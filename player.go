package main

type player struct {
	hp, maxhp int
	maxItems int
	items []*item
}

func (p *player) addItem(itemToAdd *item) {
	if itemToAdd.isAmmo() {
		for i := range p.items {
			if p.items[i].isAmmo() {
				p.items[i].count += itemToAdd.count
				return
			}
		}
	}
	if itemToAdd.isWeapon() && len(p.items) > 0 {
		for i, ii := range p.items {
			if !ii.isWeapon() || (ii.weaponInfo.weaponType == itemToAdd.weaponInfo.weaponType &&
			ii.weaponInfo.damage < itemToAdd.weaponInfo.damage) {
				// insert into slice
				temp := append([]*item{}, p.items[i:]...)
				p.items = append(p.items[0:i], itemToAdd)
				p.items = append(p.items, temp...)
				return
			}
		}
	}
	p.items = append(p.items, itemToAdd)
}


func (p *player) spendItem(item *item, g *game) {
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
		if p.items[i].isAmmo() {
			return true
		}
	}
	return false
}

func (p *player) spendAmmo() {
	for i := range p.items {
		if p.items[i].isAmmo() {
			p.items[i].count--
			if p.items[i].count == 0 {
				p.spendItem(p.items[i], nil)
			}
			return
		}
	}
}
