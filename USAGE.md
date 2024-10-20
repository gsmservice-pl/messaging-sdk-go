<!-- Start SDK Example Usage [usage] -->
### Sending single SMS Message

This example demonstrates simple sending SMS message to a single recipient:

```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Outgoing.Sms.Send(ctx, operations.CreateSendSmsRequestBodyArrayOfSmsMessage(
		[]components.SmsMessage{
			components.SmsMessage{
				Recipients: components.CreateRecipientsArrayOfStr(
					[]string{
						"+48999999999",
					},
				),
				Message: "To jest treść wiadomości",
				Sender:  messagingsdkgo.String("Bramka SMS"),
				Type:    components.SmsTypeSmsPro.ToPointer(),
				Unicode: messagingsdkgo.Bool(true),
				Flash:   messagingsdkgo.Bool(false),
				Date:    nil,
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
<!-- End SDK Example Usage [usage] -->