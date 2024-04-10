package chat

type Message struct {
	MsgType  int8   `json:"msg_type"`
	Content  string `json:"content"`
	From     int64
	To       int64
	CreateAt int64
}
