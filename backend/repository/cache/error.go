package cache

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("time expired")
	ErrorVoted          = errors.New("you have voted before")
)
