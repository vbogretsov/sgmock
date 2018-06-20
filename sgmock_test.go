package sgmock_test

import (
	"sync"
	"testing"

	"github.com/vbogretsov/sgmock"
)

const size = 100

func TestConcurrentSend(t *testing.T) {
	mock := sgmock.New()

	var sg sync.WaitGroup
	sg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			mock.Send(sgmock.Message{})
			sg.Done()
		}()
	}
	sg.Wait()

	l := len(mock.List())
	if l != size {
		t.Errorf("invalid messages count, expected %d but was %d", size, l)
	}
}
