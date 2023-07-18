package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"

	"gotiny/internal/config"
	"gotiny/internal/core/model"
	"gotiny/internal/data/dto"
)

type IpStackApiClient struct {
	config *config.Config
}

func NewIpStackApiClient(
	config *config.Config,
) *IpStackApiClient {
	return &IpStackApiClient{
		config: config,
	}
}

func (apiClient *IpStackApiClient) GetIpDetails(
	ctx context.Context,
	ip string,
) (model.IpDetails, error) {
	token := apiClient.config.IpStackToken()
	if token == "" {
		slog.WarnCtx(ctx, "ipstack token is empty")
		return model.IpDetails{}, errors.New("ipstack token is empty")
	}

	url := fmt.Sprintf(
		"%s/%s?access_key=%s",
		apiClient.config.IpStackBaseUrl(),
		ip,
		apiClient.config.IpStackToken(),
	)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.ErrorCtx(ctx, "error creating request", "err", err)
		return model.IpDetails{}, err
	}
	result, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.ErrorCtx(ctx, "error making request", "err", err)
		return model.IpDetails{}, err
	}
	jsonDecoder := json.NewDecoder(result.Body)
	var body dto.IpStackResponseDto
	err = jsonDecoder.Decode(&body)

	if err != nil {
		slog.ErrorCtx(ctx, "error decoding response", "err", err)
		return model.IpDetails{}, err
	}

	return body.ToIpDetails(), nil
}
