package postoffice

import "github.com/chainpusher/chainpusher/model"

type PostOffice interface {
	Deliver(transactions []*model.Transaction) error

	GetTransports() []Transport
}
