package main

import (
	"container/ring"
	"fmt"
)

func main() {
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
	 p := s.Prev()
	 // Note: Cannot use multiple assignment because
	 // evaluation order of LHS is not specified.
	 r.next = s
	 s.prev = r
	 n.prev = p
	 p.next = n
	}
	return n
   }
}