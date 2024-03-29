package nap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

type AuthToken struct {
	Token string
}

type AuthBasic struct {
	Username string
	Password string
}

func NewAuthBasic(username, password string) *AuthBasic {
	return &AuthBasic{
		Username: username,
		Password: password,
	}
}

func (a *AuthBasic) AuthorizationHeader() string {
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	encContent := fmt.Sprintf("%s:%s", a.Username, a.Password)
	enc.Write([]byte(encContent))
	enc.Close()
	content, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Read failed:", err)
	}
	return fmt.Sprintf("basic %s", string(content))
}

func (auth *AuthToken) AuthorizationHeader() string {
	return fmt.Sprintf("token %s", auth.Token)
}

type Authentication interface {
	AuthorizationHeader() string // "Basic <base64-encoded string>"
}

func NewAuthToken(token string) *AuthToken {
	return &AuthToken{
		Token: token,
	}
}
