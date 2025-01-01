package synk

import "context"

func ReadInputState(ctx context.Context, tokens *Tokens, endpoint string, inverterSN string) (map[string]any, error) {
	path := []string{"inverter", inverterSN, "realtime", "input"}
	queryParams := map[string]string{}
	queryParams["sn"] = inverterSN
	queryParams["lan"] = "en"
	resp, err := readApiV1(ctx, tokens, endpoint, queryParams, path...)
	if err != nil {
		return nil, err
	}
	battery := &map[string]any{}
	err = unmarshallResponseData(resp, battery)
	return *battery, err
}
