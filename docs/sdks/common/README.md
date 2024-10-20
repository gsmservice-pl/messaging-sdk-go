# Common
(*Common*)

## Overview

### Available Operations

* [Ping](#ping) - Checks API availability and version

## Ping

Check the API connection and the current API availability status. Also you will get the current API version number.

As a successful result a `PingResponse` object will be returned.

### Example Usage

```go
package main

import(
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"context"
	"log"
)

func main() {
    s := messagingsdkgo.New()

    ctx := context.Background()
    res, err := s.Common.Ping(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.PingResponse != nil {
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

**[*operations.PingResponse](../../models/operations/pingresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 400, 4XX, 503, 5XX       | application/problem+json |