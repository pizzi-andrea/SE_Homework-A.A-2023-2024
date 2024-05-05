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
give *uid* and *photoId* and get photo associated
*/
func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var uid int
	var err error
	var user *database.User
	var tk *security.Token
	var isBan bool
	var photoId int
	var photo *database.Photo
	var likes []database.User

	//  Parsing URL parameters
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !(database.ValidateId(photoId) && database.ValidateId(uid)) {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return

	}

	//  check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if user == nil || photo == nil {
		w.Header().Set("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Applay barrear authentication. A user can see photos of another user as long as they have not been banned
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if isBan {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if likes, err = rt.db.GetLikes(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	w.Header().Set("content-type", "application/json")
	if err = json.NewEncoder(w).Encode(likes); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

}
