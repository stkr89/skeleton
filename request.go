package skeleton

// Request specifies general fields required to make an http call
type Request struct {
	Url     string
	Method  string
	Body    map[string]string
	Timeout int
	Auth    *Auth
}

// Auth contains pointers to the type of authentications supported by skeleton
type Auth struct {
	Basic       *AuthBasic
	BearerToken *AuthBearerToken
	Custom      map[string]string
}

// AuthBasic provides fields required for http basic authentication
type AuthBasic struct {
	Username string
	Password string
}

// AuthBearerToken provides field required for http token based authentication
type AuthBearerToken struct {
	Token string
}
