# Senders
(*Senders*)

## Overview

### Available Operations

* [List](#list) - List allowed senders names
* [Add](#add) - Add a new sender name
* [Delete](#delete) - Delete a sender name
* [SetDefault](#setdefault) - Set default sender name

## List

Get a list of allowed senders defined in your account.

As a successful result a `ListSendersResponse` object will be returned witch `Senders` property of type `[]Sender` array containing `Sender` objects, each object per single sender.

### Example Usage

<!-- UsageSnippet language="go" operationID="listSenders" method="get" path="/senders" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

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
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## Add

Define a new allowed sender on your account. You should pass as parameter a `SenderInput` struct with two fields: `Sender` (defines sender name) and `Description`. Please carefully fill this field with the extensive description of a sender name (what will be its use, what the name mean, etc).

As a successful result a `AddSenderResponse` object will be returned with a `Sender` property containing a `Sender` object with details and status of added sender name.

### Example Usage

<!-- UsageSnippet language="go" operationID="addSender" method="post" path="/senders" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

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
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## Delete

Removes defined sender name from your account. This method accepts a `string` type with a **sender name** you want to remove. Sender name will be deleted immediately.

As a successful response there would be `DeleteSenderResponse` object returned with no Exception thrown.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSender" method="delete" path="/senders/{sender}" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 401, 403, 404, 4XX  | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## SetDefault

Set default sender name to one of the senders names already defined on your account. This method accepts a `string` type containing a **sender name** to be set as default on your account.

As a successful response a `SetDefaultSenderResponse` object will be returned no Exception to be thrown.

### Example Usage

<!-- UsageSnippet language="go" operationID="setDefaultSender" method="patch" path="/senders/{sender}" -->
```go
package main

import(
	"context"
	"github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

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
| sdkerrors.ErrorResponse  | 404                      | application/json         |
| sdkerrors.ErrorResponse  | 400, 401, 403, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |