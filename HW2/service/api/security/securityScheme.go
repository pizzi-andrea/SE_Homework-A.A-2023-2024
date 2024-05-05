package security

/*
This object rappresent a token identification. The token will be used by users to authenticate to the system.

	The token is composed from two fields:
	  - tokenId: corresponding to uid of owner
	  - owner:   username of owner
*/
type Token struct {
	Value int
}
