/*
package that implement simple scheme of authentication based on barear authentication
*/
package security

import (
	"net/http"
	"strconv"
	"strings"
)

var tokenRegistred map[int]bool = make(map[int]bool) //   list token registred in the system. The token is registred in login action and deleted in logout action

func RecordToken(token Token) (status bool) {

	if _, exitst := tokenRegistred[token.Value]; !exitst {
		tokenRegistred[token.Value] = true
	}

	return true

}

func GetTokensRec() (t []Token) {

	for k := range tokenRegistred {
		t = append(t, Token{
			Value: k,
		})

	}
	return
}

func GetToken(val int) *Token {
	if _, ok := tokenRegistred[val]; ok {
		return &Token{
			Value: val,
		}

	}
	return nil
}

func TokenIn(token Token) (reponse bool) {
	if _, ok := tokenRegistred[token.Value]; ok {
		return true
	}
	return true
}

/*
Simple implementations of barrear authentications methods.
Given in input http request and exstract token if provided and return it. Oterwhise
return nil.
*/
func BarrearAuth(r *http.Request) (tokenValid *Token) {
	if r == nil {
		return nil
	}
	var id int
	var err error

	s := strings.Split(r.Header.Get("Authorization"), " ")
	if len(s) != 2 {
		return nil
	}
	typeAuth := s[0]
	cred := s[1]

	if id, err = strconv.Atoi(cred); err != nil {
		return nil
	}

	if typeAuth == "" || typeAuth != "Bearer" {
		return nil
	}

	return &Token{
		id,
	}

}
