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
gived uid and *followedId* then remove follower *followerId* from user followers
*/
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var from, to int
	var err error
	var tk *security.Token
	var isBan bool
	var action bool
	var user *database.User
	/*
		Parse URL parameters in path
	*/
	if from, err = strconv.Atoi(ps.ByName("followerId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}
	if to, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		Check if value in parameters are valid values in accord to type definittiion
	*/
	if !(database.ValidateId(from) && database.ValidateId(to)) {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/

	if user, err = rt.db.GetUserFromId(to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	if user, err = rt.db.GetUserFromId(from); err != nil {
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

	if user == nil {
		w.Header().Set("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if from == to {
		ctx.Logger.Errorln("the user cannot follow himself")
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(http.StatusForbidden)
		return

	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if tk.Value != from {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return

	}

	/*
		get banned user and check if not banned
	*/

	if isBan, err = rt.db.IsBanned(to, from); isBan {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return

	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if action, err = rt.db.PutFollow(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   204
		w.WriteHeader(http.StatusNoContent)
		return

	}

	if action {
		if user, err = rt.db.GetUserFromId(from); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)
			return

		}

		if err = json.NewEncoder(w).Encode(*user); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)
			return
		}

		w.Header().Set("content-type", "application/json") //  201
		w.WriteHeader(http.StatusCreated)

	} else {
		ctx.Logger.Infoln("user just follow")
		w.Header().Set("content-type", "text/plain") //  204
		w.WriteHeader(http.StatusNoContent)

	}
}
