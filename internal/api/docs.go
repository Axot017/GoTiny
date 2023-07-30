// Package api GoTiny API
//
// GoTiny is a web app that lets you create short URLs from long ones.
// You can set a TTL or a visit limit for your links and track their clicks.
// GoTiny is open source and free to use.
//
//	BasePath: /
//	Version: 0.0.1
//	Contact: Mateusz Ledwoń<mateuszledwon@duck.com> https://github.com/Axot017
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package api

import "gotiny/internal/api/dto"

// Empty response
// swagger:response emptyResponse
type emptyResponse struct{}

// Error response
// swagger:response errorResponse
type errorResponse struct {
	// The error message
	// in: body
	Body dto.ErrorDto
}
