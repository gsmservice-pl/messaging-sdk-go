# Incoming
(*Incoming*)

## Overview

### Available Operations

* [List](#list) - List the received SMS messages
* [GetByIds](#getbyids) - Get the incoming messages by IDs

## List

Get the details of all of received messages from your account incoming messages box. This endpoint supports pagination so you have to pass as query parameters a `page` value (number of page with received messages which you want to access) and a `limit` value (max of received messages per page). Messages are fetched from the latest one. The system will accept maximum **50** as `limit` parameter value. If you need to get details of larger volume of messages, please access them with next pages.
    
As a successful result an array with `IncomingMessage` objects will be returned, each object per single received message. Response will also include meta-data headers: `X-Total-Results` (a total count of all received messages which are available in incoming box on your account), `X-Total-Pages` (a total number of all pages with results), `X-Current-Page` (A current page number) and `X-Limit` (messages count per single page). This request have to be authenticated using **API Access Token**. 

A response contains also a special `Link` header which includes *URIs* to access next, previous, first and last page with received messages (which complies with [RFC 5988](https://www.rfc-editor.org/rfc/rfc5988)).

In case of an error, the `ErrorResponse` object will be returned with proper HTTP header status code (our error response complies with [RFC 9457](https://www.rfc-editor.org/rfc/rfc7807)).

### Example Usage

```go
package main

import(
	"os"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"context"
	"log"
)

func main() {
    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity(os.Getenv("CLIENT_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Incoming.List(ctx, messagingsdkgo.Int64(1), messagingsdkgo.Int64(10))
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 4XX, 5XX | application/problem+json     |

## GetByIds

Get the details of one or more received messages using their `ids`. You have to pass the unique incoming message *IDs* as path parameter, which were given while receiving a messages. If you want to get the details of multiple messages at once, please separate their IDs with a comma. The system will accept maximum 50 identifiers in one call. If you need to get details of larger volume of incoming messages, please split it to several separate requests.
    
As a successful result an array with `IncomingMessage` objects will be returned, each object per single found message. Response will also include meta-data headers: `X-Success-Count` (a count of incoming messages which were found and returned correctly) and `X-Error-Count` (count of incoming messages which were not found).

If you pass duplicated IDs, API will return data of this message only once. This request have to be authenticated using **API Access Token**. 

In case of an error, the `ErrorResponse` object will be returned with proper HTTP header status code (our error response complies with [RFC 9457](https://www.rfc-editor.org/rfc/rfc7807)).

### Example Usage

```go
package main

import(
	"os"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"context"
	"log"
)

func main() {
    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity(os.Getenv("CLIENT_BEARER")),
    )

    ctx := context.Background()
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `ids`                                                                                                                    | []*int64*                                                                                                                | :heavy_check_mark:                                                                                                       | Message IDs assigned by the system (separated by comma). The system will accept a maximum of 50 identifiers in one call. |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetIncomingMessagesResponse](../../models/operations/getincomingmessagesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 404, 4XX, 5XX  | application/problem+json |