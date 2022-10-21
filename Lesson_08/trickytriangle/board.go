package trickytriangle

import "fmt"

type Hole struct {
	empty bool
}

func (h *Hole) AddPeg() {
	h.empty = false
}

func (h *Hole) RemovePeg() {
	h.empty = true
}

func (h Hole) Empty() bool {
	return h.empty
}

func (h Hole) String() string {
	if h.empty {
		return "○"
	}
	return "●"
}

type Board struct {
	Holes [15]Hole
}

func (b Board) String() string {
	ret := ""
	ret += fmt.Sprintf("          %s\n", b.Holes[0])
	ret += fmt.Sprintf("        %s   %s\n", b.Holes[1], b.Holes[2])
	ret += fmt.Sprintf("      %s   %s   %s\n", b.Holes[3], b.Holes[4], b.Holes[5])
	ret += fmt.Sprintf("    %s   %s   %s   %s\n", b.Holes[6], b.Holes[7], b.Holes[8], b.Holes[9])
	ret += fmt.Sprintf("  %s   %s   %s   %s   %s\n", b.Holes[10], b.Holes[11], b.Holes[12], b.Holes[13], b.Holes[14])

	return ret
}

func (b Board) Remaining() int {
	ret := 0
	for _, v := range b.Holes {
		if !v.empty {
			ret++
		}
	}
	return ret
}
