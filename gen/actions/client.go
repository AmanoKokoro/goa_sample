// Code generated by goa v3.2.5, DO NOT EDIT.
//
// actions client
//
// Command:
// $ goa gen sample/design

package actions

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "actions" service client.
type Client struct {
	PingEndpoint  goa.Endpoint
	HelloEndpoint goa.Endpoint
	IDEndpoint    goa.Endpoint
}

// NewClient initializes a "actions" service client given the endpoints.
func NewClient(ping, hello, id goa.Endpoint) *Client {
	return &Client{
		PingEndpoint:  ping,
		HelloEndpoint: hello,
		IDEndpoint:    id,
	}
}

// Ping calls the "ping" endpoint of the "actions" service.
func (c *Client) Ping(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.PingEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Hello calls the "hello" endpoint of the "actions" service.
func (c *Client) Hello(ctx context.Context, p *HelloPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.HelloEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// ID calls the "ID" endpoint of the "actions" service.
func (c *Client) ID(ctx context.Context, p *IDPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.IDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
