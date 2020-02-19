package gokv

import "go.uber.org/zap"

// Option func
type Option func(*memory) error

// WithLogger is Option func which sets the instance of logger
func WithLogger(logger *zap.Logger) Option {
	return func(m *memory) error {
		m.logger = logger
		return nil
	}
}

// WithKeys is Option func which sets the instance of Keys
func WithKeys(keys Keys) Option {
	return func(m *memory) error {
		m.keys = keys
		return nil
	}
}
