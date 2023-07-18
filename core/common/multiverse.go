package common

import (
	"context"

	"github.com/spf13/afero"
)

type Multiverse interface {
	Exist(universe string) bool
	Universe(name string) Universe
	Ui() afero.Fs
	Context() context.Context
	ValidServices() []string
	ValidFixtures() []string
	ValidClients() []string
	Status() interface{}
}
