/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/hammingweight/synkctl/configuration"
)

type PublicKeyResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Data    string `json:"data"` // The Base64 encoded RSA key
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type TokenResponse struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
	Data    TokenData `json:"data"`
}

type LoginPayload struct {
	Nonce     int64  `json:"nonce"`
	AreaCode  string `json:"areaCode"`
	ClientID  string `json:"client_id"`
	GrantType string `json:"grant_type"`
	Password  string `json:"password"`
	Source    string `json:"source"`
	Sign      string `json:"sign"`
	Username  string `json:"username"`
}

// md5Hash calculates the MD5 hash of a string and returns it as a hex string.
func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetSunsynkToken executes the full authentication workflow and returns the token data.
func getSunsynkToken(endpoint, username, password string) (*TokenData, error) {
	client := &http.Client{}
	source := "sunsynk"

	// 1. Generate millisecond nonce
	nonce := time.Now().UnixMilli()

	// 2. Compute Sign 1
	raw1 := fmt.Sprintf("nonce=%d&source=%sPOWER_VIEW", nonce, source)
	sign1 := md5Hash(raw1)

	// 3. Fetch the RSA public key
	urlPK := fmt.Sprintf("%s/anonymous/publicKey?nonce=%d&source=%s&sign=%s", endpoint, nonce, source, sign1)
	reqPK, err := http.NewRequest("GET", urlPK, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create public key request: %w", err)
	}

	// Recommended headers
	reqPK.Header.Set("accept", "application/json")
	reqPK.Header.Set("origin", "https://sunsynk.net")
	reqPK.Header.Set("referer", "https://sunsynk.net/")

	respPK, err := client.Do(reqPK)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key: %w", err)
	}
	defer respPK.Body.Close()

	bodyPK, err := io.ReadAll(respPK.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key response: %w", err)
	}

	var pkResp PublicKeyResponse
	if err := json.Unmarshal(bodyPK, &pkResp); err != nil {
		return nil, fmt.Errorf("failed to parse public key JSON: %w", err)
	}
	if !pkResp.Success {
		return nil, fmt.Errorf("public key API error: %s", pkResp.Msg)
	}

	base64Key := pkResp.Data
	if len(base64Key) < 10 {
		return nil, errors.New("received public key is unexpectedly short")
	}

	// 4. Compute Sign 2
	head := base64Key[:10]
	raw2 := fmt.Sprintf("nonce=%d&source=%s%s", nonce, source, head)
	sign2 := md5Hash(raw2)

	// 5. Encrypt the password using RSA PKCS1v15
	pemKey := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", base64Key)
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DER public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("key is not a valid RSA public key")
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(password))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt password: %w", err)
	}
	encryptedPasswordB64 := base64.StdEncoding.EncodeToString(encryptedBytes)

	// 6. Login to get the bearer token
	payload := LoginPayload{
		Nonce:     nonce,
		AreaCode:  "sunsynk",
		ClientID:  "csp-web",
		GrantType: "password",
		Password:  encryptedPasswordB64,
		Source:    "sunsynk",
		Sign:      sign2,
		Username:  username,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal login payload: %w", err)
	}

	reqLogin, err := http.NewRequest("POST", "https://api.sunsynk.net/oauth/token/new", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create login request: %w", err)
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to execute login request: %w", err)
	}
	defer respLogin.Body.Close()

	bodyLogin, err := io.ReadAll(respLogin.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read login response: %w", err)
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(bodyLogin, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse login JSON: %w (body: %s)", err, string(bodyLogin))
	}

	if !tokenResp.Success {
		return nil, fmt.Errorf("login failed: %s", tokenResp.Msg)
	}

	return &tokenResp.Data, nil
}

type tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// SynkClient is a type that is needed for accessing the SunSynk API; it includes the API endpoint and OAuth
// tokens. It also includes the serial number for an inverter since most requests use the serial number to
// identify which object is being requested (e.g. the details of a battery connected to the inverter.)
type SynkClient struct {
	endpoint     string
	tokens       tokens
	SerialNumber string
}

func init() {
	// Undocumented feature to work around TLS certificate problems at SunSynk (e.g. the outage of 10 March 2025).
	// Run 'export SYNK_DISABLE_TLS_CERTIFICATE_VALIDATION=1' before running synkctl to disable TLS certificate
	// validation. Generally, it's not advised to do this.
	if os.Getenv("SYNK_DISABLE_TLS_CERTIFICATE_VALIDATION") != "" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

// Authenticate uses a specify configuration to authenticate a user. If successful, a SynkClient is
// returned that can be used to make requests against the API.
func Authenticate(ctx context.Context, config *configuration.Configuration) (*SynkClient, error) {
	tokenData, err := getSunsynkToken(config.Endpoint, config.User, config.Password)
	if err != nil {
		return nil, err
	}
	tokens := &tokens{
		AccessToken: tokenData.AccessToken,
		ExpiresIn:   tokenData.ExpiresIn,
		TokenType:   tokenData.TokenType,
		Scope:       tokenData.Scope,
	}
	return &SynkClient{endpoint: config.Endpoint, tokens: *tokens, SerialNumber: config.DefaultInverterSN}, nil
}
