package skeleton

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var url string

func init() {
	os.Setenv("url", "https://esme-runner-ct5vendzhq-uc.a.run.app")

	url = os.Getenv("url")
}

type user struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func Test_getAuthBasic_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBasic_success(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 2)
}

func Test_postAuthBasic_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPost,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_postAuthBasic_noBody_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPost,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_postAuthBasic_success(t *testing.T) {
	user := user{
		FirstName: "foo",
		LastName:  "bar",
	}

	userBytes, _ := json.Marshal(user)

	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPost,
		Body:    userBytes,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, data[0]["firstName"], "foo")
	assert.Equal(t, data[0]["lastName"], "bar")
	assert.Equal(t, data[0]["id"], "3")
}

func Test_putAuthBasic_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_putAuthBasic_noBody_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_putAuthBasic_invalidCredentials_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "passR",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_putAuthBasic_success(t *testing.T) {
	user := user{
		FirstName: "Foo",
		LastName:  "Bar",
	}

	userBytes, _ := json.Marshal(user)

	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user", url),
		Method:  http.MethodPut,
		Body:    userBytes,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, data[0]["firstName"], "Foo")
	assert.Equal(t, data[0]["lastName"], "Bar")
	assert.Equal(t, data[0]["id"], "3")
}

func Test_deleteAuthBasic_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user/1", url),
		Method:  http.MethodDelete,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_deleteAuthBasic_success(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/basic/user/1", url),
		Method:  http.MethodDelete,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Basic: &AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_getAuthBearerToken_missingToken_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/bearer_token/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBearerToken_invalidToken_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/bearer_token/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth:    &Auth{BearerToken: &AuthBearerToken{Token: "t"}},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBearerToken_success(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/bearer_token/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			BearerToken: &AuthBearerToken{
				Token: "token",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_getAuthCustomHeader_missingHeader_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/custom/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_getAuthCustomHeader_invalidHeader_failure(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/custom/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{Custom: map[string]string{
			"my_custom_header": "val",
		}},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_getAuthCustomHeader_success(t *testing.T) {
	r := Request{
		Url:     fmt.Sprintf("%s/auth/custom/users", url),
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &Auth{
			Custom: map[string]string{
				"my_custom_header": "header val",
			},
		},
	}

	resp, err := Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
