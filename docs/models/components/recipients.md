# Recipients

The recipient number or multiple recipients numbers of single message. To set one recipient, simply use `components.CreateRecipientsStr()` method simply passing to it a `string` with his phone number. To set multiple recipients, please use `components.CreateRecipientsArrayOfStr()` method passing to it `[]string` with recipients.\r\n\r\nOptionally you can also set custom id (user identifier) for each message - use `components.CreateRecipientsPhoneNumberWithCid()` method passing `PhoneNumberWithCid` struct (in case of single recipient) or `operations.CreateRecipientsArrayOfPhoneNumberWithCid()` method passing `[]PhoneNumberWithCid` (in case of multiple recipients).


## Supported Types

### 

```go
recipients := components.CreateRecipientsStr(string{/* values here */})
```

### 

```go
recipients := components.CreateRecipientsArrayOfStr([]string{/* values here */})
```

### PhoneNumberWithCid

```go
recipients := components.CreateRecipientsPhoneNumberWithCid(components.PhoneNumberWithCid{/* values here */})
```

### 

```go
recipients := components.CreateRecipientsArrayOfPhoneNumberWithCid([]components.PhoneNumberWithCid{/* values here */})
```

