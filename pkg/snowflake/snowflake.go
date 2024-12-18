package snowflake

import (
	"sync"
	"time"

	"github.com/mutezebra/tiktok/pkg/consts"
)

var (
	snow *snowflake
	once sync.Once
)

func GenerateID(workerID int64, centerID int64) int64 {
	once.Do(initSnow(workerID, centerID))
	return snow.generateID()
}

type snowflake struct {
	sync.Mutex
	timestamp    int64
	workerID     int64
	dataCenterID int64
	sequence     int64
}

func (s *snowflake) generateID() int64 {
	s.Lock()
	now := time.Now().UnixMilli()
	if s.timestamp == now {
		// Increment the sequence number and ensure it does not exceed its maximum value by applying a bitmask.
		s.sequence = (s.sequence + 1) & consts.SequenceMask
		// If the sequence number has reached its maximum value this loop waits until the next millisecond so that
		// a new timestamp can be generated.
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}
	t := now - consts.Epoch
	if t > consts.TimestampMax {
		s.Unlock()
		return 0
	}
	s.timestamp = now
	id := (t << consts.TimestampShift) | (s.dataCenterID << consts.DataCenterIDShift) | (s.workerID << consts.WorkerIDShift) | (s.sequence)
	s.Unlock()
	return id
}

func initSnow(wid, cid int64) func() {
	return func() {
		workerID := wid
		if workerID > consts.WorkerIDMax {
			panic("workerID exceeds its maximum value")
		}
		dataCenterID := cid
		if dataCenterID > consts.WorkerIDMax {
			panic("dataCenterID exceeds its maximum value")
		}
		snow = &snowflake{
			Mutex:        sync.Mutex{},
			timestamp:    0,
			workerID:     workerID,
			dataCenterID: dataCenterID,
			sequence:     0,
		}
	}
}
