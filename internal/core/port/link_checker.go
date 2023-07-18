package port

import "context"

type LinkChecker interface {
	CheckUrl(ctx context.Context, url string) (bool, error)
}
