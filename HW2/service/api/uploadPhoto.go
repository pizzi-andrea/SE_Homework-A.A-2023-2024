package api

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
UploadPhoto method allow users to update their photos in WASAPhoto system. Request MUST formatted like
multipart/form-data. It will parserd values and stored its in db. The photo will be encoding in base64 format
and stored in db.
*/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid int
	var err error
	var user *database.User
	var tk *security.Token
	var file multipart.File
	var newPost *database.Post = &database.Post{}

	//   Parsing URL parameters
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//   check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

	//  if user not exist path is not valid
	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)
		return

	}

	//	Applay barrear authentication. only owner can post photo in his account
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		ctx.Logger.Error("token missing\n")
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)
		return
	}

	//	checks if the user who wants post photo is account owner
	if tk.Value != uid {
		ctx.Logger.Error("User not author.\n")
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		return
	}
	//  parsing body values (image upload operation)
	var photo *database.Photo = &database.Photo{}

	// max byte for photo is 5MB
	if err = r.ParseMultipartForm(database.MaxByteFormData); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// read photo from form
	if file, _, err = r.FormFile("img"); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	// parse photo
	if photo.ImageData, err = io.ReadAll(file); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)
	}

	newPost.DescriptionImg = r.PostFormValue("desc")

	// check if post is valid
	if photo.ImageData == nil || len(photo.ImageData) == 0 {
		ctx.Logger.Errorf("photo <%v> not valid", photo)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	//  check if format photo in conform to APIs specifications

	if newPost, err = rt.db.CreatePost(user.Uid, photo.ImageData, newPost.DescriptionImg); err != nil {
		ctx.Logger.Errorf("CreatePost::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	// check if post is created
	if newPost == nil {
		ctx.Logger.Errorf("photo <%v> not updated", photo)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	// write post created in response message
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	photo.ImageData = nil //   discard img data
	if err = json.NewEncoder(w).Encode(*newPost); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}
}
