package main

// var hash = 133

// type doubled_list struct {
// 	prev *doubled_list
// 	next *doubled_list
// 	key  int
// 	data int
// }

// type list struct {
// 	next *list
// 	data *doubled_list
// }

// type HashMap struct {
// 	heads [133]*list
// }

// type LRUCache struct {
// 	cap    int
// 	lHead  *doubled_list
// 	lTail  *doubled_list
// 	hm     HashMap
// 	length int
// }

// func (hm HashMap) get(key int) *doubled_list {
// 	num := key % hash
// 	head := hm.heads[num]
// 	for head != nil {
// 		if head.data.key == key {
// 			return head.data
// 		}
// 	}
// 	return nil
// }

// func (hm *HashMap) put(da *doubled_list) {
// 	num := da.key % hash
// 	head := hm.heads[num]
// 	temp := &list{
// 		next: head,
// 		data: da,
// 	}
// 	hm.heads[num] = temp
// }

// func (hm *HashMap) delete(da *doubled_list) {
// 	num := da.key % hash
// 	head := hm.heads[num]
// 	if head.data == da {
// 		hm.heads[num] = head.next
// 		return
// 	}
// 	for head.next != nil {
// 		if head.next.data == da {
// 			head.next = head.next.next
// 		}
// 	}
// }

// func Constructor(capacity int) LRUCache {
// 	lru := LRUCache{
// 		cap:    capacity,
// 		length: 0,
// 		lHead:  &doubled_list{key: -1, data: -1}, // 虚伪的头结点
// 		lTail:  &doubled_list{key: -2, data: -2},
// 	}
// 	lru.lHead.next = lru.lTail
// 	lru.lTail.prev = lru.lHead
// 	return lru
// }

// // 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// func (this *LRUCache) Get(key int) int {
// 	if this.length == 0 {
// 		return -1
// 	}
// 	data := this.hm.get(key)
// 	if data == nil {
// 		return -1
// 	}

// 	// 更新链表
// 	data.next.prev = data.prev
// 	data.prev.next = data.next
// 	// bug
// 	data.next = this.lHead.next
// 	data.prev = this.lHead
// 	this.lHead.next = data
// 	data.next.prev = data
// 	return data.data
// }

// // 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
// // 如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// func (this *LRUCache) Put(key int, value int) {
// 	data := this.hm.get(key)
// 	if data != nil {
// 		// 记得更新
// 		data.data = value
// 		data.prev.next = data.next
// 		data.next.prev = data.prev
// 		// 这一部分双向链表更新要考虑清楚
// 		data.next = this.lHead.next
// 		data.prev = this.lHead
// 		this.lHead.next = data
// 		data.next.prev = data
// 	} else {
// 		if this.cap == this.length {
// 			delete_data := this.lTail.prev
// 			delete_data.prev.next = this.lTail
// 			this.lTail.prev = delete_data.prev
// 			this.hm.delete(delete_data)
// 			this.length--
// 		}
// 		new_data := &doubled_list{
// 			prev: this.lHead,
// 			next: this.lHead.next,
// 			data: value,
// 			key:  key,
// 		}
// 		this.lHead.next = new_data
// 		new_data.next.prev = new_data
// 		this.hm.put(new_data)
// 		this.length++
// 	}
// }
