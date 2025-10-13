package main

import (
	"github.com/atoaidoocr7/distributed-systems/raft/api"
	"github.com/atoaidoocr7/distributed-systems/raft/server/util"
)

type RaftServer struct {
	serverType util.ServerType
	log        api.Log
}
