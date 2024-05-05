package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var photoId, uid, likeUserId int
	var err error
	var like *database.User
	var tk *security.Token
	var user *database.User
	var photo *database.Photo
	var rr bool

	// parse :uid
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	// parse :photoId
	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	// parse :likeUserId
	if likeUserId, err = strconv.Atoi(ps.ByName("likeUserId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	// check if user exists
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	// check if post exists
	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	// check if user that put like exists
	if like, err = rt.db.GetUserFromId(likeUserId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	// check if path is valid
	if photo == nil || user == nil || like == nil {
		w.Header().Set("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	//	Check if user that wont put follows can do it

	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	// check if user is banned
	if rr, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	// 	if user is banned do not allow to follow/unfollow action
	if rr {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	// only user that putted like can remove it
	if like.Uid != tk.Value {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	// remove like
	if _, err = rt.db.DelLike(like.Uid, photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
