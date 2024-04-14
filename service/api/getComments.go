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

// give *uid* and *photoId* and get photo associated
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var uid, photoId int
	var err error
	var user *database.User
	var tk *security.Token
	var isBan bool
	var photo *database.Photo
	var comments []database.Comment
	var username string
	limit := 0
	offset := 0

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

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil && limit >= 1 {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username = r.URL.Query().Get("username")

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

	if comments, err = rt.db.GetComments(photoId, username, true); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	if limit == 0 || limit > len(comments) {
		limit = len(comments)
	}

	if offset > len(comments) { // 500 response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		return
	}

	comments = comments[offset:limit]
	// manca l'ordinamento secondo la data

	w.Header().Set("content-type", "application/json")
	if err = json.NewEncoder(w).Encode(comments); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

}
