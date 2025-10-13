package main

type Client struct {
}

func (c *Client) NewClient() *Client {
	return &Client{}
}

func (c *Client) RequestVote(request *RequestVoteRequest) *RequestVoteResponse {
	return &RequestVoteResponse{}
}
