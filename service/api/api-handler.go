package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Register routes

	rt.router.GET("/", rt.wrap(rt.getContextReply))

	//  Logs in the user
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // required

	//  list registred users
	rt.router.GET("/users/", rt.wrap(rt.listUsers))

	//   assign or update username
	rt.router.PUT("/users/:uid/", rt.wrap(rt.setMyUserName)) // required

	//   get specific user profile
	rt.router.GET("/users/:uid/", rt.wrap(rt.getUserProfile)) // required

	//   get all followers
	rt.router.GET("/users/:uid/followers/", rt.wrap(rt.listFollowers))

	//   unfollow user
	rt.router.DELETE("/users/:uid/followers/:followerId", rt.wrap(rt.unfollowUser)) // required

	//  follow user
	rt.router.PUT("/users/:uid/followers/:followerId", rt.wrap(rt.followUser)) // required

	//   list personal stream photos
	rt.router.GET("/users/:uid/myStream/", rt.wrap(rt.getMyStream)) // required

	//   ban user identificated by *uid*
	rt.router.PUT("/users/:uid/banned/:bannedId", rt.wrap(rt.banUser)) // required

	//   unban user identificated by uid
	rt.router.DELETE("/users/:uid/banned/:bannedId", rt.wrap(rt.unbanUser)) // required

	// check if user had banned other user
	rt.router.GET("/users/:uid/banned/:bannedId", rt.wrap(rt.isBanned))

	//   update photo
	rt.router.POST("/users/:uid/myPhotos/", rt.wrap(rt.uploadPhoto)) // required

	//   list stream photos updated
	rt.router.GET("/users/:uid/myPhotos/", rt.wrap(rt.listPost))

	//   delete photo updated
	rt.router.DELETE("/users/:uid/myPhotos/:photoId/", rt.wrap(rt.deletePhoto)) // required

	//   get likes collected by photo
	rt.router.GET("/users/:uid/myPhotos/:photoId/likes/", rt.wrap(rt.getLikes))

	//   put like a photo
	rt.router.PUT("/users/:uid/myPhotos/:photoId/likes/:likeUserId", rt.wrap(rt.likePhoto)) // required

	//   check if user likeUserId has putted like on photo
	rt.router.GET("/users/:uid/myPhotos/:photoId/likes/:likeUserId", rt.wrap(rt.checkLike)) // required

	//   remove like a photo
	rt.router.DELETE("/users/:uid/myPhotos/:photoId/likes/:likeUserId", rt.wrap(rt.unlikePhoto)) // required

	//  add comment a photo
	rt.router.POST("/users/:uid/myPhotos/:photoId/comments/", rt.wrap(rt.commentPhoto)) // required

	//  get comments on photo
	rt.router.GET("/users/:uid/myPhotos/:photoId/comments/", rt.wrap(rt.getComments))

	//  delete comment on photo
	rt.router.DELETE("/users/:uid/myPhotos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto)) // required

	// specific endpoint used for get photo data
	rt.router.GET("/images/:photoId", rt.wrap(rt.getImage))
	//  Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router

}
