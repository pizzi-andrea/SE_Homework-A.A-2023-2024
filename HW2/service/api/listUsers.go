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

// this method returns the list of currently registered users.
// All users are identified by the unique ID (uid) and unique username chosen by the users
func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var limit int
	var err error
	var users []database.User
	var tk *security.Token

	//   get values from URL

	pLimit := r.URL.Query().Get("limit")
	pUsername := r.URL.Query().Get("username")

	/*
		Applay barrear authentication.
		if the user has passed the login phase then he will be able to access.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*parse URL parameters */
	if limit, err = strconv.Atoi(pLimit); pLimit != "" && err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return

	}
	// get list of users
	if users, err = rt.db.GetUsers(pUsername, true); err != nil {
		ctx.Logger.Errorf("GetUsers::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	//  future dev-----------------------------------------
	offset := 0
	//  -----------------------------------------------------------

	//   response object

	if limit == 0 || limit > len(users) {
		limit = len(users)
	}
	var usersOk []database.User
	users = users[offset:limit]
	//   if response body is empty
	if len(users) == 0 {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return
	}
	var b bool

	// put users to return in the list
	for _, u := range users {

		if b, err = rt.db.IsBanned(u.Uid, tk.Value); err != nil {
			ctx.Logger.Errorf("IsBanned::%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return
		}

		if !b {
			usersOk = append(usersOk, u)
		}
	}

	//   if response body is empty
	if len(usersOk) == 0 {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return
	}

	//   write values getted in to response

	if err = json.NewEncoder(w).Encode(usersOk); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}
	w.Header().Set("content-type", "application/json") //   200
}
