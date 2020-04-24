package recaptcha

import "testing"

func TestVerify(t *testing.T) {
	req := Request{
		Secret:   "dummy-secret",
		Response: "dummy-response",
		RemoteIP: "127.0.0.1",
	}
	res, err := req.Verify()
	if err != nil {
		t.Fatal(err)
	}
	_ = res
	// t.Error(res)
}
