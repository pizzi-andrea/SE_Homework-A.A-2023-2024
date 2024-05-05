package api

import (
	"bytes"
	"image"
	"image/png"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/database"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photoId int
	var err error
	var img *database.Photo
	var buffer image.Image

	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if img, err = rt.db.GetPhoto(photoId); err != nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "image/png")
	b := bytes.NewBuffer(img.ImageData)

	if buffer, err = png.Decode(b); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	if err = png.Encode(w, buffer); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return
	}

}
