package synk

import (
	"context"
)

func ReadInverterSettings(ctx context.Context, tokens *Tokens, endpoint string, inverterSN string) (map[string]any, error) {
	path := []string{"common", "setting", inverterSN, "read"}
	resp, err := readApiV1(ctx, tokens, endpoint, path...)
	if err != nil {
		return nil, err
	}
	inverterSettings := &map[string]any{}
	err = umarshallResponseData(resp, inverterSettings)
	return *inverterSettings, err
}
