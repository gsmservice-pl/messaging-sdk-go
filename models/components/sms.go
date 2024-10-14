// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/internal/utils"
	"time"
)

type RecipientsType string

const (
	RecipientsTypeStr                       RecipientsType = "str"
	RecipientsTypeArrayOfStr                RecipientsType = "arrayOfStr"
	RecipientsTypePhoneNumberWithCid        RecipientsType = "PhoneNumberWithCid"
	RecipientsTypeArrayOfPhoneNumberWithCid RecipientsType = "arrayOfPhoneNumberWithCid"
)

// Recipients - The recipient number or multiple recipients numbers of single message. To set one recipient, simply pass here a `string` with his phone number. To set multiple recipients, pass here a simple `array` of `string`. Optionally you can also set custom id (user identifier) for each message - pass `PhoneNumberWithCid` object (in case of single recipient) or `Array` of `PhoneNumberWithCid` (in case of multiple recipients).
type Recipients struct {
	Str                       *string
	ArrayOfStr                []string
	PhoneNumberWithCid        *PhoneNumberWithCid
	ArrayOfPhoneNumberWithCid []PhoneNumberWithCid

	Type RecipientsType
}

func CreateRecipientsStr(str string) Recipients {
	typ := RecipientsTypeStr

	return Recipients{
		Str:  &str,
		Type: typ,
	}
}

func CreateRecipientsArrayOfStr(arrayOfStr []string) Recipients {
	typ := RecipientsTypeArrayOfStr

	return Recipients{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func CreateRecipientsPhoneNumberWithCid(phoneNumberWithCid PhoneNumberWithCid) Recipients {
	typ := RecipientsTypePhoneNumberWithCid

	return Recipients{
		PhoneNumberWithCid: &phoneNumberWithCid,
		Type:               typ,
	}
}

func CreateRecipientsArrayOfPhoneNumberWithCid(arrayOfPhoneNumberWithCid []PhoneNumberWithCid) Recipients {
	typ := RecipientsTypeArrayOfPhoneNumberWithCid

	return Recipients{
		ArrayOfPhoneNumberWithCid: arrayOfPhoneNumberWithCid,
		Type:                      typ,
	}
}

func (u *Recipients) UnmarshalJSON(data []byte) error {

	var phoneNumberWithCid PhoneNumberWithCid = PhoneNumberWithCid{}
	if err := utils.UnmarshalJSON(data, &phoneNumberWithCid, "", true, true); err == nil {
		u.PhoneNumberWithCid = &phoneNumberWithCid
		u.Type = RecipientsTypePhoneNumberWithCid
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = RecipientsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = RecipientsTypeArrayOfStr
		return nil
	}

	var arrayOfPhoneNumberWithCid []PhoneNumberWithCid = []PhoneNumberWithCid{}
	if err := utils.UnmarshalJSON(data, &arrayOfPhoneNumberWithCid, "", true, true); err == nil {
		u.ArrayOfPhoneNumberWithCid = arrayOfPhoneNumberWithCid
		u.Type = RecipientsTypeArrayOfPhoneNumberWithCid
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Recipients", string(data))
}

func (u Recipients) MarshalJSON() ([]byte, error) {
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

	return nil, errors.New("could not marshal union type Recipients: all fields are null")
}

// Sms - An object with a new SMS message properties
type Sms struct {
	// The recipient number or multiple recipients numbers of single message. To set one recipient, simply pass here a `string` with his phone number. To set multiple recipients, pass here a simple `array` of `string`. Optionally you can also set custom id (user identifier) for each message - pass `PhoneNumberWithCid` object (in case of single recipient) or `Array` of `PhoneNumberWithCid` (in case of multiple recipients).
	Recipients Recipients `json:"recipients"`
	// SMS message content
	Message string `json:"message"`
	// SMS sender name
	Sender *string `default:"Bramka SMS" json:"sender"`
	// SMS type according to the table
	//
	// |type|Description|
	// |----|-----------|
	// |  1 |  SMS PRO  |
	// |  3 |  SMS ECO  |
	// |  4 |  SMS 2WAY |
	Type *SmsType `default:"1" json:"type"`
	// Should the message be sent with special characters, e.g. Polish diacritics (if any)? If *false*, those characters will be automatically replaced with their equivalents. If *true* your message will be sent as **unicode** but the message will be able to consist of fewer characters.
	Unicode *bool `default:"false" json:"unicode"`
	// Should the message to be sent with class 0 (FLASH)?
	Flash *bool `default:"false" json:"flash"`
	// Scheduled future date and time of sending the message (in ISO 8601 format). If missing or null - message will be sent immediately
	Date *time.Time `default:"null" json:"date"`
}

func (s Sms) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Sms) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Sms) GetRecipients() Recipients {
	if o == nil {
		return Recipients{}
	}
	return o.Recipients
}

func (o *Sms) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Sms) GetSender() *string {
	if o == nil {
		return nil
	}
	return o.Sender
}

func (o *Sms) GetType() *SmsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Sms) GetUnicode() *bool {
	if o == nil {
		return nil
	}
	return o.Unicode
}

func (o *Sms) GetFlash() *bool {
	if o == nil {
		return nil
	}
	return o.Flash
}

func (o *Sms) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}