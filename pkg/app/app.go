package app

import "context"

type Lifecycle interface {
	Setup() (ShutdownFunc, error)
	Run() (ShutdownFunc, error)
}

type Config struct {
	Name     string
	Version  string
	BasePath string
}

type ShutdownFunc func(context.Context) error
