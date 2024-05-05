package database

import (
	"regexp"
	"time"
)

type Id = int              //   Identificator at 64-bit
type TimeStamp = time.Time //   this components describe timestamp value conform to [RFC3339](https://  datatracker.ietf.org/doc/html/rfc3339#section-5.6) specification
type Username = string     //   username of a user
type Validator interface {
	Verify() bool
}

/*this object rappresent a photo*/
type Photo struct {
	PhotoId   Id     `json:"photoId"`
	ImageData []byte `json:"imageData"` //   data
}

/*this object rappresent a user*/
type User struct {
	Uid      Id       `json:"uid"`
	Username Username `json:"username"`
}

// user profile rappresentation
type Profile struct {
	User      User   `json:"user"`
	Stream    Stream `json:"stream"`
	Follower  int    `json:"follower"`  // number user that follow a specific user
	Following int    `json:"following"` //  numer of users following by specific user

}

// this object rappresent a Image update from user
type Image struct {
	PhotoId Id     `json:"photoId"`
	Data    []byte `json:"data"` // binary string rappresenting the data of image
}

// this object rappresent a post.
// A post is provided to photo and list of like and comments that recived.
type Post struct {
	Owner          Id        `json:"owner"`
	Refer          Id        `json:"refer"`
	Location       string    `json:"location"`
	Likes          []User    `json:"likes"`
	Comments       []Comment `json:"comments"`
	TimeUpdate     TimeStamp `json:"timeUpdate"`
	DescriptionImg string    `json:"descriptionImg"`
}

/*this object rappresent a photo*/
type Stream = []Post //  model of stream of photos

/*this object rappresent a comment on a photo.*/
type Comment struct {
	CommentId Id        `json:"commentId"`
	Author    User      `json:"author"`    //  this object rappresent a user
	Text      string    `json:"text"`      //  comment text encoded in UNICODE format
	TimeStamp TimeStamp `json:"timeStamp"` //  this components describe timestamp value conform to RFC3339 specification

}

func ValidateId(v Id) bool {
	return true
}

func ValidateTimeStamp(time string) bool {
	s, err := regexp.MatchString(" '[1-9]\\d{3}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$", time)
	return 20 <= len(time) && len(time) >= 21 && s && err == nil
}

func ValidateUsername(u string) bool {

	s, err := regexp.MatchString("^.*?$", u)
	return len(u) >= 3 && len(u) <= MaxByteUsername && s && err == nil
}

func ValidateStream(s Stream) bool {

	r := len(s) <= MaxSizeStream
	if r {
		for _, p := range s {

			if !p.Verify() {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (p *Photo) Verify() bool {
	return true
}

func (u *User) Verify() bool {
	return ValidateUsername(u.Username) && ValidateId(u.Uid)
}

func (p *Profile) Verify() bool {
	return p.User.Verify() && p.Follower >= 0 && p.Following >= 0 && ValidateStream(p.Stream)
}

func (c *Comment) Verify() bool {

	r, err := regexp.MatchString("^.*?$", c.Text)
	return c.Author.Verify() && len(c.Text) >= 1 && len(c.Text) <= MaxByteDescription && r && err == nil
}

func (p *Post) Verify() bool {
	return true
}
