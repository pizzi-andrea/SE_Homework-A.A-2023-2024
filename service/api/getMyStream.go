package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
given *uid* of user that who wants to get all photo in his stream
*/
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid int
	var err error
	var stream []database.Post
	var user *database.User
	var tk *security.Token
	var limit int
	offset := 0

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}
	pLimit := r.URL.Query().Get("limit")
	pUsername := r.URL.Query().Get("username")

	if limit, err = strconv.Atoi(pLimit); pLimit != "" && err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return

	}

	rr, err := regexp.MatchString("^.*?$", pUsername)
	if !(rr && err == nil && len(pUsername) <= 16) {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*

	 */
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		The stream is personal only owner user can see it
	*/
	if tk.Value != user.Uid {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if stream, err = rt.db.GetMyStream(user.Uid, pUsername, true, []database.OrderBy{}); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if limit == 0 || limit > len(stream) {
		limit = len(stream)
	}

	if offset > len(stream) { //  500 response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	stream = stream[offset:limit]

	for i := range stream {
		stream[i].Location = strings.TrimSuffix("/images/:photoId", ":photoId") + strconv.Itoa(stream[i].Refer)
		ctx.Logger.Infof("Path:%s\n", stream[i].Location)
	}

	if len(stream) == 0 { //   204 response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return

	} else {
		if err = json.NewEncoder(w).Encode(stream); err != nil {
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)
		}
	}

}
