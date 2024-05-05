
class Photo {

}

/**
 * User
 */
export class User {
	constructor(id, username){
		this.id = id 
		this.username = username 
	}
}

/**
 * User Profile
 */
export default class Profile {
	
	constructor(id, username, follower, following){
		this.user = new User(id, username)
		this.follower = follower
		this.following = following
	}

	
}