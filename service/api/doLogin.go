package api

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

type loginUser struct { //   struct for json Marshaler
	Name database.Username `json:"name"`
}

/*
If the user does not exist, it will be created,

	and an identifier is returned.
	If the user exists, the user identifier is returned.
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	var u loginUser
	var token security.Token
	var user *database.User

	/*
		Decode values in body request *r
	*/
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return

	}

	//    check if user is registred
	if user, err = rt.db.GetUserFromUser(u.Name); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500 code
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	//   check format data

	/*
		If user registred, make logged
	*/
	if user != nil {

		token = security.Token{
			Value: user.Uid,
		}

		security.RecordToken(token)
		if err = json.NewEncoder(w).Encode(token); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)
			return

		}
		w.Header().Set("content-type", "application/json") //  200 code

		return

	}
	//   check db query error

	if !(database.ValidateUsername(u.Name)) {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if user, err = rt.db.PostUser(u.Name); err != nil || user == nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500 code
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	/*
		if user not registred,  registred it and make logged
	*/

	token = security.Token{
		Value: user.Uid,
	}

	if err = json.NewEncoder(w).Encode(token); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	security.RecordToken(token) //   201 code
	w.Header().Set("content-type", "application/json")

}
