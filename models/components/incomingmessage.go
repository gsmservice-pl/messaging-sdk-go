// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gsmservice-pl/messaging-sdk-go/internal/utils"
	"time"
)

// IncomingMessage - An object with the properties of the received message
type IncomingMessage struct {
	// Unique identifier of incoming message
	ID *int64 `json:"id,omitempty"`
	// Login of the account (sub-account) on which the message was received
	Login *string `json:"login,omitempty"`
	// Recipient number (or service name)
	Recipient *string `json:"recipient,omitempty"`
	// Message sender number (or alphanumeric name)
	Sender *string `json:"sender,omitempty"`
	// Sender name (matched with phonebook)
	PhonebookSenderName *string `json:"phonebook_sender_name,omitempty"`
	// Date and time of message receipt in ISO 8601 format
	Date *time.Time `json:"date,omitempty"`
	// Received message content
	Message *string `json:"message,omitempty"`
	// Dedicated Mobile Originated service number (if the message was received on this number)
	DedicatedNumber *string `json:"dedicated_number,omitempty"`
	// Dedicated Mobile Originated service prefix (if the message was received on this number with such prefix)
	DedicatedPrefix *string `json:"dedicated_prefix,omitempty"`
}

func (i IncomingMessage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncomingMessage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncomingMessage) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncomingMessage) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

func (o *IncomingMessage) GetRecipient() *string {
	if o == nil {
		return nil
	}
	return o.Recipient
}

func (o *IncomingMessage) GetSender() *string {
	if o == nil {
		return nil
	}
	return o.Sender
}

func (o *IncomingMessage) GetPhonebookSenderName() *string {
	if o == nil {
		return nil
	}
	return o.PhonebookSenderName
}

func (o *IncomingMessage) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *IncomingMessage) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *IncomingMessage) GetDedicatedNumber() *string {
	if o == nil {
		return nil
	}
	return o.DedicatedNumber
}

func (o *IncomingMessage) GetDedicatedPrefix() *string {
	if o == nil {
		return nil
	}
	return o.DedicatedPrefix
}