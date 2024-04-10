package chat

type Message struct {
	MsgType  int8   `json:"msg_type"`
	Content  string `json:"content"`
	CreateAt int64
	From     int64
	To       int64
}

func (m *Message) SetCreateAt(at int64) {
	m.CreateAt = at
}

func (m *Message) SetFrom(from int64) {
	m.From = from
}

func (m *Message) SetTo(to int64) {
	m.To = to
}
