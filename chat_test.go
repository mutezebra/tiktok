package tiktok

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	web2 "github.com/hertz-contrib/websocket"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"
)

func BenchmarkChatHandler(b *testing.B) {
	//registerUsers(b)
	info := loginUsers(b)
	conn := getWsConn(b, info)
	b.ResetTimer()

	msg := "{\n    \"msg_type\":1,\n    \"content\":\"just for test\"\n}"
	for i := 0; i < b.N; i++ {
		err := conn.WriteMessage(web2.TextMessage, []byte(msg))
		if err != nil {
			b.Error(err)
			return
		}
		_, _, err = conn.ReadMessage()
		if err != nil {
			b.Error(err)
			return
		}
	}

	_ = conn.Close()
}

var key = 15

func registerUsers(b *testing.B) {
	b.Helper()
	for i := 0; i < 2; i++ { // 可以确保接收消息的人也注册了，不然用户不存在话不能发消息
		url := "http://192.168.84.7:4000/api/user/register"
		method := "POST"
		id := strconv.Itoa(key + i)

		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		_ = writer.WriteField("user_name", "user"+id)
		_ = writer.WriteField("password", "123456")
		_ = writer.WriteField("email", "user"+id+"@gmail.com")
		_ = writer.Close()

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			b.Error(err)
			return
		}

		req.Header.Add("User-Agent", "PostmanRuntime/7.26.8")
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := client.Do(req)
		if err != nil {
			b.Error(err)
			return
		}

		_ = resp.Body.Close()
	}

}

type LoginResp struct {
	Base interface{} `json:"base"`
	Data struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

func loginUsers(b *testing.B) *LoginResp {
	b.Helper()
	url := "http://192.168.84.7:4000/api/user/login"
	method := "POST"
	id := strconv.Itoa(key)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("user_name", "user"+id)
	_ = writer.WriteField("password", "123456")
	_ = writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		b.Error(err)
		return nil
	}

	req.Header.Add("User-Agent", "PostmanRuntime/7.26.8")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		b.Error(err)
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()

	info := LoginResp{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		b.Error(err)
		return nil
	}
	return &info
}

func getWsConn(b *testing.B, info *LoginResp) *websocket.Conn {
	b.Helper()
	url := "ws://192.168.84.7:4000/api/relation/auth/chat?to=4480054134837248"
	header := http.Header{}
	header.Add("Access-Token", info.Data.AccessToken)
	header.Add("Refresh-Token", info.Data.RefreshToken)
	conn, resp, err := websocket.DefaultDialer.Dial(url, header)
	if errors.Is(err, websocket.ErrBadHandshake) {
		log.Printf("handshake failed with status %d", resp.StatusCode)
	}
	if err != nil {
		b.Error(err)
		return nil
	}

	return conn
}
