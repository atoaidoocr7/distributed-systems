package main

import (
	"github.com/atoaidoocr7/distributed-systems/raft/api"
	"github.com/atoaidoocr7/distributed-systems/raft/server/util"
)

// RaftServer represents a very basic raft server
type RaftServer struct {
	id         string
	serverType util.ServerType
	log        *api.Log
	serverTerm int
	nextIndex  map[string]int
	matchIndex map[string]int
}

func NewRaftServer(id string) *RaftServer {
	return &RaftServer{
		serverType: util.Follower,
		log:        api.NewLog(),
		serverTerm: 0,
		id:         id,
	}
}

func (server *RaftServer) run() {

	/**
	TODO:
	- Every n seconds, where n == configured timeout:

	- If server is follower and leader sends heartbeat before timer hits, reset timer
	- If timer is up, upgrade

	 **/
}

func (server *RaftServer) upgradeToLeader() {}

func (server *RaftServer) upgradeToCandidate() {}

func (server *RaftServer) downgradeToFollower() {}
