package synk

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
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

func CountInverters(ctx context.Context, tokens *Tokens, endpoint string) (int, error) {
	path := []string{"inverters", "count"}
	resp, err := readApiV1(ctx, tokens, endpoint, nil, path...)
	if err != nil {
		return 0, err
	}
	data := map[string]any{}
	unmarshallResponseData(resp, &data)
	return int(data["total"].(float64)), err
}

func GetInverterSerialNumbers(ctx context.Context, tokens *Tokens, endpoint string, page int, pageSize int) ([]string, error) {
	path := []string{"inverters"}
	queryParams := map[string]string{}
	queryParams["page"] = strconv.Itoa(page)
	queryParams["limit"] = strconv.Itoa(pageSize)
	queryParams["type"] = "-2"
	queryParams["status"] = "-1"
	resp, err := readApiV1(ctx, tokens, endpoint, queryParams, path...)
	if err != nil {
		return nil, err
	}
	data := map[string]any{}
	err = unmarshallResponseData(resp, &data)
	if err != nil {
		return nil, err
	}
	allInverters, ok := data["infos"]
	if !ok {
		return nil, errors.New("Can't retrieve serial numbers from response")
	}
	inverterList := allInverters.([]any)
	responseList := []string{}
	for _, inv := range inverterList {
		sn := inv.(map[string]any)["sn"]
		responseList = append(responseList, sn.(string))
	}
	return responseList, nil
}
