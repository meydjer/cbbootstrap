// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/couchbase/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbase/cbbootstrap/goa
// --version=v1.0.0
//
// API "cbbootstrap": Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the cbbootstrap service client.
type Client struct {
	*goaclient.Client
	Encoder *goa.HTTPEncoder
	Decoder *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}
