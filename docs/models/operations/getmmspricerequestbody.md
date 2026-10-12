# GetMmsPriceRequestBody

To check the price of a single message or messages with the same content to multiple recipients, pass a single `MmsMessage` object with the properties of this message using `operations.CreateGetMmsPriceRequestBodyMmsMessage()` method. To check the price of multiple messages with different content at the same time, pass a `[]MmsMessage` array with the properties of each message using `operations.CreateGetMmsPriceRequestBodyArrayOfMmsMessage()` method.


## Supported Types

### MmsMessage

```go
getMmsPriceRequestBody := operations.CreateGetMmsPriceRequestBodyMmsMessage(components.MmsMessage{/* values here */})
```

### 

```go
getMmsPriceRequestBody := operations.CreateGetMmsPriceRequestBodyArrayOfMmsMessage([]components.MmsMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getMmsPriceRequestBody.Type {
	case operations.GetMmsPriceRequestBodyTypeMmsMessage:
		// getMmsPriceRequestBody.MmsMessage is populated
	case operations.GetMmsPriceRequestBodyTypeArrayOfMmsMessage:
		// getMmsPriceRequestBody.ArrayOfMmsMessage is populated
}
```
