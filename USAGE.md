<!-- Start SDK Example Usage [usage] -->
### Sending single SMS Message

This example demonstrates simple sending SMS message to a single recipient:

```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
	)

	res, err := s.Outgoing.Sms.Send(ctx, operations.CreateSendSmsRequestBodyArrayOfSmsMessage(
		[]components.SmsMessage{
			components.SmsMessage{
				Recipients: components.CreateSmsMessageRecipientsStr(
					"+48999999999",
				),
				Message: "To jest treść wiadomości",
			},
		},
	))
	if err != nil {
		log.Fatal(err)
	}
	if res.Messages != nil {
		// handle response
	}
}

```

### Sending single MMS Message

This example demonstrates simple sending MMS message to a single recipient:

```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
	)

	res, err := s.Outgoing.Mms.Send(ctx, operations.CreateSendMmsRequestBodyMmsMessage(
		components.MmsMessage{
			Recipients: components.CreateRecipientsPhoneNumberWithCid(
				components.PhoneNumberWithCid{
					Nr:  "+48999999999",
					Cid: messagingsdkgo.String("my-id-1113"),
				},
			),
			Subject: messagingsdkgo.String("To jest temat wiadomości"),
			Message: messagingsdkgo.String("To jest treść wiadomości"),
			Attachments: messagingsdkgo.Pointer(components.CreateAttachmentsStr(
				"<file_body in base64 format>",
			)),
		},
	))
	if err != nil {
		log.Fatal(err)
	}
	if res.Messages != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->