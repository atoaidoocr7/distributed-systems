package main

import (
	"github.com/atoaidoocr7/distributed-systems/raft/api"
	"github.com/atoaidoocr7/distributed-systems/raft/server/util"
)

// RaftServer represents a very basic raft server
type RaftServer struct {
	serverType util.ServerType
	log        *api.Log
	serverTerm int
	serverId   string
	nextIndex  map[string]int
	matchIndex map[string]int
}

func NewRaftServer(serverName string) *RaftServer {
	return &RaftServer{
		serverType: util.Follower,
		log:        api.NewLog(),
		serverTerm: 0,
		serverId:   serverName,
	}
}
