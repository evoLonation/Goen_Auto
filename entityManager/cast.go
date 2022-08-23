package entityManager

type entityForCast interface {
	GetRealType() GoenInheritType
}

func (p *manager[T, PT]) CastFrom(e any) (PT, error) {
	//e.GetRealType()
	ei := (any(e)).(EntityForManager)
	return p.Get(ei.GetGoenId())
}

func (p *manager[T, PT]) GetRealType(e PT) GoenInheritType {
	return (any(e)).(EntityForInheritManager).GetRealType()
}
