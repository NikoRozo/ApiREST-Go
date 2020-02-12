package datastore

import "github.com/nikorozo/apirest/loader"

type CustomerStore interface {
	Initialize()
	SearchCustomer(CustomerName string, limit, skip int) *[]*loader.CustomerData
	SearchId(id string) *loader.CustomerData
	CreateCustomer(customer *loader.CustomerData) bool
}
