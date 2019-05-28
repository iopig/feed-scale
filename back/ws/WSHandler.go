package gateway

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/iopig/feed-scale/common"
)

// 每隔1秒, 检查一次连接是否健康
func (wsConnection *WSConnection) heartbeatChecker() {
	var (
		timer *time.Timer
	)
	timer = time.NewTimer(time.Duration(common.G_config.WsHeartbeatInterval) * time.Second)
	for {
		select {
		case <-timer.C:
			if !wsConnection.IsAlive() {
				wsConnection.Close()
				goto EXIT
			}
			timer.Reset(time.Duration(common.G_config.WsHeartbeatInterval) * time.Second)
		case <-wsConnection.closeChan:
			timer.Stop()
			goto EXIT
		}
	}

EXIT:
	// 确保连接被关闭
}

// 处理PING请求
func (wsConnection *WSConnection) handlePing(bizReq *common.BizMessage) (bizResp *common.BizMessage, err error) {
	var (
		buf []byte
	)

	wsConnection.KeepAlive()

	if buf, err = json.Marshal(common.BizPongData{}); err != nil {
		return
	}
	bizResp = &common.BizMessage{
		Type: "PONG",
		Data: json.RawMessage(buf),
	}
	return
}

// 处理websocket请求
func (wsConnection *WSConnection) WSHandle() {
	var (
		message *common.WSMessage
		bizReq  *common.BizMessage
		bizResp *common.BizMessage
		err     error
		buf     []byte
	)

	// 心跳检测线程
	go wsConnection.heartbeatChecker()

	// 请求处理协程
	for {
		if message, err = wsConnection.ReadMessage(); err != nil {
			goto ERR
		}

		// 只处理文本消息
		if message.MsgType != websocket.TextMessage {
			continue
		}

		// 解析消息体
		if bizReq, err = common.DecodeBizMessage(message.MsgData); err != nil {
			goto ERR
		}

		bizResp = nil

		// 1,收到PING则响应PONG: {"type": "PING"}, {"type": "PONG"}
		// 2,收到JOIN则加入ROOM: {"type": "JOIN", "data": {"room": "chrome-plugin"}}
		// 3,收到LEAVE则离开ROOM: {"type": "LEAVE", "data": {"room": "chrome-plugin"}}

		// 请求串行处理
		switch bizReq.Type {
		case "PING":
			if bizResp, err = wsConnection.handlePing(bizReq); err != nil {
				goto ERR
			}
		case "JOIN":
			if bizResp, err = wsConnection.handleJoin(bizReq); err != nil {
				goto ERR
			}
		case "LEAVE":
			if bizResp, err = wsConnection.handleLeave(bizReq); err != nil {
				goto ERR
			}
		}

		if bizResp != nil {
			if buf, err = json.Marshal(*bizResp); err != nil {
				goto ERR
			}
			// socket缓冲区写满不是致命错误
			if err = wsConnection.SendMessage(&common.WSMessage{websocket.TextMessage, buf}); err != nil {
				if err != common.ERR_SEND_MESSAGE_FULL {
					goto ERR
				} else {
					err = nil
				}
			}
		}
	}

ERR:
	// 确保连接关闭
	wsConnection.Close()

	// 离开所有房间

	// 从连接池中移除
	//	G_connMgr.DelConn(wsConnection)
	return
}
