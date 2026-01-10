package reading_files_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/don-da/study-go/learn-go-with-tests/reading_files"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md": {Data: []byte("hi")},
	}

	posts, err := reading_files.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d, want %d", len(posts), len(fs))
	}
}
