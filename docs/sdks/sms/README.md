# Sms
(*Outgoing.Sms*)

## Overview

### Available Operations

* [GetPrice](#getprice) - Check the price of SMS Messages
* [Send](#send) - Send SMS Messages

## GetPrice

Check the price of single or multiple SMS messages at the same time before sending them. You can pass a single `SmsMessage` object using `operations.CreateGetSmsPriceRequestBodySmsMessage()` method (for single message) or `[]SmsMessage` array using `operations.CreateGetSmsPriceRequestBodyArrayOfSmsMessage()` method (for multiple messages). Each `SmsMessage` object has several properties, describing message parameters such as recipient phone number, content of the message, type, etc.
The method will accept maximum **100** messages in one call.

As a successful result a `GetSmsPriceResponse` object will be returned with `Prices` property of type `[]Price` containing a `Price` objects, one object per each single message. You should check the `Error` property of each `Price` object to make sure which messages were priced successfully and which finished with an error. Successfully priced messages will have `null` value of `Error` property.

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

    res, err := s.Outgoing.Sms.GetPrice(ctx, operations.CreateGetSmsPriceRequestBodyArrayOfSmsMessage(
        []components.SmsMessage{
            components.SmsMessage{
                Recipients: components.CreateSmsMessageRecipientsPhoneNumberWithCid(
                    components.PhoneNumberWithCid{
                        Nr: "+48999999999",
                        Cid: messagingsdkgo.String("my-id-1113"),
                    },
                ),
                Message: "To jest treść wiadomości",
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
| `request`                                                                              | [operations.GetSmsPriceRequestBody](../../models/operations/getsmspricerequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetSmsPriceResponse](../../models/operations/getsmspriceresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 4XX            | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## Send

Send single or multiple SMS messages at the same time. You can pass a single `SmsMessage` object using `operations.CreateSendSmsRequestBodySmsMessage()` method (for single message) or `[]SmsMessage` array using `operations.CreateSendSmsRequestBodyArrayOfSmsMessage()` method (for multiple messages). Each `SmsMessage` object has several properties, describing message parameters such recipient phone number, content of the message, type or scheduled sending date, etc. This method will accept maximum 100 messages in one call.

As a successful result a `SendSmsResponse` object will be returned with `Messages` property of type `[]Message` containing `Message` objects, one object per each single message. You should check the `StatusCode` property of each `Message` object to make sure which messages were accepted by gateway (queued) and which were rejected. In case of rejection, `StatusDescription` property will include a reason.

`SendSmsResponse` will also include `Headers` property with `X-Success-Count` (a count of messages which were processed successfully), `X-Error-Count` (count of messages which were rejected) and `X-Sandbox` (if a request was made in Sandbox or Production system) elements.

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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.SendSmsRequestBody](../../models/operations/sendsmsrequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.SendSmsResponse](../../models/operations/sendsmsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |