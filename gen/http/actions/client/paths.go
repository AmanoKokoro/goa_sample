// Code generated by goa v3.2.5, DO NOT EDIT.
//
// HTTP request path constructors for the actions service.
//
// Command:
// $ goa gen sample/design

package client

import (
	"fmt"
)

// PingActionsPath returns the URL path to the actions service ping HTTP endpoint.
func PingActionsPath() string {
	return "/ping"
}

// HelloActionsPath returns the URL path to the actions service hello HTTP endpoint.
func HelloActionsPath() string {
	return "/hello"
}

// IDActionsPath returns the URL path to the actions service ID HTTP endpoint.
func IDActionsPath(id int) string {
	return fmt.Sprintf("/%v", id)
}
