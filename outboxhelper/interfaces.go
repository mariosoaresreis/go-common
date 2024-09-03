package outboxhelper

import "github.com/mariosoaresreis/go-common/datamodels"

type IdTypes interface {
	string | int64
}

type OutboxTypes interface {
	GetOutboxTypeInfo() (typ datamodels.AggregateType, version string)
}
