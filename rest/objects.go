package rest

import "context"

func (synkClient *SynkClient) ReadBattery(ctx context.Context) (SynkObject, error) {
	path := []string{"inverter", "battery", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	return synkClient.readApiV1(ctx, queryParams, path...)
}

func (synkClient *SynkClient) ReadGrid(ctx context.Context) (SynkObject, error) {
	path := []string{"inverter", "grid", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber}
	return synkClient.readApiV1(ctx, queryParams, path...)
}

func (synkClient *SynkClient) ReadInputState(ctx context.Context) (SynkObject, error) {
	path := []string{"inverter", synkClient.SerialNumber, "realtime", "input"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	return synkClient.readApiV1(ctx, queryParams, path...)
}

func (synkClient *SynkClient) ReadLoad(ctx context.Context) (SynkObject, error) {
	path := []string{"inverter", "load", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	return synkClient.readApiV1(ctx, queryParams, path...)
}
