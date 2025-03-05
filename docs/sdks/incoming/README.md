# Incoming
(*Incoming*)

## Overview

### Available Operations

* [List](#list) - List the received SMS messages
* [GetByIds](#getbyids) - Get the incoming messages by IDs

## List

Get the details of all received messages from your account incoming messages box. This method supports pagination so you have to pass `page` (number of page with received messages which you want to access) and a `limit` (max of received messages per page) parameters. Messages are fetched from the latest one. This method will accept maximum **50** as `limit` parameter value.

As a successful result a `ListIncomingMessagesResponse` object will be returned with `IncomingMessages` property of type `[]IncomingMessage` containing `IncomingMessage` objects, each object per single received message.

`ListIncomingMessagesResponse` object will contain also a `Headers` property where you can find `X-Total-Results` (a total count of all received messages which are available in incoming box on your account), `X-Total-Pages` (a total number of all pages with results), `X-Current-Page` (A current page number) and `X-Limit` (messages count per single page) elements.

### Example Usage

```go
package main

import(
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Incoming.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncomingMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number of results                                   | 1                                                        |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Number of results on one page                            | 10                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListIncomingMessagesResponse](../../models/operations/listincomingmessagesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 404, 4XX  | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## GetByIds

Get the details of one or more received messages using their `ids`. This method accepts an array of type `[]int64` containing unique incoming message *IDs*, which were given while receiving a messages. The method will accept maximum 50 identifiers in one call.

As a successful result a `GetIncomingMessagesResponse` object will be returned with an `IncomingMessages` property of type `[]IncomingMessage` containing `IncomingMessage` objects, each object per single received message.

`GetIncomingMessagesResponse` object will contain also a `Headers` property where you can find `X-Success-Count` (a count of incoming messages which were found and returned correctly) and `X-Error-Count` (count of incoming messages which were not found) elements.

### Example Usage

```go
package main

import(
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Incoming.GetByIds(ctx, []int64{
        43456,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncomingMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `ids`                                                                                                        | []*int64*                                                                                                    | :heavy_check_mark:                                                                                           | Array of Message IDs assigned by the system. The system will accept a maximum of 50 identifiers in one call. |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetIncomingMessagesResponse](../../models/operations/getincomingmessagesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 404, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |