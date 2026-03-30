type LinkedList struct {
    head *Node
}

type Node struct {
    value int
    next *Node
}

func NewLinkedList() *LinkedList {
    return &LinkedList{nil}
}

func (ll *LinkedList) Get(index int) int {
    node := ll.get(index)
    if node == nil {
        return -1
    }
    return node.value
}

func (ll *LinkedList) get(index int) *Node {
    curr := 0
    node := ll.head

    for node != nil && curr < index {
        curr++
        node = node.next
    }
    
    if index < 0 {
        return nil
    }

    return node
}

func (ll *LinkedList) InsertHead(val int) {
    node := &Node{
        val,
        ll.head,
    }
    ll.head = node
}

func (ll *LinkedList) InsertTail(val int) {
    if ll.head == nil {
        ll.InsertHead(val)
        return
    }
    node := ll.walk(ll.head)
    node.next = &Node{
        val,
        nil,
    }
}

func (ll *LinkedList) walk(node *Node) *Node {
     if node.next != nil {
        return ll.walk(node.next)
     }

     return node
}


func (ll *LinkedList) Remove(index int) bool {
    if index == 0 {
        if ll.head == nil {
            return false
        }
        ll.head = ll.head.next
        return true
    }

    prevnode := ll.get(index - 1)
    if prevnode == nil || prevnode.next == nil {
        return false
    }
    
    node := prevnode.next
    prevnode.next = node.next

    return true
}

func (ll *LinkedList) GetValues() []int {
    values := []int{
    }
    node := ll.head

    for node != nil {
        values = append(values, node.value)
        node = node.next
    }

    return values
}