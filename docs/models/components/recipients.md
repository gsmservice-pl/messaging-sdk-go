# Recipients

The recipient number or multiple recipients numbers of single message. To set one recipient, simply pass here a `string` with his phone number. To set multiple recipients, pass here a simple `array` of `string`. Optionally you can also set custom id (user identifier) for each message - pass `PhoneNumberWithCid` object (in case of single recipient) or `Array` of `PhoneNumberWithCid` (in case of multiple recipients).


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

