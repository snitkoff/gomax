package gomax

type node struct {
	child map[uInt]*node
	route *route
	flag  bool
}
