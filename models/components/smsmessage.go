// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/internal/utils"
	"time"
)

type SmsMessageRecipientsType string

const (
	SmsMessageRecipientsTypeStr                       SmsMessageRecipientsType = "str"
	SmsMessageRecipientsTypeArrayOfStr                SmsMessageRecipientsType = "arrayOfStr"
	SmsMessageRecipientsTypePhoneNumberWithCid        SmsMessageRecipientsType = "PhoneNumberWithCid"
	SmsMessageRecipientsTypeArrayOfPhoneNumberWithCid SmsMessageRecipientsType = "arrayOfPhoneNumberWithCid"
)

// SmsMessageRecipients - The recipient number or multiple recipients numbers of single message. To set one recipient, please use `components.CreateSmsMessageRecipientsStr()` method simply passing to it a `string` with his phone number. To set multiple recipients, please use `components.CreateSmsMessageRecipientsArrayOfStr()` method passing to it `[]string` with recipients.\r\n\r\nOptionally you can also set custom id (user identifier) for each message - use `components.CreateSmsMessageRecipientsPhoneNumberWithCid()` method passing `PhoneNumberWithCid` struct (in case of single recipient) or `operations.CreateSmsMessageRecipientsArrayOfPhoneNumberWithCid()` method passing []PhoneNumberWithCid (in case of multiple recipients).
type SmsMessageRecipients struct {
	Str                       *string              `queryParam:"inline"`
	ArrayOfStr                []string             `queryParam:"inline"`
	PhoneNumberWithCid        *PhoneNumberWithCid  `queryParam:"inline"`
	ArrayOfPhoneNumberWithCid []PhoneNumberWithCid `queryParam:"inline"`

	Type SmsMessageRecipientsType
}

func CreateSmsMessageRecipientsStr(str string) SmsMessageRecipients {
	typ := SmsMessageRecipientsTypeStr

	return SmsMessageRecipients{
		Str:  &str,
		Type: typ,
	}
}

func CreateSmsMessageRecipientsArrayOfStr(arrayOfStr []string) SmsMessageRecipients {
	typ := SmsMessageRecipientsTypeArrayOfStr

	return SmsMessageRecipients{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func CreateSmsMessageRecipientsPhoneNumberWithCid(phoneNumberWithCid PhoneNumberWithCid) SmsMessageRecipients {
	typ := SmsMessageRecipientsTypePhoneNumberWithCid

	return SmsMessageRecipients{
		PhoneNumberWithCid: &phoneNumberWithCid,
		Type:               typ,
	}
}

func CreateSmsMessageRecipientsArrayOfPhoneNumberWithCid(arrayOfPhoneNumberWithCid []PhoneNumberWithCid) SmsMessageRecipients {
	typ := SmsMessageRecipientsTypeArrayOfPhoneNumberWithCid

	return SmsMessageRecipients{
		ArrayOfPhoneNumberWithCid: arrayOfPhoneNumberWithCid,
		Type:                      typ,
	}
}

func (u *SmsMessageRecipients) UnmarshalJSON(data []byte) error {

	var phoneNumberWithCid PhoneNumberWithCid = PhoneNumberWithCid{}
	if err := utils.UnmarshalJSON(data, &phoneNumberWithCid, "", true, true); err == nil {
		u.PhoneNumberWithCid = &phoneNumberWithCid
		u.Type = SmsMessageRecipientsTypePhoneNumberWithCid
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = SmsMessageRecipientsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = SmsMessageRecipientsTypeArrayOfStr
		return nil
	}

	var arrayOfPhoneNumberWithCid []PhoneNumberWithCid = []PhoneNumberWithCid{}
	if err := utils.UnmarshalJSON(data, &arrayOfPhoneNumberWithCid, "", true, true); err == nil {
		u.ArrayOfPhoneNumberWithCid = arrayOfPhoneNumberWithCid
		u.Type = SmsMessageRecipientsTypeArrayOfPhoneNumberWithCid
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SmsMessageRecipients", string(data))
}

func (u SmsMessageRecipients) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	if u.PhoneNumberWithCid != nil {
		return utils.MarshalJSON(u.PhoneNumberWithCid, "", true)
	}

	if u.ArrayOfPhoneNumberWithCid != nil {
		return utils.MarshalJSON(u.ArrayOfPhoneNumberWithCid, "", true)
	}

	return nil, errors.New("could not marshal union type SmsMessageRecipients: all fields are null")
}

// SmsMessage - An object with a new SMS message properties
type SmsMessage struct {
	// The recipient number or multiple recipients numbers of single message. To set one recipient, please use `components.CreateSmsMessageRecipientsStr()` method simply passing to it a `string` with his phone number. To set multiple recipients, please use `components.CreateSmsMessageRecipientsArrayOfStr()` method passing to it `[]string` with recipients.\r\n\r\nOptionally you can also set custom id (user identifier) for each message - use `components.CreateSmsMessageRecipientsPhoneNumberWithCid()` method passing `PhoneNumberWithCid` struct (in case of single recipient) or `operations.CreateSmsMessageRecipientsArrayOfPhoneNumberWithCid()` method passing []PhoneNumberWithCid (in case of multiple recipients).
	Recipients SmsMessageRecipients `json:"recipients"`
	// SMS message content
	Message string `json:"message"`
	// SMS sender name
	Sender *string `default:"Bramka SMS" json:"sender"`
	// SMS type (components.SmsTypeSmsPro -> SMS PRO, components.SmsTypeSmsEco -> SMS ECO, components.SmsTypeSmsTwoWay -> SMS 2WAY)
	Type *SmsType `default:"1" json:"type"`
	// Should the message be sent with special characters, e.g. Polish diacritics (if any)? If *false*, those characters will be automatically replaced with their equivalents. If *true* your message will be sent as **unicode** but the message will be able to consist of fewer characters.
	Unicode *bool `default:"false" json:"unicode"`
	// Should the message to be sent with class 0 (FLASH)?
	Flash *bool `default:"false" json:"flash"`
	// Scheduled future date and time of sending the message (in ISO 8601 format). If missing or null - message will be sent immediately
	Date *time.Time `default:"null" json:"date"`
}

func (s SmsMessage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SmsMessage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SmsMessage) GetRecipients() SmsMessageRecipients {
	if o == nil {
		return SmsMessageRecipients{}
	}
	return o.Recipients
}

func (o *SmsMessage) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *SmsMessage) GetSender() *string {
	if o == nil {
		return nil
	}
	return o.Sender
}

func (o *SmsMessage) GetType() *SmsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SmsMessage) GetUnicode() *bool {
	if o == nil {
		return nil
	}
	return o.Unicode
}

func (o *SmsMessage) GetFlash() *bool {
	if o == nil {
		return nil
	}
	return o.Flash
}

func (o *SmsMessage) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}
