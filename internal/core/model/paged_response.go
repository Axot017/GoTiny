package model

type PagedResponse[T any] struct {
	Items     []T
	PageToken *string
}
