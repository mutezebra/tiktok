package chat

import (
	"context"
	"fmt"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/hertz-contrib/websocket"
	"time"
)

type Client struct {
	srv     *Service
	From    int64
	To      int64
	Channel chan []byte
	Conn    *websocket.Conn
}

func (c *Client) Register(ctx context.Context) error {
	return c.srv.Register(ctx, c)
}

func (c *Client) Unregister() error {
	return c.srv.Unregister(c)
}

func (c *Client) Read() {
	go func(c *Client) {
		defer func() {
			if err := c.Unregister(); err != nil {
				log.LogrusObj.Error(err)
			}
		}()

		for {
			mt, data, err := c.Conn.ReadMessage()
			if err != nil {
				log.LogrusObj.Error(err)
				break
			}

			switch mt {
			case websocket.CloseMessage:
				break
			case websocket.TextMessage:
				var msg Message
				if err = json.Unmarshal(data, &msg); err != nil {
					log.LogrusObj.Error(err)

					if err = c.srv.WriteErrorToConn("you message format is wrong", c.From); err != nil {
						log.LogrusObj.Error(err)
					}
					break
				}
				msg.From = c.From
				msg.To = c.To
				msg.CreateAt = time.Now().Unix()

				if err = c.srv.SendMsg(&msg); err != nil {
					log.LogrusObj.Error(err)
					break
				}
				continue
			default:
				if err = c.Conn.WriteMessage(websocket.CloseMessage, []byte("closed")); err != nil {
					log.LogrusObj.Error(err)
				}
				break
			}
		}
	}(c)
}

func (c *Client) Write() {
	for {
		data, ok := <-c.Channel
		if ok {
			if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.LogrusObj.Error(fmt.Sprintf("websocket conn write failed,cause: %v", err))
			}
		} else {
			break
		}
	}
}
