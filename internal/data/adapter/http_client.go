package adapter

import (
	"context"
	"net/http"
	"strings"

	"golang.org/x/exp/slog"
)

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (t *HttpClient) CheckUrl(ctx context.Context, url string) (bool, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		slog.ErrorCtx(ctx, "error creating request", "err", err)
		return false, err
	}
	result, err := http.DefaultClient.Do(request)
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return false, nil
		}
		slog.ErrorCtx(ctx, "error making request", "err", err)

		return false, err
	}

	return result.StatusCode < 400, nil
}
