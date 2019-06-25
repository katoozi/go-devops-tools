package nap

import (
	"testing"
)

func TestAuthToken(t *testing.T) {
	token := NewAuthToken("somerandomtoken")
	header := token.AuthorizationHeader()
	if header != "toekn somerandomtoken" {
		t.Fail()
	}
}

func TestAuthBasic(t *testing.T) {
	basic := NewAuthBasic("user", "password")
	header := basic.AuthorizationHeader()
	if header != "basic dXNlcjpwYXNzd29yZA==" {
		t.Fail()
	}
}
