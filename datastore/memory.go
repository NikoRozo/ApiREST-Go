package datastore

import (
	"log"
	"os"
	"strings"

	"github.com/nikorozo/apirest/loader"
)

type Customer struct {
	Store *[]*loader.CustomerData `json:"store"`
}

func (b *Customer) Initialize() {
	filename := "./assets/customer.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

func (b *Customer) SearchCustomer(name string, limit, skip int) *[]*loader.CustomerData {
	ret := Filter(b.Store, func(v *loader.CustomerData) bool {
		return strings.Contains(strings.ToLower(v.Name), strings.ToLower(name))
	})
	if limit == 0 || limit > len(*ret) {
		limit = len(*ret)
	}

	data := (*ret)[skip:limit]
	return &data
}

func (b *Customer) SearchId(id string) *loader.CustomerData {
	ret := Filter(b.Store, func(v *loader.CustomerData) bool {
		return strings.ToLower(v.ID) == strings.ToLower(id)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

func (b *Customer) CreateCustomer(customer *loader.CustomerData) bool {
	*b.Store = append(*b.Store, customer)
	return true
}

func Filter(vs *[]*loader.CustomerData, f func(*loader.CustomerData) bool) *[]*loader.CustomerData {
	vsf := make([]*loader.CustomerData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
