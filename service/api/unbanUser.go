package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
)

/*
taken uid of the user who wants unbband id of user banned and delete last one
*/
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var from, to int
	var err error
	var tk *security.Token
	var isBan bool

	/*
		Parse URL parameters
	*/

	//   get :uid parameter in path
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	//   get :bannedId parameter in path
	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/

	if isBan, err = rt.db.IsBanned(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	if !isBan {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

	}

	/*
		Secure Bearer Authentication
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		Check if user can would put ban is account owner
	*/
	if tk.Value != from {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return

	}

	/*
		Try to remove ban
	*/

	if _, err = rt.db.DelBan(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)

	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
