// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/couchbase/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbase/cbbootstrap/goa
// --version=v1.0.0
//
// API "cbbootstrap": cluster Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateOrJoinClusterPath computes a request path to the create_or_join action of cluster.
func CreateOrJoinClusterPath() string {

	return fmt.Sprintf("/cluster")
}

// Create a new Couchbase Cluster
func (c *Client) CreateOrJoinCluster(ctx context.Context, path string, payload *CreateOrJoinClusterPayload) (*http.Response, error) {
	req, err := c.NewCreateOrJoinClusterRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateOrJoinClusterRequest create the request corresponding to the create_or_join action endpoint of the cluster resource.
func (c *Client) NewCreateOrJoinClusterRequest(ctx context.Context, path string, payload *CreateOrJoinClusterPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetStatusClusterPayload is the cluster get_status action payload.
type GetStatusClusterPayload struct {
	ClusterID            string  `form:"cluster_id" json:"cluster_id" xml:"cluster_id"`
	NodeIPAddrOrHostname *string `form:"node_ip_addr_or_hostname,omitempty" json:"node_ip_addr_or_hostname,omitempty" xml:"node_ip_addr_or_hostname,omitempty"`
}

// GetStatusClusterPath computes a request path to the get_status action of cluster.
func GetStatusClusterPath() string {

	return fmt.Sprintf("/cluster/get_status")
}

// Get Couchbase Cluster by ID.  Works around URL encoding issues seen in GET with :cluster_id URL param
func (c *Client) GetStatusCluster(ctx context.Context, path string, payload *GetStatusClusterPayload) (*http.Response, error) {
	req, err := c.NewGetStatusClusterRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetStatusClusterRequest create the request corresponding to the get_status action endpoint of the cluster resource.
func (c *Client) NewGetStatusClusterRequest(ctx context.Context, path string, payload *GetStatusClusterPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StatusClusterPath computes a request path to the status action of cluster.
func StatusClusterPath(clusterID string) string {
	param0 := clusterID

	return fmt.Sprintf("/cluster/%s", param0)
}

// Get Couchbase Cluster by ID
func (c *Client) StatusCluster(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatusClusterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusClusterRequest create the request corresponding to the status action endpoint of the cluster resource.
func (c *Client) NewStatusClusterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
