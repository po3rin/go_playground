//+build wireinject

package main

import "github.com/google/go-cloud/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
