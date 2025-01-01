package synk

import "context"

func ReadBattery(ctx context.Context, tokens *Tokens, endpoint string, inverterSN string) (map[string]any, error) {
	path := []string{"inverter", "battery", inverterSN, "realtime"}
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
