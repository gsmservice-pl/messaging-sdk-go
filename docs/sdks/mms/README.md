# Mms
(*Outgoing.Mms*)

## Overview

### Available Operations

* [GetPrice](#getprice) - Check the price of MMS Messages
* [Send](#send) - Send MMS Messages

## GetPrice

Check the price of single or multiple MMS messages at the same time before sending them. You can pass a single `MmsMessage` object using `operations.CreateGetMmsPriceRequestBodyMmsMessage()` method (for single message) or `[]MmsMessage` array using `operations.CreateGetMmsPriceRequestBodyArrayOfMmsMessage()` method (for multiple messages). Each `MmsMessage` object has several properties, describing message parameters such as recipient phone number, content of the message, attachments, etc. 
The system will accept maximum **50** messages in one call.

As a successful result a `GetMmsPriceResponse` object will be returned with `Prices` property of type `[]Price` containing a `Price` objects, one object per each single message. You should check the `Error` property of each `Price` object to make sure which messages were priced successfully and which finished with an error. Successfully priced messages will have `null` value of `Error` property.

`GetSmsPriceResponse` object will include also `Headers` property with `X-Success-Count` (a count of messages which were processed successfully) and `X-Error-Count` (count of messages which were rejected) elements.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMmsPrice" method="post" path="/messages/mms/price" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Outgoing.Mms.GetPrice(ctx, operations.CreateGetMmsPriceRequestBodyMmsMessage(
        components.MmsMessage{
            Recipients: components.CreateRecipientsStr(
                "+48999999999",
            ),
            Subject: v3.Pointer("This is a subject of the message"),
            Message: v3.Pointer("This is MMS message content."),
            Attachments: v3.Pointer(components.CreateAttachmentsStr(
                "<file body in base64 format>",
            )),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Prices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetMmsPriceRequestBody](../../models/operations/getmmspricerequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetMmsPriceResponse](../../models/operations/getmmspriceresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 4XX            | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## Send

Send single or multiple MMS messages at the same time. You can pass a single `MmsMessage` object using `operations.CreateSendMmsRequestBodyMmsMessage()` method (for single message) or `[]MmsMessage` array using `operations.CreateSendMmsRequestBodyArrayOfMmsMessage()` method (for multiple messages). Each `MmsMessage` object has several properties, describing message parameters such recipient phone number, content of the message, attachments or scheduled sending date, etc. This method will accept maximum 50 messages in one call.

As a successful result a `SendMmsResponse` object will be returned with `Messages` property of type `[]Message` containing `Message` objects, one object per each single message. You should check the `StatusCode` property of each `Message` object to make sure which messages were accepted by gateway (queued) and which were rejected. In case of rejection, `StatusDescription` property will include a reason.

`SendMmsResponse` will also include `Headers` property with `X-Success-Count` (a count of messages which were processed successfully), `X-Error-Count` (count of messages which were rejected) and `X-Sandbox` (if a request was made in Sandbox or Production system) elements.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendMms" method="post" path="/messages/mms" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Outgoing.Mms.Send(ctx, operations.CreateSendMmsRequestBodyMmsMessage(
        components.MmsMessage{
            Recipients: components.CreateRecipientsStr(
                "+48999999999",
            ),
            Subject: v3.Pointer("This is a subject of the message"),
            Message: v3.Pointer("This is MMS message content."),
            Attachments: v3.Pointer(components.CreateAttachmentsStr(
                "<file body in base64 format>",
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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.SendMmsRequestBody](../../models/operations/sendmmsrequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.SendMmsResponse](../../models/operations/sendmmsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |