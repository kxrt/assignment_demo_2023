package main

import (
	"context"
	"fmt"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()
	err := PushMessage(req.GetMessage())
	if err != nil {
		return nil, err
	}
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()
	messages, err := PullMessages(req)
	if err != nil {
		return nil, err
	}
	resp.SetMessages(messages)
	fmt.Println(messages)
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}
