package main

type RequestVoteRequest struct{}
type RequestVoteResponse struct{}

func (r *RaftServer) RequestVote(request *RequestVoteRequest) *RequestVoteResponse {
	return nil
}
