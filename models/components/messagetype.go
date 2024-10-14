// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// MessageType - Message type according to the table:
//
// |type|Description|
// |----|-----------|
// |  1 | SMS PRO   |
// |  3 | SMS ECO   |
// |  4 | SMS 2WAY  |
// | 10 | MMS       |
type MessageType int64

const (
	MessageTypeSmsPro    MessageType = 1
	MessageTypeSmsEco    MessageType = 3
	MessageTypeSmsTwoWay MessageType = 4
	MessageTypeMms       MessageType = 10
)

func (e MessageType) ToPointer() *MessageType {
	return &e
}
