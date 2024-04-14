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

func (rt *_router) isBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var from, to int
	var err error
	var result *database.User
	var tk *security.Token

	/*
		Parse parameters in path
	*/
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		Secure Bearer Authentication, check if user wont put ban is account owner.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		Check if user  that want ban user can it. Only user have put ban
		can delete it
	*/
	if tk.Value != from {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return

	}

	/*
		Try to ban user, if queries fails server response with error message
	*/

	if result, err = rt.db.GetUserBanned(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)

	}

	if result != nil {
		if err = json.NewEncoder(w).Encode(*result); err != nil {

			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

		}
	} else {
		w.WriteHeader(http.StatusNoContent)

	}

}
