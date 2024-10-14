# Common
(*Common*)

## Overview

This section describes other usefull operations and tools

### Available Operations

* [Ping](#ping) - Checks API availability and version

## Ping

Check the API connection and the current API availability status. Also you will get the current API version number. The request doesn't contain a body or any parameters.

As a successful result a `PingResponse` object will be returned. This request doesn't need to be authenticated.

In case of an error, the `ErrorResponse` object will be returned with proper HTTP header status code (our error response complies with [RFC 9457](https://www.rfc-editor.org/rfc/rfc7807)).

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