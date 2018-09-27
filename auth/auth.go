package auth

import (
	"crypto/rand"
	"io/ioutil"
	"time"

	"gopenguin/config"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       string
	Username string
	Role     string
}

type Auth struct {
	Secret []byte
}

func NewAuth(config *config.Config) (*Auth, error) {
	authInfo := config.Auth

	auth := new(Auth)

	if authInfo.SecretSource == "" {
		auth.Secret = make([]byte, 256)

		rand.Read(auth.Secret)
	} else {
		buf, err := ioutil.ReadFile(authInfo.SecretSource)

		if err != nil {
			return nil, err
		}

		auth.Secret = buf
	}

	return auth, nil
}

func (a *Auth) keyFunc(token *jwt.Token) (interface{}, error) {
	return a.Secret, nil
}

func (a *Auth) DecodeToUser(tokenString string) (*User, error) {
	token, err := a.Decode(tokenString)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	user := new(User)

	user.ID = claims["id"].(string)
	user.Username = claims["username"].(string)
	user.Role = claims["role"].(string)

	return user, nil
}

func (a *Auth) Decode(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, a.keyFunc)
}

func (a *Auth) Encode(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().AddDate(0, 1, 0)

	return token.SignedString(a.Secret)
}
