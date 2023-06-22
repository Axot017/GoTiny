package core

import "gotiny/internal/core/usecase"

func Providers() []interface{} {
	return []interface{}{
		usecase.NewCreateShortLink,
	}
}
