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
	// ReCaptchaResponse struct
	ReCaptchaResponse struct {
		Success     bool      `json:"success"`
		ChallengeTS time.Time `json:"challenge_ts"`
		Hostname    string    `json:"hostname"`
		Score       float64   `json:"score"`
		ErrorCodes  []string  `json:"error-codes"`
	}

	// ReCaptchaRequest struct
	ReCaptchaRequest struct {
		Secret   string `_form:"secret"`
		Response string `_form:"response"`
		RemoteIP string `_form:"remoteip"`
	}
)

// Verify recaptcha token
func (x *ReCaptchaRequest) Verify() (*ReCaptchaResponse, error) {
	resp, err := client.PostForm(recaptchaEndpoint, x.Values())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(ReCaptchaResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	return res, err
}

// Values returns url.Values for recaptcha struct
func (x *ReCaptchaRequest) Values() url.Values {
	v := url.Values{}
	v.Add("secret", x.Secret)
	v.Add("response", x.Response)
	if x.RemoteIP != "" {
		v.Add("remoteip", x.RemoteIP)
	}
	return v
}
