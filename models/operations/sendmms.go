// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/internal/utils"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type SendMmsRequestBodyType string

const (
	SendMmsRequestBodyTypeMmsMessage        SendMmsRequestBodyType = "MmsMessage"
	SendMmsRequestBodyTypeArrayOfMmsMessage SendMmsRequestBodyType = "arrayOfMmsMessage"
)

// SendMmsRequestBody - To send a single MMS or messages with the same content to multiple recipients, please use `operations.CreateSendMmsRequestBodyMmsMessage()` method with a single `MmsMessage` object with the properties of this message. To send multiple messages with different content at the same time, please use `operations.CreateSendMmsRequestBodyArrayOfMmsMessage()` method passing to it array of type `[]MmsMessage` with the properties of each message.
type SendMmsRequestBody struct {
	MmsMessage        *components.MmsMessage
	ArrayOfMmsMessage []components.MmsMessage

	Type SendMmsRequestBodyType
}

func CreateSendMmsRequestBodyMmsMessage(mmsMessage components.MmsMessage) SendMmsRequestBody {
	typ := SendMmsRequestBodyTypeMmsMessage

	return SendMmsRequestBody{
		MmsMessage: &mmsMessage,
		Type:       typ,
	}
}

func CreateSendMmsRequestBodyArrayOfMmsMessage(arrayOfMmsMessage []components.MmsMessage) SendMmsRequestBody {
	typ := SendMmsRequestBodyTypeArrayOfMmsMessage

	return SendMmsRequestBody{
		ArrayOfMmsMessage: arrayOfMmsMessage,
		Type:              typ,
	}
}

func (u *SendMmsRequestBody) UnmarshalJSON(data []byte) error {

	var mmsMessage components.MmsMessage = components.MmsMessage{}
	if err := utils.UnmarshalJSON(data, &mmsMessage, "", true, true); err == nil {
		u.MmsMessage = &mmsMessage
		u.Type = SendMmsRequestBodyTypeMmsMessage
		return nil
	}

	var arrayOfMmsMessage []components.MmsMessage = []components.MmsMessage{}
	if err := utils.UnmarshalJSON(data, &arrayOfMmsMessage, "", true, true); err == nil {
		u.ArrayOfMmsMessage = arrayOfMmsMessage
		u.Type = SendMmsRequestBodyTypeArrayOfMmsMessage
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SendMmsRequestBody", string(data))
}

func (u SendMmsRequestBody) MarshalJSON() ([]byte, error) {
	if u.MmsMessage != nil {
		return utils.MarshalJSON(u.MmsMessage, "", true)
	}

	if u.ArrayOfMmsMessage != nil {
		return utils.MarshalJSON(u.ArrayOfMmsMessage, "", true)
	}

	return nil, errors.New("could not marshal union type SendMmsRequestBody: all fields are null")
}

type SendMmsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check statuses of particular messages in each `Message` object.
	Messages []components.Message
	Headers  map[string][]string
}

func (o *SendMmsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SendMmsResponse) GetMessages() []components.Message {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *SendMmsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
