package main

type player struct {
	hp, maxhp int
	maxItems int
	items []*item
}

func (p *player) addItem(item *item) {
	if item.itemConsumableType == ITEM_AMMO {
		for i := range p.items {
			if p.items[i].itemConsumableType == ITEM_AMMO {
				p.items[i].count += item.count
				return
			}
		}
	}
	p.items = append(p.items, item)
}


func (p *player) spendItem(item *item) {
	for i := range p.items {
		if p.items[i] == item {
			if p.items[i].count > 1 {
				p.items[i].count--
				return
			}
			p.items = append(p.items[:i], p.items[i+1:]...)
			return
		}
	}
}

func (p *player) hasAmmo() bool {
	for i := range p.items {
		if p.items[i].itemConsumableType == ITEM_AMMO {
			return true
		}
	}
	return false
}
