package components

type InventoryComponent struct {
	items map[string]Item
}

func (ic InventoryComponent) AddItem(i Item) {
	ic.items[i.Name()] = i
}

func (ic InventoryComponent) RemoveItem(name string) {
	delete(ic.items, name)
}

type Item interface {
	Name() string
	Use()
	Weight() int
}
