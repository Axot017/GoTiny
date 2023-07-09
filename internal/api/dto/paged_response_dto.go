package dto

import "gotiny/internal/core/model"

type PagedResponseDto[T any] struct {
	PageToken *string `json:"page_token"`
	Items     []T     `json:"data"`
}

func PagedResponseDtoFromModel[TIn any, TOut any](
	model model.PagedResponse[TIn],
	itemMapper func(TIn) TOut,
) PagedResponseDto[TOut] {
	result := make([]TOut, len(model.Items))
	for i, item := range model.Items {
		result[i] = itemMapper(item)
	}

	return PagedResponseDto[TOut]{
		PageToken: model.PageToken,
		Items:     result,
	}
}
