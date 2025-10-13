package main

type AppendEntriesRequest struct{}

type AppendEntriesResponse struct{}

func (s *RaftServer) AppendEntries(req *AppendEntriesRequest) *AppendEntriesResponse {
	return nil
}
