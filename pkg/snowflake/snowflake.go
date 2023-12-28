package snowflake

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Generator struct {
	mu             sync.Mutex
	workerID       int64
	sequence       int64
	last           int64
	workerIDShift  uint64
	timestampShift uint64
	sequenceMask   int64
	startTime      int64
}

type Config struct {
	StartTime    int64
	WorkerID     int
	WorkerIDBits uint8
	SequenceBits uint8
}

func MustNew(workerID int) *Generator {
	g, err := New(workerID)
	if err != nil {
		panic(err)
	}
	return g
}

func New(workerID int) (*Generator, error) {
	return NewWithConfig(Config{
		WorkerID:     workerID,
		WorkerIDBits: 5,
		SequenceBits: 12,
		StartTime:    1619827200000,
	})
}

func MustNewWithConfig(conf Config) *Generator {
	g, err := NewWithConfig(conf)
	if err != nil {
		panic(err)
	}
	return g
}

func NewWithConfig(conf Config) (*Generator, error) {
	maxWorkerID := -1 ^ (-1 << conf.WorkerIDBits)
	if conf.WorkerID > int(maxWorkerID) {
		return nil, fmt.Errorf("worker id must be between 0 and %d", maxWorkerID)
	}

	workerIDShift := uint64(conf.SequenceBits)
	timestampShift := workerIDShift + uint64(conf.WorkerIDBits)
	sequenceMask := -1 ^ (-1 << conf.SequenceBits)

	return &Generator{
		startTime:      conf.StartTime,
		workerID:       int64(conf.WorkerID),
		sequence:       0,
		last:           -1,
		workerIDShift:  workerIDShift,
		timestampShift: timestampShift,
		sequenceMask:   int64(sequenceMask),
	}, nil
}

func (sf *Generator) ID() int64 {
	sf.mu.Lock()

	now := currentMillis()
	if now == sf.last {
		sf.sequence = (sf.sequence + 1) & sf.sequenceMask
		if sf.sequence == 0 {
			now = sf.NextMillis()
		}
	} else {
		sf.sequence = 0
	}

	sf.last = now

	id := (now-sf.startTime)<<sf.timestampShift |
		sf.workerID<<sf.workerIDShift |
		sf.sequence

	sf.mu.Unlock()
	return id
}

func (sf *Generator) NextMillis() int64 {
	now := currentMillis()
	for {
		if now > sf.last {
			break
		}
		now = currentMillis()
	}
	return now
}

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type ID int64

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id ID) Uint64() uint64 {
	return uint64(id)
}
