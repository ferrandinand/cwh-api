package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ferrandinand/cwh-lib/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) (string, bool)
}

type RemoteAuthRepository struct {
	URL string
}
type AuthReponse struct {
	UserID       map[string]string
	IsAuthorized map[string]bool
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) (string, bool) {

	u := buildVerifyURL(r.URL, token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return "", false
	} else {
		m := AuthReponse{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return "", false
		}
		return m.UserID["user"], m.IsAuthorized["isAuthorized"]

	}
}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&user_id=2000
*/
func buildVerifyURL(authURL string, token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: authURL, Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository(url string) RemoteAuthRepository {
	return RemoteAuthRepository{
		URL: url,
	}
}
