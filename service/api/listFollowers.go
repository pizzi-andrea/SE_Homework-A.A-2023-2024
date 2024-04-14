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
give a UID return a list contanings all followers user
*/
func (rt *_router) listFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid int
	var err error
	var isBan bool
	var user *database.User
	var tk *security.Token
	var limit int
	var username database.Username
	offset := 0

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	//   validate username format
	username = r.URL.Query().Get("username")

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil && r.URL.Query().Get("limit") != "" {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

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

	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		check if token is given and it user is currently logged
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		get banned user and check if not banned
	*/

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if isBan {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if followers, err := rt.db.GetFollowers(uid, username, true); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.WriteHeader(http.StatusInternalServerError)

		return

	} else {

		if limit == 0 || limit > len(followers) {
			limit = len(followers)
		}

		if offset > len(followers) { //  500 response
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)

			return
		}

		followers = followers[offset:limit]

		if len(followers) == 0 { //   204 response
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusNoContent)

		} else { //  200 repsonse
			w.Header().Set("content-type", "application/json")
			if err = json.NewEncoder(w).Encode(followers); err != nil {
				ctx.Logger.Errorf("%w", err)
				w.Header().Set("content-type", "text/plain") //   500
				w.WriteHeader(ServerError.StatusCode)

			}
		}

	}

}
