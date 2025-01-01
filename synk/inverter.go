package synk

import (
	"context"
	"encoding/json"
)

func ReadInverterSettings(ctx context.Context, tokens *Tokens, endpoint string, inverterSN string) (map[string]any, error) {
	path := []string{"common", "setting", inverterSN, "read"}
	resp, err := readApiV1(ctx, tokens, endpoint, nil, path...)
	if err != nil {
		return nil, err
	}
	inverterSettings := &map[string]any{}
	err = unmarshallResponseData(resp, inverterSettings)
	return *inverterSettings, err
}

func UpdateInverterSettings(ctx context.Context, tokens *Tokens, endpoint string, inverterSN string, settings map[string]any) error {
	path := []string{"common", "setting", inverterSN, "set"}
	postData, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	return updateApiV1(ctx, tokens, endpoint, string(postData), path...)
}
