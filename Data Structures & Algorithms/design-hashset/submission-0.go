type MyHashSet struct {
	table [1000000]bool
}

func Constructor() MyHashSet {
    return MyHashSet{table: [1000000]bool{}}
}

func (this *MyHashSet) Add(key int) {
    this.table[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.table[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.table[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 