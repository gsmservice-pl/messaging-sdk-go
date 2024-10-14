# Sms
(*Outgoing.Sms*)

## Overview

### Available Operations

* [GetPrice](#getprice) - Check the price of SMS Messages
* [Send](#send) - Send SMS Messages

## GetPrice

Check the price of single or multiple SMS messages at the same time before sending them. You have to pass as request body the `Sms` object (for single message) or `array` of `Sms` objects (for multiple messages). Each object has several properties, describing message parameters such recipient phone number, content of the message, type, etc. Please mind that some of them are required.
The system will accept maximum **100** messages in one call. If you need to check the price of larger volume of messages, please split it to several separate requests.

As a successful result an `array` of `Price` objects will be returned, one object per each single message. You should check the `error` property of each message in a response body to make sure which were priced successfully and which finished with an error. Successfully priced messages will have `null` value of `error` property. Response will also include meta-data headers: `X-Success-Count` (a count of messages which were processed successfully) and `X-Error-Count` (count of messages which were rejected).

If you send duplicated messages in one call, API will process such message only once. This request have to be authenticated using **API Access Token**.

In case of an error, the `ErrorResponse` object will be returned with proper HTTP header status code (our error response complies with [RFC 9457](https://www.rfc-editor.org/rfc/rfc7807)).


### Example Usage

```go
package main

import(
	"os"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
	"log"
)

func main() {
    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity(os.Getenv("CLIENT_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Outgoing.Sms.GetPrice(ctx, operations.CreateGetSmsPriceRequestBodyArrayOfSms(
        []components.Sms{
            components.Sms{
                Recipients: components.CreateRecipientsPhoneNumberWithCid(
                    components.PhoneNumberWithCid{
                        Nr: "+48999999999",
                        Cid: messagingsdkgo.String("my-id-1113"),
                    },
                ),
                Message: "To jest treść wiadomości",
                Sender: messagingsdkgo.String("Bramka SMS"),
                Type: components.SmsTypeSmsPro.ToPointer(),
                Unicode: messagingsdkgo.Bool(true),
                Flash: messagingsdkgo.Bool(false),
                Date: nil,
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
| sdkerrors.ErrorResponse  | 400, 401, 4XX, 5XX       | application/problem+json |

## Send

Send single or multiple SMS messages at the same time. You have to pass as request body the `Sms` object (for single message) or `array` of `Sms` objects (for multiple messages). Each object has several properties, describing message parameters such recipient phone number, content of the message, type or scheduled sending date, etc. Please mind that some of them are required.
The system will accept maximum 100 messages in one call. If you need to send larger volume of messages, please split it to several separate requests.

As a successful result an `array` with `Message` objects will be returned, one object per each single message. You should check the `status_code` property of each message in a response body to make sure which were accepted by gateway (queued) and which were rejected. In case of rejection, `status_description` property will include a reason. Response will also include meta-data headers: `X-Success-Count` (a count of messages which were processed successfully), `X-Error-Count` (count of messages which were rejected) and `X-Sandbox` (if a request was made in Sandbox or Production system).

If you send duplicated messages in one call, API will process such message only once. This request have to be authenticated using **API Access Token**.

In case of an error, the `ErrorResponse` object will be returned with proper HTTP header status code (our error response complies with [RFC 9457](https://www.rfc-editor.org/rfc/rfc7807)).

### Example Usage

```go
package main

import(
	"os"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
	"log"
)

func main() {
    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity(os.Getenv("CLIENT_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Outgoing.Sms.Send(ctx, operations.CreateSendSmsRequestBodyArrayOfSms(
        []components.Sms{
            components.Sms{
                Recipients: components.CreateRecipientsArrayOfStr(
                    []string{
                        "+48999999999",
                    },
                ),
                Message: "To jest treść wiadomości",
                Sender: messagingsdkgo.String("Bramka SMS"),
                Type: components.SmsTypeSmsPro.ToPointer(),
                Unicode: messagingsdkgo.Bool(true),
                Flash: messagingsdkgo.Bool(false),
                Date: nil,
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
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX, 5XX  | application/problem+json |