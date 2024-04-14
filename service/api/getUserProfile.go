package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
give user id and put user profile. User profile is object contain
all information on user, in particular:
  - information about user
  - stream photos updated
  - number of photo have been updated
  - number of followers
  - number of following
*/
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var uid int
	var err error
	var user *database.User
	var isBan bool
	var follower, following []database.User
	var photos database.Stream
	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("GetUserFromId::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	// if user not exist path is invalid
	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		Applay barrear authentication. Only owner user and users not banned by owner can access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		check if user wont access is not beetween banned users from owner
	*/

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("IsBanned::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	// if user that would see profile is banned by owner profile user can not access to it.
	if isBan {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	/*
		get followers of user
	*/
	if follower, err = rt.db.GetFollowers(uid, "", true); err != nil {
		ctx.Logger.Errorf("GetFollowers::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	/*
		get users following by user
	*/
	if following, err = rt.db.GetFollowed(uid, "", true); err != nil {
		ctx.Logger.Errorf("GetFollowed::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	/*
		get photo stream. Photo stream is composed by photos of other followed users
	*/
	if photos, err = rt.db.GetPosts(uid, []database.OrderBy{}); err != nil {
		ctx.Logger.Errorf("GetPhotos::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	for i := range photos {
		photos[i].Location = strings.TrimSuffix("/images/:photoId", ":photoId") + strconv.Itoa(photos[i].Refer)

	}

	/*
		put in response body user profile rappresentation
	*/
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK) //   200
	// encode in json format the response
	if err = json.NewEncoder(w).Encode(database.Profile{
		User:      *user,
		Stream:    photos,
		Follower:  len(follower),
		Following: len(following),
	}); err != nil {
		ctx.Logger.Errorf("Encode::%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		return

	}

}
