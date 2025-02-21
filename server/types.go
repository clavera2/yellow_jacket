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

func (m MessagePool) GetMessage(id uuid.UUID) (utils.Message, error) {
	if !m.IDExists(id) {
		return utils.Message{}, errors.New("message pool does not contain message with id")
	} else {
		message := m.messages[id]
		return message, nil
	}
}

func (m *MessagePool) GetAllMessages() []utils.Message {
	messageSlice := make([]utils.Message, len(m.messages))
	m.Lock()
	for _, message := range m.messages {
		messageSlice = append(messageSlice, message)
	}
	m.Unlock()
	return messageSlice
}

func (m *MessagePool) DeleteMessage(id uuid.UUID) error {
	m.Lock()
	if !m.IDExists(id) {
		return errors.New("message with id does not exist")
	}
	delete(m.messages, id)
	m.Unlock()
	return nil
}

func (m *MessagePool) ClearPool() {
	m.Lock()
	clear(m.messages)
	m.Unlock()
}

func (m *MessagePool) IDExists(id uuid.UUID) bool {
	_, exists := m.messages[id]
	return exists
}

func (m MessagePool) EncodeAll() [][]byte {
	m.Lock()
	messages := m.GetAllMessages()
	encoded_messages := [][]byte{}
	for _, message := range messages {
		e, err := message.Encode()
		if err != nil {
			continue
		} else {
			encoded_messages = append(encoded_messages, e)
		}
	}
	m.Unlock()
	return encoded_messages
}
