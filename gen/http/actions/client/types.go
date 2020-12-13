// Code generated by goa v3.2.5, DO NOT EDIT.
//
// actions HTTP client types
//
// Command:
// $ goa gen sample/design

package client

import (
	actions "sample/gen/actions"
)

// HelloRequestBody is the type of the "actions" service "hello" endpoint HTTP
// request body.
type HelloRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// NewHelloRequestBody builds the HTTP request body from the payload of the
// "hello" endpoint of the "actions" service.
func NewHelloRequestBody(p *actions.HelloPayload) *HelloRequestBody {
	body := &HelloRequestBody{
		Name: p.Name,
	}
	return body
}
