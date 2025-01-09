package wscore

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type WsShellClient struct {
	client *websocket.Conn
}

func NewWsShellClient(client *websocket.Conn) *WsShellClient {
	return &WsShellClient{client: client}
}

func (this *WsShellClient) Write(p []byte) (n int, err error) {
	//这里做了改动
	err = this.client.WriteMessage(websocket.TextMessage,
		p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
func (this *WsShellClient) Read(p []byte) (n int, err error) {
	// 读取前端传来的数据
	_, b, err := this.client.ReadMessage()

	if err != nil {
		return 0, err
	}
	fmt.Print("read from client: ", string(b))
	return copy(p, string(b)+"\n"), nil
}
