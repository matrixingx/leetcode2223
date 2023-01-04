package lrucache

/*
	https://leetcode.cn/problems/lru-cache/
*/

type Node struct {
	Key,Val int
	Prev *Node
	Next *Node
}


type LRUCache struct {
	Mp map[int]*Node
	Head,Tail *Node
	Size,Cap int
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	lruCache := LRUCache{
		Mp: make(map[int]*Node, 0),
		Head: head,
		Tail: tail,
		Cap: capacity,
	}
	return lruCache
}


func (this *LRUCache) Get(key int) int {
	if v,ok := this.Mp[key];!ok {
		return -1
	} else {
		// 断开原双向链表
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		// 将节点放到最开头
		v.Prev = this.Head
		v.Next = this.Head.Next
		this.Head.Next.Prev = v
		this.Head.Next = v
		return v.Val
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if v,ok := this.Mp[key];ok {
		// 断开原双向链表
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		delete(this.Mp,v.Key)
	} else if this.Cap == this.Size {
		tail := this.Tail.Prev
		tail.Prev.Next = tail.Next
		tail.Next.Prev = tail.Prev
		delete(this.Mp,tail.Key)
	} else {
		this.Size++
	}
	node := &Node{
		Key: key,
		Val: value,
	}
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
	this.Mp[key] = node
}