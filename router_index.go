package gomax

import "net/http"

type index struct {
	tree  *node
	index map[uint64]route
}

func newIndex() *index {
	return &index{
		index: make(map[uint64]route),
	}
}

func (i *index) find(req *http.Request) *route {
	salt := i.generateSalt(req.Method)
	sHashes := [URLSectionCount]uint64{}
	level := i.generateUintSlice(req.URL.Path, salt, &sHashes)
	return i.findX(level, 0, i.tree, &sHashes)
}

func (i *index) generateSalt(s string) uint64 {
	return uint64(s[0] + s[1])
}

func (i *index) generateUintSlice(s string, salt uint64, sHashes *[URLSectionCount]uint64) int {
	c := LimiterByte
	n := 0
	length := len(s)
	if length == 1 {
		sHashes[0] = SlashHash
		return 0
	}
	var total = salt
	for i := 1; i < length; i++ {
		if s[i] == c {
			sHashes[n] = total
			total = salt
			n++
			continue
		}
		total = total<<5 + uint64(s[i])
	}
	sHashes[n] = total
	n++
	return n
}

func (i *index) findX(ln int, level int, tree2 *node, cHashes *[URLSectionCount]uint64) *route {
	if ln == level {
		return tree2.route
	} else if z1, ok := tree2.child[cHashes[level]]; ok {
		return i.findX(ln, level+1, z1, cHashes)
	} else if z2, ok := tree2.child[LimiterUnit]; ok {
		return i.findX(ln, level+1, z2, cHashes)
	}
	return nil
}
