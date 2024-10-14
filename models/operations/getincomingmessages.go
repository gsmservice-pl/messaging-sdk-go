// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type GetIncomingMessagesRequest struct {
	// Message IDs assigned by the system (separated by comma). The system will accept a maximum of 50 identifiers in one call.
	Ids []int64 `pathParam:"style=simple,explode=true,name=ids"`
}

func (o *GetIncomingMessagesRequest) GetIds() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Ids
}

type GetIncomingMessagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check received messages details in response body.
	IncomingMessages []components.IncomingMessage
	Headers          map[string][]string
}

func (o *GetIncomingMessagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIncomingMessagesResponse) GetIncomingMessages() []components.IncomingMessage {
	if o == nil {
		return nil
	}
	return o.IncomingMessages
}

func (o *GetIncomingMessagesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}