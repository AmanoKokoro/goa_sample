// Code generated by goa v3.2.5, DO NOT EDIT.
//
// actions HTTP client encoders and decoders
//
// Command:
// $ goa gen sample/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	actions "sample/gen/actions"

	goahttp "goa.design/goa/v3/http"
)

// BuildPingRequest instantiates a HTTP request object with method and path set
// to call the "actions" service "ping" endpoint
func (c *Client) BuildPingRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PingActionsPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("actions", "ping", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodePingResponse returns a decoder for responses returned by the actions
// ping endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodePingResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("actions", "ping", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("actions", "ping", resp.StatusCode, string(body))
		}
	}
}

// BuildHelloRequest instantiates a HTTP request object with method and path
// set to call the "actions" service "hello" endpoint
func (c *Client) BuildHelloRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: HelloActionsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("actions", "hello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeHelloRequest returns an encoder for requests sent to the actions hello
// server.
func EncodeHelloRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*actions.HelloPayload)
		if !ok {
			return goahttp.ErrInvalidType("actions", "hello", "*actions.HelloPayload", v)
		}
		body := NewHelloRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("actions", "hello", err)
		}
		return nil
	}
}

// DecodeHelloResponse returns a decoder for responses returned by the actions
// hello endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("actions", "hello", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("actions", "hello", resp.StatusCode, string(body))
		}
	}
}

// BuildIDRequest instantiates a HTTP request object with method and path set
// to call the "actions" service "ID" endpoint
func (c *Client) BuildIDRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*actions.IDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("actions", "ID", "*actions.IDPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: IDActionsPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("actions", "ID", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeIDResponse returns a decoder for responses returned by the actions ID
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("actions", "ID", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("actions", "ID", resp.StatusCode, string(body))
		}
	}
}
