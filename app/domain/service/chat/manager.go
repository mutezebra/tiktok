package chat

import (
	"fmt"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"sync"
)

type Manager struct {
	connMap map[int64]*websocket.Conn
	msgChan map[int64]chan []byte
	repo    repository.ChatRepository
	mu      sync.RWMutex
}

func (m *Manager) Send(content []byte, to int64) error {
	m.mu.RLock()
	ch, ok := m.msgChan[to]
	m.mu.RUnlock()
	if !ok {
		return errors.New(fmt.Sprintf("could not find %d`s channel", to))
	}

	ch <- content
	return nil
}

func (m *Manager) Register(from int64, conn *websocket.Conn, ch chan []byte) {
	m.connMap[from] = conn
	m.msgChan[from] = ch
}

func (m *Manager) Unregister(from int64) {
	m.mu.Lock()
	delete(m.connMap, from)
	delete(m.msgChan, from)
	m.mu.Unlock()
}
