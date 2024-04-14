package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
assign username or update old username with new username.
The username to set is in the body request
*/

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid int
	var username string
	var err error
	var user *database.User

	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	// if user not exist ...
	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Decode values in body request *r
	*/
	if err = json.NewDecoder(r.Body).Decode(&username); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return

	}

	/*
		Validate username to update
	*/
	if !(database.ValidateUsername(username)) {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	/*
		Applay barrear authentication. Username can update only his username
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)
		return
	}

	/*
		checks if the user who wants to change username is the owner
	*/
	if tk.Value != uid {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		return
	}

	/*
		Update username
	*/
	if _, err = rt.db.SetUsername(uid, username); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
