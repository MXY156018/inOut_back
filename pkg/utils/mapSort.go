package utils

type MapSorter []Item
type Item struct {
	Key string
	Val string
}

func NewMapSorter(m map[string]string) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter)Less(i,j int) bool {
	return ms[i].Key < ms[j].Key
}

func (ms MapSorter) Swap(i,j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

