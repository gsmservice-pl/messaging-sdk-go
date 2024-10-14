// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type SetDefaultSenderRequest struct {
	// Sender name to set as default
	Sender string `pathParam:"style=simple,explode=false,name=sender"`
}

func (o *SetDefaultSenderRequest) GetSender() string {
	if o == nil {
		return ""
	}
	return o.Sender
}

type SetDefaultSenderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *SetDefaultSenderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetDefaultSenderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}