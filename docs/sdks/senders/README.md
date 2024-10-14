# Senders
(*Senders*)

## Overview

### Available Operations

* [List](#list) - List allowed senders names
* [Add](#add) - Add a new sender name
* [Delete](#delete) - Delete a sender name
* [SetDefault](#setdefault) - Set default sender name

## List

Get a list of allowed senders defined in your account. The request doesn't contain a body or any parameters. 
        
As a successful result an array with `Sender` objects will be returned, each object per single sender. Senders are being registered by providers and operators. Registered senders get *Active* status and can be used then to send messages. *Pending* senders are also returned by API (with proper `status`) but until registration they cannot be used. This request have to be authenticated using **API Access Token**.

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
        messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Senders.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Senders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListSendersResponse](../../models/operations/listsendersresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX, 5XX  | application/problem+json |

## Add

Define a new allowed sender on your account. The request body should contain a `Sender` object with two properties: `sender` (defines sender name) and `description`. The secont parameter is very important - sender names are being registered by providers and operators. Only fully registered sender names can be used to send messages. Providers need sometimes detailed description of case in which the sender will be used to eliminate frauds. After verifing it they make a decisions if such sender name can be registered. Please carefully fill this property with the extensive description of a sender name (what will be its use, what the name mean, etc). 
        
As a successful result a single `Sender` object will be returned. Registered senders get *Active* status and can be used then to send messages. Pending Senders are also returned by API (with proper `status`) but until registration they cannot be used. Response will also include meta-data header: `X-Sandbox` (if a request was made in Sandbox or Production system). This request have to be authenticated using **API Access Token**.

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
        messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Senders.Add(ctx, components.SenderInput{
        Sender: "Bramka SMS",
        Description: "This is our company name. It contains our registered trademark.",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sender != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.SenderInput](../../models/components/senderinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.AddSenderResponse](../../models/operations/addsenderresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX, 5XX  | application/problem+json |

## Delete

Removes defined sender name from your account. This endpoint accepts a path `sender` parameter with empty request body. You should pass the full sender name to delete it. Sender name will be deleted immediately.

As a successful response only HTTP status code of *204* will be returned in header with empty response body. Response will also include meta-data header: `X-Sandbox` (if a request was made in Sandbox or Production system).
This request have to be authenticated using **API Access Token**.

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
        messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Senders.Delete(ctx, "Podpis")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `sender`                                                 | *string*                                                 | :heavy_check_mark:                                       | Sender name to be removed                                | Podpis                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteSenderResponse](../../models/operations/deletesenderresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorResponse      | 400, 401, 403, 404, 4XX, 5XX | application/problem+json     |

## SetDefault

Set default sender name to one of the senders names already defined on your account. Default sender name can be used while sending messages when you not pass any other defined sender to `Sms` object while sending message. 

This endpoint accepts a path `sender` parameter with empty request body. You should pass the full sender name to set it as default on your account.

As a successful response only HTTP status code of *204* will be returned in header with empty response body. Response will also include meta-data header: `X-Sandbox` (if a request was made in Sandbox or Production system).
This request have to be authenticated using **API Access Token**.

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
        messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
    )

    ctx := context.Background()
    res, err := s.Senders.SetDefault(ctx, "Podpis")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `sender`                                                 | *string*                                                 | :heavy_check_mark:                                       | Sender name to set as default                            | Podpis                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SetDefaultSenderResponse](../../models/operations/setdefaultsenderresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX, 5XX  | application/problem+json |
| sdkerrors.ErrorResponse  | 404                      | application/json         |