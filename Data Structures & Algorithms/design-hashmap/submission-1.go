type MyHashMap struct {
	table []int
}

func Constructor() MyHashMap {
    t := make([]int, 1000001)
	for i := range t {
		t[i] = -1
	}
	return MyHashMap{table: t}
}

func (this *MyHashMap) Put(key int, value int) {
    this.table[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.table[key]
}

func (this *MyHashMap) Remove(key int) {
    this.table[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */