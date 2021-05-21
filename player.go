package main

type player struct {
	hp, maxhp int
	items []*item
}

func (p *player) spendItem(item *item) {
	for i := range p.items {
		if p.items[i] == item {
			p.items = append(p.items[:i], p.items[i+1:]...)
			return 
		}
	}
}
