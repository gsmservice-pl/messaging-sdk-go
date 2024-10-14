# IncomingMessage

An object with the properties of the received message


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                     | **int64*                                                                                                 | :heavy_minus_sign:                                                                                       | Unique identifier of incoming message                                                                    | 45544                                                                                                    |
| `Login`                                                                                                  | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Login of the account (sub-account) on which the message was received                                     | some-user                                                                                                |
| `Recipient`                                                                                              | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Recipient number (or service name)                                                                       | +48999000555                                                                                             |
| `Sender`                                                                                                 | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Message sender number (or alphanumeric name)                                                             | +48999888777                                                                                             |
| `PhonebookSenderName`                                                                                    | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Sender name (matched with phonebook)                                                                     | Jan Nowak                                                                                                |
| `Date`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                               | :heavy_minus_sign:                                                                                       | Date and time of message receipt in ISO 8601 format                                                      | 2024-05-31T05:17:35Z                                                                                     |
| `Message`                                                                                                | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Received message content                                                                                 | To jest treść odebranego SMSa                                                                            |
| `DedicatedNumber`                                                                                        | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Dedicated Mobile Originated service number (if the message was received on this number)                  | +48111222444                                                                                             |
| `DedicatedPrefix`                                                                                        | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Dedicated Mobile Originated service prefix (if the message was received on this number with such prefix) | LATO                                                                                                     |