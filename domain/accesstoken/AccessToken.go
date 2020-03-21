package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/crypto"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

const (
	expirationTime = 24
)

//AccessToken struct
type AccessToken struct {
	Token    string `json:"access_token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

//ValidateToken sent
func (at *AccessToken) ValidateToken() errors.RestErr {
	at.Trim()
	if at.Token == "" {
		return errors.BadRequestError("invalid Access Token ID", errors.New("empty token"))
	}
	return nil
}

//Validate entire data
func (at *AccessToken) Validate() errors.RestErr {
	if err := at.ValidateToken(); err != nil {
		return err
	}
	if at.UserID <= 0 {
		return errors.BadRequestError("invalid User ID", errors.New("user id zero"))
	}
	if at.ClientID <= 0 {
		return errors.BadRequestError("invalid Client ID", errors.New("client id zero"))
	}
	return nil
}

//Trim empty spaces
func (at *AccessToken) Trim() {
	at.Token = strings.TrimSpace(at.Token)
}

//IsExpired if token is done for
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

//Generate Unique Token
func (at *AccessToken) Generate() {
	at.Token = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}

//GetNewAccessToken change expiration
func GetNewAccessToken(id int64) AccessToken {
	return AccessToken{
		UserID:  id,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
