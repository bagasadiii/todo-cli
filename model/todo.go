package model

import (
	"sync"
	"time"
)
var (
	lastID int
	mu sync.Mutex
)
type Todo struct {
	TodoID		int
	Name		string
	Status		bool
	CreatedAt	time.Time
	CompletedAt	time.Time
}
func(t *Todo)AddID(){
	mu.Lock()
	defer mu.Unlock()

	lastID++
	t.TodoID = lastID
}