package lrucache

import "container/list"

// 412ms 75MB
type ListValue struct {
	key   int
	value int
}
type LRUCache struct {
	Capacity int
	List     *list.List
	Map      map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		List:     list.New(),
		Map:      make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Map[key]
	if ok {
		ret := v.Value.(ListValue)
		this.List.MoveToFront(v)
		return ret.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.Map[key]
	if ok {
		v.Value = ListValue{key: key, value: value}
		this.List.MoveToFront(v)
		return
	}

	if this.List.Len() >= this.Capacity {
		back := this.List.Back()
		delete(this.Map, back.Value.(ListValue).key)
		this.List.Remove(back)
	}

	v = this.List.PushFront(ListValue{key: key, value: value})
	this.Map[key] = v
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
