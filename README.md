# Skeleton

[![Go Report Card](https://goreportcard.com/badge/github.com/stkr89/skeleton)](https://goreportcard.com/report/github.com/stkr89/skeleton)

Skeleton is a dead simple and beginner friendly http client with support for 
`GET`, `POST`, `PUT`, `DELETE` methods. It supports following authentication types:

- Basic
- Bearer Token
- Custom Headers

## Usage

Http `GET` request with basic authentication

```go
req := Request{
        Url:     "http://localhost:8080/auth/basic/user",
        Method:  http.MethodGet,
        Timeout: 10,
        Auth: &Auth{
            Basic: &AuthBasic{
                Username: "username",
                Password: "passR",
            },
        },
    }
```

Http `POST` request with bearer token authentication

```go
req := Request{
        Url:     "http://localhost:8080/auth/bearer_token/users",
        Method:  http.MethodPost,
        Body: map[string]string{
            "firstName": "foo",
            "lastName":  "bar",
        },
        Timeout: 10,
        Auth:    &Auth{
            BearerToken: &AuthBearerToken{
                Token: "token",
            },
        },
    }
```

Http `GET` request with custom authentication 

```go
req := Request{
        Url:     "http://localhost:8080/auth/custom/users",
        Method:  http.MethodGet,
        Timeout: 10,
        Auth: &Auth{
            Custom: map[string]string{
                "my_custom_header": "header val",
            },
        },
    }
```

Make the http request using above request object as follows

```go
resp, err := send(&req)
if err != nil {
    fmt.println(err)
}

var data []map[string]interface{}
err = json.Unmarshal(respBytes, &data)
if err != nil {
    fmt.println(err)
}

fmt.println(data)
```