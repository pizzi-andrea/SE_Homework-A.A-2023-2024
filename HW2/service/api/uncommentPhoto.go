package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

// /users/{uid}/myPhotos/{photoId}/comments/
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var photoId, uid, comm int
	var err error
	var msg *database.Comment
	var tk *security.Token
	var user *database.User
	var photo *database.Photo
	var rr bool

	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if comm, err = strconv.Atoi(ps.ByName("commentId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("GetPhoto::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if msg, err = rt.db.GetComment(comm); err != nil {
		ctx.Logger.Errorf("GetComment::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if photo == nil || user == nil || msg == nil {
		ctx.Logger.Errorf("path::%w", err)
		w.Header().Set("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		ctx.Logger.Errorf("NoToken::%w", err)
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if rr, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if rr {
		ctx.Logger.Errorf("Banned\n")
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if msg.Author.Uid != tk.Value {
		ctx.Logger.Errorf("Not Author:: (%d != %d)\n", msg.Author.Uid, tk.Value)
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if rr, err = rt.db.DelComment(msg.CommentId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if rr {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		ctx.Logger.Infof("User %d uncommented photo %d", user.Uid, photo.PhotoId)
		return

	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	ctx.Logger.Errorln("Comment not found")

}
