package gomax

import "net/http"

type index struct {
	tree  *node
	index map[uInt]route
}

func newIndex() *index {
	return &index{
		index: make(map[uInt]route),
	}
}

func (i *index) find(req *http.Request) *route {
	salt := i.generateSalt(req.Method)
	sHashes := [HTTP_SECTION_COUNT]uInt{}
	level := i.generateUintSlice(req.URL.Path, salt, &sHashes)
	return i.findX(level, 0, i.tree, &sHashes)
}

func (i *index) generateSalt(s string) uInt {
	return uInt(s[0] + s[1])
}

func (i *index) generateUintSlice(s string, salt uInt, cHashes *[HTTP_SECTION_COUNT]uInt) int {
	c := DELIMITER_BYTE
	na := 0
	length := len(s)
	if length == 1 {
		cHashes[0] = SLASH_HASH
		return 0
	}
	var total typeHash = salt
	for i := 1; i < length; i++ {
		if s[i] == c {
			cHashes[na] = total
			total = salt
			na++
			continue
		}
		total = total<<5 + typeHash(s[i])
	}
	cHashes[na] = total
	na++
	return na
}
