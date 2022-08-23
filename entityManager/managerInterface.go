package entityManager

type InheritManagerForOther[T any] interface {
	ManagerForOther[T]
	CastFrom(any) (T, error)
	GetRealType(T) GoenInheritType
}

type ManagerForOther[T any] interface {
	New() T
	GetFromAllInstanceBy(member string, value any) (T, error)
	FindFromAllInstanceBy(member string, value any) ([]T, error)
	AddInAllInstance(e T)
	RemoveFromAllInstance(e T)
}

type ManagerForEntity[PT any] interface {
	// Get 实际上不需要检查是否在allinstance里面
	Get(goenId int) (PT, error)
	FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error)
	GetGoenId(PT) int
}
