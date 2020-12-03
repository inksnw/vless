package common

type Runnable interface {

	Start() error
	Closable
}

type Closable interface {
	Close() error
}

