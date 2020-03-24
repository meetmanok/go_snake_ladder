package domain

type Pawn struct {
	currentPos int8
}

func (p *Pawn) SetPos(pos int8) {
	p.currentPos = pos
}

func (p *Pawn) GetPos() int8 {
	return p.currentPos
}

func NewPawn() Pawn {
	return Pawn{currentPos:0}
}