package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const recaptchaEndpoint = "https://www.google.com/recaptcha/api/siteverify"

var client = &http.Client{
	Transport: http.DefaultTransport,
}

type (
	// Response struct
	Response struct {
		Success     bool      `json:"success"`
		ChallengeTS time.Time `json:"challenge_ts"`
		Hostname    string    `json:"hostname"`
		Score       float64   `json:"score"`
		ErrorCodes  []string  `json:"error-codes"`
	}

	// Request struct
	Request struct {
		Secret   string `_form:"secret"`
		Response string `_form:"response"`
		RemoteIP string `_form:"remoteip"`
	}
)

// Verify recaptcha token
func (x *Request) Verify() (*Response, error) {
	resp, err := client.PostForm(recaptchaEndpoint, x.Values())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(Response)
	err = json.NewDecoder(resp.Body).Decode(res)
	return res, err
}

// Values returns url.Values for recaptcha struct
func (x *Request) Values() url.Values {
	v := url.Values{}
	v.Add("secret", x.Secret)
	v.Add("response", x.Response)
	if x.RemoteIP != "" {
		v.Add("remoteip", x.RemoteIP)
	}
	return v
}
