package main

type CharSet struct {
	Count    int
	Char     string
	WasMoved bool
}

func NewCharSet(count int, char string, wasMoved bool) *CharSet {
	return &CharSet{Count: count, Char: char, WasMoved: wasMoved}
}

func (p CharSet) Clone() *CharSet {
	return &CharSet{Count: p.Count, Char: p.Char, WasMoved: p.WasMoved}
}
