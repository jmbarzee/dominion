package connect

import "time"

type connectionConfig struct {
	timeout time.Duration
}

func NewConnectionConfig(timeout time.Duration) ConnectionConfig {
	return connectionConfig{
		timeout: timeout,
	}
}

func (cc connectionConfig) GetTimeout() time.Duration {
	return cc.timeout
}
