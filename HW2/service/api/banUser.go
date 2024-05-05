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

// taken uid of the user who wants to ban and uid of the user to be banned.
// if the user is banned, will be nofthig to do, otherwise the user will be banned.
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var result bool
	var from, to int
	var err error
	var t *database.User
	var tk *security.Token

	//	Parse parameters path url

	// parse :uid
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	// parse :bannedId
	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	//	get user from id that want ban and user that want ban
	if _, err = rt.db.GetUserFromId(from); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if t, err = rt.db.GetUserFromId(to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// check if user want ban himself
	if from == to {
		ctx.Logger.Errorln("the user cannot ban himself")
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(http.StatusForbidden)
		return

	}

	//	Secure Bearer Authentication, check if user wont put ban is account owner.
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		ctx.Logger.Errorln("token not valid or missing")
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	//	Check if user  that want ban user can it. Only user have put ban
	//	and can remove it
	if tk.Value != from {
		ctx.Logger.Errorln("not authorized")
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		return

	}

	//	Try to ban user, if queries fails server response with error message
	if result, err = rt.db.PutBan(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	//   check now if user follow other user
	if result {

		// Delete follows  beetween  users
		if _, err = rt.db.DelFollow(to, from); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return
		}
		if _, err = rt.db.DelFollow(from, to); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return
		}

		// write user banned into body response
		if err = json.NewEncoder(w).Encode(*t); err != nil {

			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)
			return
		}

		return
	}

	w.Header().Set("content-type", "text/plain") //  204
	w.WriteHeader(http.StatusNoContent)

}
