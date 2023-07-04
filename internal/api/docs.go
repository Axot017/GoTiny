// Package api GoTiny API
//
// TODO: add description
//
//	Schemes: https
//	Host: yfe9pxjetr.eu-central-1.awsapprunner.com
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
