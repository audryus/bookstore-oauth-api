package accesstoken

import "gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"

const (
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

//LoginRequest struct from the Users API
type LoginRequest struct {
	GrantType  string `json:"grant_type"`
	Scope      string `json:"scope"`
	AuthID     string `json:"auth_id"`
	AuthSecret string `json:"auth_secret"`
}

//Validate the login request
func (lr *LoginRequest) Validate() errors.RestErr {
	switch lr.GrantType {
	case grantTypePassword:
	case grantTypeClientCredentials:
		break
	default:
		return errors.BadRequestError("invalid grant_type parameter", errors.New("grant type not implemented"))
	}
	return nil
}
