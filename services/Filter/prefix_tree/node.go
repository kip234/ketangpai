package prefix_tree

import "sync"

type node struct {
	lock sync.RWMutex
	unit byte
	sons map[byte]*node
}

func newNode(c byte) *node {
	return &node{
		unit: c,
		sons:make(map[byte]*node),
	}
}

func (n *node)IsTrailer()bool{
	if len(n.sons)==0{
		return true
	}
	return false
}

func (n *node)GetUnit()byte{
	return n.unit
}

func (n *node)GetSon(s byte)(re *node,ok bool){
	n.lock.RLock()
	re,ok=n.sons[s]
	n.lock.RUnlock()
	return
}

func (n *node)IsExit(b byte) (ok bool) {
	n.lock.RLock()
	_,ok=n.sons[b]
	n.lock.RUnlock()
	return
}

func (n *node)Add(b *node){
	n.lock.Lock()
	n.sons[b.GetUnit()]=b
	n.lock.Unlock()
}