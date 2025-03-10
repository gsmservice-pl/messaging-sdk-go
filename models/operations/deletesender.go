// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
)

type DeleteSenderRequest struct {
	// Sender name to be removed
	Sender string `pathParam:"style=simple,explode=false,name=sender"`
}

func (o *DeleteSenderRequest) GetSender() string {
	if o == nil {
		return ""
	}
	return o.Sender
}

type DeleteSenderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *DeleteSenderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteSenderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
