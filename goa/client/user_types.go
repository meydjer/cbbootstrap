//************************************************************************//
// API "cbbootstrap": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/couchbaselabs/cbbootstrap/design
// --out=$(GOPATH)/src/github.com/couchbaselabs/cbbootstrap/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// createOrJoinClusterPayload user type.
type createOrJoinClusterPayload struct {
	ClusterID            *string `form:"cluster_id,omitempty" json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	NodeIPAddrOrHostname *string `form:"node_ip_addr_or_hostname,omitempty" json:"node_ip_addr_or_hostname,omitempty" xml:"node_ip_addr_or_hostname,omitempty"`
}

// Validate validates the createOrJoinClusterPayload type instance.
func (ut *createOrJoinClusterPayload) Validate() (err error) {
	if ut.ClusterID != nil {
		if utf8.RuneCountInString(*ut.ClusterID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.cluster_id`, *ut.ClusterID, utf8.RuneCountInString(*ut.ClusterID), 1, true))
		}
	}
	if ut.NodeIPAddrOrHostname != nil {
		if utf8.RuneCountInString(*ut.NodeIPAddrOrHostname) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.node_ip_addr_or_hostname`, *ut.NodeIPAddrOrHostname, utf8.RuneCountInString(*ut.NodeIPAddrOrHostname), 1, true))
		}
	}
	return
}

// Publicize creates CreateOrJoinClusterPayload from createOrJoinClusterPayload
func (ut *createOrJoinClusterPayload) Publicize() *CreateOrJoinClusterPayload {
	var pub CreateOrJoinClusterPayload
	if ut.ClusterID != nil {
		pub.ClusterID = ut.ClusterID
	}
	if ut.NodeIPAddrOrHostname != nil {
		pub.NodeIPAddrOrHostname = ut.NodeIPAddrOrHostname
	}
	return &pub
}

// CreateOrJoinClusterPayload user type.
type CreateOrJoinClusterPayload struct {
	ClusterID            *string `form:"cluster_id,omitempty" json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	NodeIPAddrOrHostname *string `form:"node_ip_addr_or_hostname,omitempty" json:"node_ip_addr_or_hostname,omitempty" xml:"node_ip_addr_or_hostname,omitempty"`
}

// Validate validates the CreateOrJoinClusterPayload type instance.
func (ut *CreateOrJoinClusterPayload) Validate() (err error) {
	if ut.ClusterID != nil {
		if utf8.RuneCountInString(*ut.ClusterID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.cluster_id`, *ut.ClusterID, utf8.RuneCountInString(*ut.ClusterID), 1, true))
		}
	}
	if ut.NodeIPAddrOrHostname != nil {
		if utf8.RuneCountInString(*ut.NodeIPAddrOrHostname) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.node_ip_addr_or_hostname`, *ut.NodeIPAddrOrHostname, utf8.RuneCountInString(*ut.NodeIPAddrOrHostname), 1, true))
		}
	}
	return
}