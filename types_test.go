package ingo_test

import (
	"io"
)

type Logger interface {
	Debugf(message string, args ...interface{})
}

type NoopLogger struct {
	writer io.Writer
}

func (c *NoopLogger) Debugf(message string, args ...interface{}) {
}

type Storage interface {
	Set(key, value string)
	Get(key string) string
}

type InMemoryStorage struct {
	m      map[string]string
	logger Logger
}

func NewInMemoryStorage(logger Logger) *InMemoryStorage {
	return &InMemoryStorage{
		m:      make(map[string]string),
		logger: logger,
	}
}

func (s *InMemoryStorage) Set(key, value string) {
	s.logger.Debugf("InMemoryStorage:SET:%s:%s", key, value)
	s.m[key] = value
}

func (s *InMemoryStorage) Get(key string) string {
	s.logger.Debugf("InMemoryStorage:GET:%s", key)
	return s.m[key]
}
