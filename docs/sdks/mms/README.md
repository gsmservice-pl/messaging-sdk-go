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

```go
package main

import(
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

    res, err := s.Outgoing.Mms.GetPrice(ctx, operations.CreateGetMmsPriceRequestBodyArrayOfMmsMessage(
        []components.MmsMessage{
            components.MmsMessage{
                Recipients: components.CreateRecipientsArrayOfStr(
                    []string{
                        "+48999999999",
                    },
                ),
                Subject: messagingsdkgo.String("To jest temat wiadomości"),
                Message: messagingsdkgo.String("To jest treść wiadomości"),
                Attachments: messagingsdkgo.Pointer(components.CreateAttachmentsStr(
                    "<file_body in base64 format>",
                )),
            },
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

```go
package main

import(
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
                    Nr: "+48999999999",
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