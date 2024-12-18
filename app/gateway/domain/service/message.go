package chat

import (
	"time"

	"github.com/mutezebra/tiktok/app/gateway/domain/repository"
	"github.com/mutezebra/tiktok/pkg/consts"
)

type Message struct {
	MsgType   int8 `json:"msg_type"`
	HaveRead  bool
	PageSize  int8   `json:"page_size"`
	PageNum   int32  `json:"page_num"`
	Content   string `json:"content"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	timeSince int64
	timeEnd   int64
	GroupID   int64 `json:"group_id"`
	From      int64
	To        int64
	CreateAt  int64
}

const (
	ContentEmptyError = "content should not be empty"
	ContentNotEmpty   = "content should be empty"
	PageNumError      = "page_num should not smaller than 1"
	PageSizeError     = "page_size should not smaller than 1"
	TimeStartError    = "time_start format is wrong,eg: 2006-01-02 15:04:05"
	TimeEndError      = "time_end format is wrong,eg: 2006-01-02 15:04:05"

	UnSupportedType = "unsupported msg type"
)

// CheckMessageParams check the message params
func (m *Message) CheckMessageParams() (result string) {
	switch m.MsgType {
	case consts.ChatTextMessage:
		if m.Content == "" {
			return ContentEmptyError
		}

	case consts.HistoryChatMessage:
		if m.PageNum < 1 {
			return PageNumError
		}
		if m.PageSize < 1 {
			return PageSizeError
		}
		if m.Content != "" {
			return ContentNotEmpty
		}

		var t time.Time
		var err error

		if t, err = time.Parse("2006-01-02 15:04:05", m.TimeStart); err != nil {
			return TimeStartError
		}
		m.timeSince = t.Unix()

		if t, err = time.Parse("2006-01-02 15:04:05", m.TimeEnd); err != nil {
			return TimeEndError
		}
		m.timeEnd = t.Unix()

	case consts.NotReadMessage:

	default:
		return UnSupportedType
	}

	return ""
}

// buildHistoryQueryReq build the msg history query req for repository
func (m *Message) buildHistoryQueryReq() *repository.HistoryQueryReq {
	return &repository.HistoryQueryReq{
		From:     m.From,
		To:       m.To,
		PageNum:  m.PageNum,
		PageSize: m.PageSize,
		Start:    m.timeSince,
		End:      m.timeEnd,
	}
}

// buildRepoMessage convert the message to repository message
func (m *Message) buildRepoMessage() *repository.Message {
	return &repository.Message{
		UID:        m.From,
		ReceiverID: m.To,
		Content:    m.Content,
		HaveRead:   m.HaveRead,
		CreateAt:   m.CreateAt,
	}
}
