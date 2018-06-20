package sgmock

import (
	"sync"
)

// Address represents an email address.
type Address struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// Content represents SendGrid V3 API email body.
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Personalizations represents SendGrid V3 API personalizations.
type Personalizations struct {
	Subject string    `json:"subject"`
	To      []Address `json:"to"`
	Cc      []Address `json:"cc,omitempty"`
	Bcc     []Address `json:"bcc,omitempty"`
}

// Message represents SendGrid V3 API message.
type Message struct {
	From             Address            `json:"from"`
	Personalizations []Personalizations `json:"personalizations"`
	Content          Content            `json:"content"`
}

// Mock represents SendGrid V3 test API.
type Mock interface {
	Send(Message) error
	List() []Message
	Clear()
}

type mock struct {
	mutex sync.Mutex
	store []Message
}

// New creates a new API.
func New() Mock {
	return &mock{store: []Message{}}
}

func (m *mock) Send(msg Message) error {
	m.mutex.Lock()
	m.store = append(m.store, msg)
	m.mutex.Unlock()
	return nil
}

func (m *mock) List() []Message {
	return m.store
}

func (m *mock) Clear() {
	m.mutex.Lock()
	m.store = []Message{}
	m.mutex.Unlock()
}
