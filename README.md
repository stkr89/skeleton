# Skeleton

Skeleton is a dead simple and beginner friendly http client with support for 
`GET`, `POST`, `PUT`, `DELETE` methods. It supports following authentication types:

- Basic
- Bearer Token
- Custom Headers

## Usage

Http request with basic authentication

```go
req := Request{
        Url:     "http://localhost:8080/auth/basic/user",
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
```

Http request with bearer token authentication

```go
req := Request{
        Url:     "http://localhost:8080/auth/bearer_token/users",
        Method:  http.MethodGet,
        Body:    nil,
        Timeout: 10,
        Auth:    &Auth{
            BearerToken: &AuthBearerToken{
                Token: "token",
            },
        },
    }
```

Http request with custom authentication 

```go
req := Request{
        Url:     "http://localhost:8080/auth/custom/users",
        Method:  http.MethodGet,
        Body:    nil,
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