package skeleton

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

func send(r *Request) (*http.Response, error) {
	if r.Body == nil {
		r.Body = make(map[string]string)
	}

	jsonBytes, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if r.Auth != nil {
		addAuthorizationHeaders(r, req)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(r.Timeout)}

	return client.Do(req)
}

func addAuthorizationHeaders(r *Request, req *http.Request) {
	if r.Auth.Basic != nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(r.Auth.Basic.Username, r.Auth.Basic.Password))
	} else if r.Auth.BearerToken != nil {
		req.Header.Add("Authorization", "Bearer "+r.Auth.BearerToken.Token)
	} else if r.Auth.Custom != nil {
		for k, v := range r.Auth.Custom {
			req.Header.Add(k, v)
		}
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
