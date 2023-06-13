package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc"
	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc/imservice"
	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/proto_gen/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var cli imservice.Client

func main() {
	r, err := etcd.NewEtcdResolver([]string{"etcd:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli = imservice.MustNewClient("demo.rpc.server",
		client.WithResolver(r),
		client.WithRPCTimeout(1*time.Second),
		client.WithHostPorts("rpc-server:8888"),
	)

	h := server.Default(server.WithHostPorts("0.0.0.0:8080"))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.POST("/api/send", sendMessage)
	h.GET("/api/pull", pullMessage)

	h.Spin()
}

func sendMessage(ctx context.Context, c *app.RequestContext) {
	var req api.SendRequest
	err := c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, "Failed to parse request body: %v", err)
		return
	}
	sender, senderStatus := c.GetQuery("sender")
	receiver, receiverStatus := c.GetQuery("receiver")
	text, textStatus := c.GetQuery("text")
	if !senderStatus || !receiverStatus || !textStatus {
		c.String(consts.StatusBadRequest, "Missing request query parameters.")
		return
	}
	chat := parseChat(sender, receiver)
	resp, err := cli.Send(ctx, &rpc.SendRequest{
		Message: &rpc.Message{
			Chat:     chat,
			Text:     text,
			Sender:   sender,
			SendTime: time.Now().UnixNano() / int64(time.Microsecond),
		},
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
	} else if resp.Code != 0 {
		c.String(consts.StatusInternalServerError, resp.Msg)
	} else {
		c.Status(consts.StatusOK)
	}
}

func pullMessage(ctx context.Context, c *app.RequestContext) {
	var req api.PullRequest
	err := c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, "Failed to parse request body: %v", err)
		return
	}
	chat, chatStatus := c.GetQuery("chat")
	if !chatStatus {
		c.String(consts.StatusBadRequest, "No chat found.")
		return
	}
	if len(strings.Split(chat, ":")) != 2 {
		c.String(consts.StatusBadRequest, "Wrong chat format.")
		return
	}
	chat = parseChatPull(chat)
	cursor, cursorStatus := c.GetQuery("cursor")
	var cursorNum int64
	if cursorStatus {
		cursorNum, err = strconv.ParseInt(cursor, 10, 64)
		if err != nil {
			c.String(consts.StatusBadRequest, "Failed to parse cursor: %v", err)
			return
		}
	}
	limit, limitStatus := c.GetQuery("limit")
	var limitNum int32
	if limitStatus {
		tempLimitNum, err := strconv.ParseInt(limit, 10, 32)
		if err != nil {
			c.String(consts.StatusBadRequest, "Failed to parse limit: %v", err)
			return
		}
		limitNum = int32(tempLimitNum)
	} else {
		limitNum = 10
	}
	reverse, reverseStatus := c.GetQuery("reverse")
	var reverseBool bool
	if reverseStatus {
		reverseBool, err = strconv.ParseBool(reverse)
		if err != nil {
			c.String(consts.StatusBadRequest, "Failed to parse reverse: %v", err)
			return
		}
	}

	resp, err := cli.Pull(ctx, &rpc.PullRequest{
		Chat:    chat,
		Cursor:  cursorNum,
		Limit:   limitNum,
		Reverse: &reverseBool,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	} else if resp.Code != 0 {
		c.String(consts.StatusInternalServerError, resp.Msg)
		return
	}
	messages := make([]*api.Message, 0, len(resp.Messages))
	for _, msg := range resp.Messages {
		messages = append(messages, &api.Message{
			Chat:     msg.Chat,
			Text:     msg.Text,
			Sender:   msg.Sender,
			SendTime: msg.SendTime,
		})
	}
	c.JSON(consts.StatusOK, &api.PullResponse{
		Messages:   messages,
		HasMore:    resp.GetHasMore(),
		NextCursor: resp.GetNextCursor(),
	})
}

func parseChat(a string, b string) string {
	if a < b {
		return a + ":" + b
	}
	return b + ":" + a
}

func parseChatPull(chat string) string {
	participants := strings.Split(chat, ":")
	return parseChat(participants[0], participants[1])
}
