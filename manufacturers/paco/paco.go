package paco

import (
	"sort"
)

type Paco struct{}

func NewPaco() *Paco {
	return &Paco{}
}

func (g *Paco) ListParts() []string {
	l := make([]string, 0, len(partsMap))

	for k := range partsMap {
		l = append(l, k)
	}

	sort.Strings(l)

	return l
}
