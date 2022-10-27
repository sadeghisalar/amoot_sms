# AmootSMS REST API
READ [AmootSMS Docs](https://doc.amootsms.com/documentation/api/webservice2asmx/accountstatus) !IMPORTANT

## Installation
```go
go get github.com/sadeghisalar/amoot_sms
```

## 1 - Create Instance
```go
	amootSMS := amootsms.Api{
		Username: "09000000000",
		Password: "09000000000",
	}
	data := amootSMS.AccountStatus()
	fmt.Println(data["AccountName"])
```

## 2 - Use Default Methods
```go

	// Send Simple Message
	amootSMS.SendSimple("Test ...","public","", []string{
		"09380000000",
		"09320000000",
	})
	
	// Send With Pattern
	amootSMS.SendWithPattern("490","123456","09380000000")

	// Send Quick OTP
	amootSMS.SendQuickOTP("1","123456","09380000000")

```

## 3 - OR, Make a Call!
```go
	amootSMS.Call("/SendSimple", map[string]interface{}{
		"SendDateTime":   time.Now().Format(time.RFC3339),
		"SMSMessageText": "Test ...",
		"LineNumber":     "public",
		"Mobiles":        "09380000000,09320000000",
	})
```
