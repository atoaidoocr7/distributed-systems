package util

type ServerType int

const (
	Unknown ServerType = iota
	Leader
	Candidate
	Follower
)
