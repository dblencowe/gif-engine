package database

import "context"

type ImageRecord interface {
	Url() string
}

type DB interface {
	Stop(context.Context) error
	Write(context.Context, any) error
	FindByTags(context.Context, []string) (ImageRecord, error)
}
