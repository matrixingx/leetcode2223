package solution


type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	Cap        int
	Mp         map[int]*ListNode
	Head, Tail *ListNode
}

func InitLRUCache(cap int) *LRUCache {
	head := &ListNode{}
	tail := &ListNode{}
	head.Next = tail
	tail.Prev = head
	return &LRUCache{
		Cap:  cap,
		Mp:   make(map[int]*ListNode),
		Head: head,
		Tail: tail,
	}
}

func (this *LRUCache) Push(v int) {
	var delTail = len(this.Mp) == this.Cap
	if node, ok := this.Mp[v]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = this.Head
		node.Next = this.Head.Next
	} else {
		newNode := &ListNode{
			Val:  v,
			Prev: this.Head,
			Next: this.Head.Next,
		}
		this.Mp[v] = newNode
	}
	if delTail {
		node := this.Tail.Prev
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = nil
		node.Next = nil
		delete(this.Mp, node.Val)
	}
}

func (this *LRUCache) Get(v int) *ListNode {
	if node, ok := this.Mp[v]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = this.Head
		node.Next = this.Head.Next
		return node
	} else {
		return nil
	}
}