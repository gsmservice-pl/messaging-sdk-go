<!-- Start SDK Example Usage [usage] -->
### Sending single SMS Message

This example demonstrates simple sending SMS message to a single recipient:

```go
package main

import (
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v4"
	"github.com/gsmservice-pl/messaging-sdk-go/v4/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v4/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := v4.New(
		v4.WithSecurity("<YOUR API ACCESS TOKEN>"),
	)

	res, err := s.Outgoing.Sms.Send(ctx, operations.CreateSendSmsRequestBodySmsMessage(
		components.SmsMessage{
			Recipients: components.CreateSmsMessageRecipientsStr(
				"+48999999999",
			),
			Message: "This is SMS message content.",
			Unicode: v4.Pointer(true),
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
	"github.com/gsmservice-pl/messaging-sdk-go/v4"
	"github.com/gsmservice-pl/messaging-sdk-go/v4/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v4/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := v4.New(
		v4.WithSecurity("<YOUR API ACCESS TOKEN>"),
	)

	res, err := s.Outgoing.Mms.Send(ctx, operations.CreateSendMmsRequestBodyMmsMessage(
		components.MmsMessage{
			Recipients: components.CreateRecipientsStr(
				"+48999999999",
			),
			Subject: v4.Pointer("This is a subject of the message"),
			Message: v4.Pointer("This is MMS message content."),
			Attachments: v4.Pointer(components.CreateAttachmentsStr(
				"R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7",
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