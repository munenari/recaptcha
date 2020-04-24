
# ReCAPTCHA client

## usage

```go
req := recaptcha.Request{
	Secret:   "dummy-secret",
	Response: "dummy-response",
	RemoteIP: "127.0.0.1",
}
res, err := req.Verify()
if err != nil {
	panic(err)
}
if res.Success {
	fmt.Println("success!")
}
```
