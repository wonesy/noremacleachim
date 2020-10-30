package main

import (
	"errors"

	"github.com/wonesy/noremacleachim/manufacturers/paco"
)

type Arbiter struct {
	manufacturers map[string]Manufacturer
}

func NewArbiter() *Arbiter {
	return &Arbiter{}
}

// Manufacturer represents a pump manufacturer
type Manufacturer interface {
	ListParts() []string
}

var manufacturerNames = []string{
	"paco",
}

func (a *Arbiter) InitManufacturers() {
	a.manufacturers = make(map[string]Manufacturer, len(manufacturerNames))

	for _, mf := range manufacturerNames {
		switch mf {
		case "paco":
			a.manufacturers[mf] = paco.NewPaco()
		}
	}
}

func (a *Arbiter) ListManufacturers() []string {
	return manufacturerNames
}

func (a *Arbiter) GetManufacturer(name string) (Manufacturer, error) {
	mf, ok := a.manufacturers[name]
	if !ok {
		return nil, errors.New("No such manufacturer: " + name)
	}
	return mf, nil
}
