package recaptcha

import (
	"net/http"
	"testing"
)

func TestVerify(t *testing.T) {
	c := RecaptchaClient{Client: http.DefaultClient, Secret: "dummy-secret"}
	res, err := c.Verify(t.Context(), "dummy-response", "127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	_ = res
	// t.Error(res)
}
