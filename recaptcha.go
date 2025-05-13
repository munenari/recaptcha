package recaptcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const recaptchaEndpoint = "https://www.google.com/recaptcha/api/siteverify"

type (
	RecaptchaClient struct {
		Client *http.Client
		Secret string
	}
	Response struct {
		Success     bool      `json:"success"`
		Score       float64   `json:"score"`
		Action      string    `json:"action"`
		ChallengeTS time.Time `json:"challenge_ts"`
		Hostname    string    `json:"hostname"`
		ErrorCodes  []string  `json:"error-codes"`
	}

	request struct {
		Secret   string `_form:"secret"`
		Response string `_form:"response"`
		RemoteIP string `_form:"remoteip"`
	}
)

func (me RecaptchaClient) Verify(ctx context.Context, response, remoteIP string) (*Response, error) {
	req := &request{
		Secret:   me.Secret,
		Response: response,
		RemoteIP: remoteIP,
	}
	client := me.Client
	if client == nil {
		client = http.DefaultClient
	}
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, recaptchaEndpoint, strings.NewReader(req.Values().Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := &Response{}
	err = json.NewDecoder(resp.Body).Decode(res)
	return res, err
}

func (x *request) Values() url.Values {
	v := url.Values{}
	v.Add("secret", x.Secret)
	v.Add("response", x.Response)
	if x.RemoteIP != "" {
		v.Add("remoteip", x.RemoteIP)
	}
	return v
}
