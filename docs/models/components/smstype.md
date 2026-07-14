# SmsType

SMS type (components.SmsTypeSmsPro -> SMS PRO, components.SmsTypeSmsEco -> SMS ECO, components.SmsTypeSmsTwoWay -> SMS 2WAY)

## Example Usage

```go
import (
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
)

value := components.SmsTypeSmsPro

// Open enum: custom values can be created with a direct type cast
custom := components.SmsType(999)
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SmsTypeSmsPro`    | 1                  |
| `SmsTypeSmsEco`    | 3                  |
| `SmsTypeSmsTwoWay` | 4                  |