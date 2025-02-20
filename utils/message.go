package utils

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Message struct {
	id   uuid.UUID
	data interface{}
}

func (m Message) GetID() uuid.UUID {
	return m.id
}

func NewMessage(message interface{}) (*Message, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("could not generate unique id for message")
	}
	return &Message{
		id:   uid,
		data: message,
	}, nil
}

func (m Message) GetData() interface{} {
	return m.data
}

func (m *Message) Encode() ([]byte, error) {
	marshal, err := json.Marshal(m.data)
	if err != nil {
		return nil, errors.New("cannot convert message data to json")
	}
	return marshal, nil
}
