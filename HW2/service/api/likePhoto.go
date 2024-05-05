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

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var photoId, uid, likeUserId int
	var err error
	var tk *security.Token
	var user *database.User
	var post *database.Post
	var rr bool

	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if likeUserId, err = strconv.Atoi(ps.ByName("likeUserId")); err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if !(database.ValidateId(photoId) && database.ValidateId(uid)) {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return

	}

	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("GetUserFromId::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if post, err = rt.db.GetPost(photoId); err != nil {
		ctx.Logger.Errorf("GetPost::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if post == nil || user == nil {
		w.Header().Set("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if tk.Value != likeUserId {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		return

	}

	if likeUserId == uid {
		ctx.Logger.Errorln("User can not put like on its photo")
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(http.StatusForbidden)
		return

	}

	if rr, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("IsBanned::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if rr {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if rr, err = rt.db.PutLike(likeUserId, photoId); err != nil {
		ctx.Logger.Errorf("PutLike::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if !rr {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(*user); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

}
