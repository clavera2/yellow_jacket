package utils

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Message struct {
	Id   uuid.UUID   `json:"id"`
	Data interface{} `json:"data"`
}

func (m Message) GetID() uuid.UUID {
	return m.Id
}

func NewMessage(message interface{}) (*Message, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("could not generate unique id for message")
	}
	return &Message{
		Id:   uid,
		Data: message,
	}, nil
}

func (m Message) GetData() interface{} {
	return m.Data
}

func (m *Message) Encode() ([]byte, error) {
	marshal, err := json.Marshal(m.Data)
	if err != nil {
		return nil, errors.New("cannot convert message data to json")
	}
	return marshal, nil
}
