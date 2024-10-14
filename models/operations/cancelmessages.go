// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type CancelMessagesRequest struct {
	// Message IDs assigned by the system (separated by comma). The system will accept a maximum of 50 identifiers in one call.
	Ids []int64 `pathParam:"style=simple,explode=true,name=ids"`
}

func (o *CancelMessagesRequest) GetIds() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Ids
}

type CancelMessagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check the status of messages cancellation in response body.
	CancelledMessages []components.CancelledMessage
	Headers           map[string][]string
}

func (o *CancelMessagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CancelMessagesResponse) GetCancelledMessages() []components.CancelledMessage {
	if o == nil {
		return nil
	}
	return o.CancelledMessages
}

func (o *CancelMessagesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}