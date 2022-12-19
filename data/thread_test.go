package data

import "testing"

// Delete all threads from database
func ThreadDeleteAll() (err error) {
	statement := "delete from threads"
	_, err = Db.Exec(statement)
	return
}

// Delete all threads from database
func PostDeleteAll() (err error) {
	statement := "delete from posts"
	_, err = Db.Exec(statement)
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
}

func Test_Threads(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}

	threads, err := Threads()
	if err != nil {
		t.Error(err, "Cannot retrieve threads")
	}
	if len(threads) != 1 {
		t.Error(err, "Wrong number of threads retrieved")
	}
	if threads[0].Topic != conv.Topic {
		t.Error(err, "Different threads retrieved")
	}
}

func Test_CreatePost(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	post, err := users[0].CreatePost(conv, "My first post")
	if err != nil {
		t.Error(err, "Cannot create post")
	}
	if post.UserId != users[0].Id {
		t.Error("User not linked with post")
	}
}

func Test_Posts(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	post, err := users[0].CreatePost(conv, "My first post")
	if err != nil {
		t.Error(err, "Cannot create post")
	}
	posts, err := conv.Posts()
	if len(posts) != 1 {
		t.Error(err, "Wrong number of posts retrieved")
	}
	if posts[0].Body != post.Body {
		t.Error(err, "Different posts retrieved")
	}
}
