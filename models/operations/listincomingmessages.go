// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gsmservice-pl/messaging-sdk-go/v3/internal/utils"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
)

type ListIncomingMessagesRequest struct {
	// Page number of results
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of results on one page
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListIncomingMessagesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListIncomingMessagesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListIncomingMessagesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListIncomingMessagesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type ListIncomingMessagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check received message details in each `IncomingMessage` object
	IncomingMessages []components.IncomingMessage
	Headers          map[string][]string
}

func (o *ListIncomingMessagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListIncomingMessagesResponse) GetIncomingMessages() []components.IncomingMessage {
	if o == nil {
		return nil
	}
	return o.IncomingMessages
}

func (o *ListIncomingMessagesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
