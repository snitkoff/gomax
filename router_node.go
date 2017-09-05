package gomax

type node struct {
	child map[uint64]*node
	route *route
	flag  bool
}
