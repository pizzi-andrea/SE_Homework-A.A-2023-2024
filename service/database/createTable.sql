-- enable foreign constraint
PRAGMA foreign_keys = ON;
-- Create Users schema

CREATE TABLE IF NOT EXISTS Users (
	uid INTEGER  PRIMARY KEY, 	-- user id
	username TEXT UNIQUE NOT NULL	-- username of user
);

-- Create Photo schema

CREATE TABLE IF NOT EXISTS Photos (
	photoId INTEGER  PRIMARY KEY,	-- photo identificator
	owner INTEGER NOT NULL,		-- uid of owner
	descriptionImg TEXT NOT NULL,	-- image description
	imageData BLOB NOT NULL,	-- file
	timeUpdate TIMESTAMP NOT NULL,	-- time to update
	
	
	CONSTRAINT ref_to_user FOREIGN KEY(owner) REFERENCES Users(uid) ON DELETE CASCADE
);


-- Create Comments schema

CREATE TABLE IF NOT EXISTS Comments (
	commentId INTEGER   PRIMARY KEY,
	author INTEGER NOT NULL,
	photo INTEGER NOT NULL,
	text_ TEXT NOT NULL,
	timeStamp_ TIMESTAMP NOT NULL,

	CONSTRAINT ref_to_user FOREIGN KEY(author) REFERENCES Users(uid) ON DELETE CASCADE,
	CONSTRAINT ref_to_photo FOREIGN KEY(photo) REFERENCES Photos(photoId) ON DELETE CASCADE
);


-- Create likes schema

CREATE TABLE IF NOT EXISTS Likes (
	user INTEGER NOT NULL,
	photo INTEGER NOT NULL,
	PRIMARY KEY(user, photo),
	CONSTRAINT ref_to_user FOREIGN KEY(user) REFERENCES Users(uid) ON DELETE CASCADE,
	CONSTRAINT ref_to_photo FOREIGN KEY(photo) REFERENCES Photos(photoId) ON DELETE CASCADE
);


-- Create followers schema

CREATE TABLE IF NOT EXISTS Followers (
	from_ INTEGER NOT NULL,
	to_ INTEGER NOT NULL,
	
	PRIMARY KEY(from_, to_)
	CONSTRAINT ref_to_from FOREIGN KEY(from_) REFERENCES Users(uid) ON DELETE CASCADE,
	CONSTRAINT ref_to FOREIGN KEY(to_) REFERENCES Users(uid) ON DELETE CASCADE
);
-- Create banned schema

CREATE TABLE IF NOT EXISTS Bans (
	from_ INTEGER NOT NULL,
	to_ INTEGER NOT NULL,
	
	PRIMARY KEY(from_, to_),
	CONSTRAINT ref_to_from FOREIGN KEY(from_) REFERENCES Users(uid) ON DELETE CASCADE,
	CONSTRAINT ref_to FOREIGN KEY(to_) REFERENCES Users(uid) ON DELETE CASCADE
);
