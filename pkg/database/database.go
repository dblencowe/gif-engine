package database

import "context"

type ImageRecord interface{}

type DB interface {
	Stop(context.Context) error
	Write(context.Context, any) error
	Read(context.Context, any) (ImageRecord, error)
}
