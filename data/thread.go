package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	// 仮データ
	thread := Thread{
		Id:        1,
		Uuid:      "testUuid",
		Topic:     "testTopic",
		UserId:    1,
		CreatedAt: time.Now(),
	}
	threads = append(threads, thread)
	return
}
