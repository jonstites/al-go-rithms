package main

// HashSet implements standard HashSet datastructure
type HashSet struct {
	data map[interface{}]interface{}
}

func newHashSet() HashSet {
	return HashSet{make(map[interface{}]interface{})}
}

func (set *HashSet) add(item interface{}) bool {
	_, ok := set.data[item]
	if ok == true {
		return false
	}

	set.data[item] = true
	return true
}

func (set *HashSet) isEmpty() bool {
	return len(set.data) == 0
}

func (set *HashSet) contains(item interface{}) bool {
	_, ok := set.data[item]
	return ok
}

func (set *HashSet) size() int {
	return len(set.data)
}

func (set *HashSet) clear() {
	set.data = make(map[interface{}]interface{})
}

func (set *HashSet) remove(item interface{}) bool {
	_, ok := set.data[item]
	if ok == false {
		return false
	}
	delete(set.data, item)
	return true
}