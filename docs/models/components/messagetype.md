# MessageType

Message type (components.MessageTypeSmsPro -> SMS PRO, components.MessageTypeSmsEco -> SMS ECO, components.MessageTypeSmsTwoWay ->SMS 2WAY, components.MessageTypeMms -> MMS)

## Example Usage

```go
import (
	"github.com/gsmservice-pl/messaging-sdk-go/v4/models/components"
)

value := components.MessageTypeSmsPro

// Open enum: custom values can be created with a direct type cast
custom := components.MessageType(999)
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `MessageTypeSmsPro`    | 1                      |
| `MessageTypeSmsEco`    | 3                      |
| `MessageTypeSmsTwoWay` | 4                      |
| `MessageTypeMms`       | 10                     |