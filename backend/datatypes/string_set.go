package datatypes

type StringSet struct {
	entries map[string]int
	list    []string
}

func NewStringSet() *StringSet {
	return &StringSet{entries: make(map[string]int)}
}

func (set *StringSet) updateList() {
	var list []string
	i := 0

	for _, listEntry := range set.list {
		_, ok := set.entries[listEntry]

		if ok {
			list = append(list, listEntry)
			set.entries[listEntry] = i
			i++
		}
	}

	set.list = list
}

func (set *StringSet) Add(entry string) {
	_, ok := set.entries[entry]

	if !ok {
		set.entries[entry] = len(set.list)
		set.list = append(set.list, entry)
	}
}

func (set *StringSet) AddMultiple(entries []string) {
	for _, entry := range entries {
		set.Add(entry)
	}
}

func (set *StringSet) Remove(entry string) {
	if _, ok := set.entries[entry]; ok {
		delete(set.entries, entry)
		set.updateList()
	}
}

func (set *StringSet) RemoveMultiple(entries []string) {
	for _, entry := range entries {
		if _, ok := set.entries[entry]; ok {
			delete(set.entries, entry)
		}
	}

	set.updateList()
}

func (set *StringSet) Has(entry string) bool {
	_, ok := set.entries[entry]
	return ok
}

func (set *StringSet) Entries() []string {
	return set.list
}
