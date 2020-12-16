package skeleton

type Request struct {
	Url     string
	Method  string
	Body    map[string]string
	Timeout int
	Auth    *Auth
}

type Auth struct {
	Basic       *AuthBasic
	BearerToken *AuthBearerToken
	Custom      map[string]string
}

type AuthBasic struct {
	Username string
	Password string
}

type AuthBearerToken struct {
	Token string
}
