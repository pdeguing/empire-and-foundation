package data

import (
	"github.com/pdeguing/empire-and-foundation/ent"
)

func hasResources(p *ent.Planet, s Amounts) bool {
	return p.Metal >= s.Metal &&
		p.Hydrogen >= s.Hydrogen &&
		p.Silica >= s.Silica
}

func addStock(p *ent.Planet, s Amounts) {
	p.Metal += s.Metal
	p.Hydrogen += s.Hydrogen
	p.Silica += s.Silica
}

func subStock(p *ent.Planet, s Amounts) {
	p.Metal -= s.Metal
	p.Hydrogen -= s.Hydrogen
	p.Silica -= s.Silica
}