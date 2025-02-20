package server

import (
	"errors"
	"sync"

	"github.com/clavera2/yellow_jacket/utils"
	"github.com/google/uuid"
)

type MessagePool struct {
	sync.Mutex
	messages map[uuid.UUID]utils.Message
}

func NewMessagePool() *MessagePool {
	return &MessagePool{Mutex: sync.Mutex{}, messages: make(map[uuid.UUID]utils.Message)}
}

func (m *MessagePool) AddMessage(message utils.Message) error {
	if m.IDExists(message.GetID()) {
		return errors.New("message already exists in message pools")
	}
	m.Lock()
	m.messages[message.GetID()] = message
	m.Unlock()
	return nil
}

func (m *MessagePool) GetAllMessages() []utils.Message {
	messageSlice := make([]utils.Message, len(m.messages))

	for _, message := range m.messages {
		messageSlice = append(messageSlice, message)
	}
	return messageSlice
}

func (m *MessagePool) DeleteMessage(id uuid.UUID) error {
	if !m.IDExists(id) {
		return errors.New("message with id does not exist")
	}
	delete(m.messages, id)
	return nil
}

func (m *MessagePool) ClearPool() {
	clear(m.messages)
}

func (m *MessagePool) IDExists(id uuid.UUID) bool {
	_, exists := m.messages[id]
	return exists
}
