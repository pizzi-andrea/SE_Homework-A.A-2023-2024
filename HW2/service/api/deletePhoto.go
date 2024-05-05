package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

// deletePhoto permit to photo owner to delete updated photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid int
	var photoId int
	var err error
	var user *database.User
	var photo *database.Photo
	var tk *security.Token
	var s bool

	//   Parsing URL parameters in path
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	//   check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if user == nil || photo == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Applay barrear authentication.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		checks if the user who wants delete photo is the owner
	*/
	if tk.Value != uid {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	//   delete photo
	if s, err = rt.db.DelPhoto(photo.PhotoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return

	}
	if s {
		ctx.Logger.Debug("Photo deleted with success")
	} else {
		ctx.Logger.Debug("Photo not deleted")
	}
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
