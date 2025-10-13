package api

import "errors"

const (
	initialCapacity = 50
)

// TODO(atoaidoocr7): Look into how we can make these fields private. For now, the fields must be public
// inorder to be able to use them with golangs net/rpc pacakge. The net/rpc package requires fields to be
// exported in order to encode and decode request parameters
type LogEntry struct {
	Term  int64
	Index int64
	Data  []byte
}

type Log struct {
	entries []*LogEntry
}

func NewLog() *Log {
	return &Log{entries: make([]*LogEntry, 0, initialCapacity)}
}

func (l *Log) AddEntry(entry *LogEntry) {
	l.entries = append(l.entries, entry)
}

func (l *Log) GetEntry(index int64) (LogEntry, error) {
	if index < 0 || index >= int64(len(l.entries)) {
		return LogEntry{}, errors.New("index out of range")
	}
	return *l.entries[index], nil
}

func (l *Log) Len() int {
	return len(l.entries)
}

func (l *Log) Pop() bool {
	if len(l.entries) == 0 {
		return false
	}
	return true
}

func (l *Log) Peek() LogEntry {
	if len(l.entries) == 0 {
		return LogEntry{}
	}
	return *l.entries[0]
}
