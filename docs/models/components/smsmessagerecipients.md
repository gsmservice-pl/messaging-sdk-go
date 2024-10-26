# SmsMessageRecipients

The recipient number or multiple recipients numbers of single message. To set one recipient, please use `components.CreateSmsMessageRecipientsStr()` method simply passing to it a `string` with his phone number. To set multiple recipients, please use `components.CreateSmsMessageRecipientsArrayOfStr()` method passing to it `[]string` with recipients.\r\n\r\nOptionally you can also set custom id (user identifier) for each message - use `components.CreateSmsMessageRecipientsPhoneNumberWithCid()` method passing `PhoneNumberWithCid` struct (in case of single recipient) or `operations.CreateSmsMessageRecipientsArrayOfPhoneNumberWithCid()` method passing []PhoneNumberWithCid (in case of multiple recipients).


## Supported Types

### 

```go
smsMessageRecipients := components.CreateSmsMessageRecipientsStr(string{/* values here */})
```

### 

```go
smsMessageRecipients := components.CreateSmsMessageRecipientsArrayOfStr([]string{/* values here */})
```

### PhoneNumberWithCid

```go
smsMessageRecipients := components.CreateSmsMessageRecipientsPhoneNumberWithCid(components.PhoneNumberWithCid{/* values here */})
```

### 

```go
smsMessageRecipients := components.CreateSmsMessageRecipientsArrayOfPhoneNumberWithCid([]components.PhoneNumberWithCid{/* values here */})
```

