package ingo_test

import "testing"
import "github.com/goenning/ingo"
import "reflect"

func TestUsingInteger(t *testing.T) {
	container := ingo.NewContainer()
	container.Register(reflect.TypeOf((*int)(nil)).Elem(), 2)

	add := func(a, b int) int {
		return a + b
	}

	results, err := container.Execute(add)
	if err != nil {
		t.Errorf("Err should be nil, but was %s", err)
	}
	if results[0] != 4 {
		t.Errorf("First result should be 4, but was %s", results[0])
	}
}

func TestUsingNoopLogger(t *testing.T) {
	container := ingo.NewContainer()
	logger := &NoopLogger{}
	container.Register(reflect.TypeOf((*Logger)(nil)).Elem(), logger)

	doIt := func(logger Logger) {
		logger.Debugf("This will panic if logger is nil")
	}

	results, err := container.Execute(doIt)
	if err != nil {
		t.Errorf("Err should be nil, but was %s", err)
	}
	if len(results) != 0 {
		t.Errorf("Results should be empty, but was %s", results)
	}
}

func TestUsingNestedDependencies(t *testing.T) {
	container := ingo.NewContainer()
	logger := &NoopLogger{}
	container.Register(reflect.TypeOf((*Logger)(nil)).Elem(), logger)
	container.Register(reflect.TypeOf((*Storage)(nil)).Elem(), NewInMemoryStorage)

	doIt := func(logger Logger, storage Storage) string {
		storage.Set("message", "Hello World")
		return storage.Get("message")
	}

	results, err := container.Execute(doIt)
	if err != nil {
		t.Errorf("Err should be nil, but was %s", err)
	}
	if results[0] != "Hello World" {
		t.Errorf("First result should be 'Hello World', but was %s", results[0])
	}
}
