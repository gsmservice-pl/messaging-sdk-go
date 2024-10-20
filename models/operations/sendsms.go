// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/internal/utils"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
)

type SendSmsRequestBodyType string

const (
	SendSmsRequestBodyTypeSmsMessage        SendSmsRequestBodyType = "SmsMessage"
	SendSmsRequestBodyTypeArrayOfSmsMessage SendSmsRequestBodyType = "arrayOfSmsMessage"
)

// SendSmsRequestBody - To send a single SMS or messages with the same content to multiple recipients, please use `operations.CreateGetSmsPriceRequestBodySmsMessage()` method with a single `SmsMessage` object with the properties of this message. To send multiple messages with different content at the same time, please use `operations.CreateSendSmsRequestBodyArrayOfSmsMessage()` method passing to it array of type `[]SmsMessage` with the properties of each message.
type SendSmsRequestBody struct {
	SmsMessage        *components.SmsMessage
	ArrayOfSmsMessage []components.SmsMessage

	Type SendSmsRequestBodyType
}

func CreateSendSmsRequestBodySmsMessage(smsMessage components.SmsMessage) SendSmsRequestBody {
	typ := SendSmsRequestBodyTypeSmsMessage

	return SendSmsRequestBody{
		SmsMessage: &smsMessage,
		Type:       typ,
	}
}

func CreateSendSmsRequestBodyArrayOfSmsMessage(arrayOfSmsMessage []components.SmsMessage) SendSmsRequestBody {
	typ := SendSmsRequestBodyTypeArrayOfSmsMessage

	return SendSmsRequestBody{
		ArrayOfSmsMessage: arrayOfSmsMessage,
		Type:              typ,
	}
}

func (u *SendSmsRequestBody) UnmarshalJSON(data []byte) error {

	var smsMessage components.SmsMessage = components.SmsMessage{}
	if err := utils.UnmarshalJSON(data, &smsMessage, "", true, true); err == nil {
		u.SmsMessage = &smsMessage
		u.Type = SendSmsRequestBodyTypeSmsMessage
		return nil
	}

	var arrayOfSmsMessage []components.SmsMessage = []components.SmsMessage{}
	if err := utils.UnmarshalJSON(data, &arrayOfSmsMessage, "", true, true); err == nil {
		u.ArrayOfSmsMessage = arrayOfSmsMessage
		u.Type = SendSmsRequestBodyTypeArrayOfSmsMessage
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SendSmsRequestBody", string(data))
}

func (u SendSmsRequestBody) MarshalJSON() ([]byte, error) {
	if u.SmsMessage != nil {
		return utils.MarshalJSON(u.SmsMessage, "", true)
	}

	if u.ArrayOfSmsMessage != nil {
		return utils.MarshalJSON(u.ArrayOfSmsMessage, "", true)
	}

	return nil, errors.New("could not marshal union type SendSmsRequestBody: all fields are null")
}

type SendSmsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request was processed successfully. Please check statuses of particular messages in each `Message` object.
	Messages []components.Message
	Headers  map[string][]string
}

func (o *SendSmsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SendSmsResponse) GetMessages() []components.Message {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *SendSmsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
