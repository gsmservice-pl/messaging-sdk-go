// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type GetMessagesRequest struct {
	// Message IDs assigned by the system (separated by comma). The system will accept a maximum of 50 identifiers in one call.
	Ids []int64 `pathParam:"style=simple,explode=true,name=ids"`
}

func (o *GetMessagesRequest) GetIds() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Ids
}

type GetMessagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check messages details in response body.
	Messages []components.Message
	Headers  map[string][]string
}

func (o *GetMessagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMessagesResponse) GetMessages() []components.Message {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *GetMessagesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}