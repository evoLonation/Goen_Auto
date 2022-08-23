package entityManager

import (
	"database/sql"
	"errors"
	"fmt"
)

type EntityForInheritManager interface {
	entityForRecur
	EntityForManager
	GetRealType() GoenInheritType
}

type entityForRecur interface {
	GetParentEntity() EntityForInheritManager
	// 下两个方法用于recurInheritAfterNew
	// 最底层会调用
	inheritAfterNew(goenId int, inheritType GoenInheritType)
	// 除了最底层的每层会调用
	setCreated()
	// 下两个方法用于recurInheritAfterFind
	// 最底层会调用
	afterFind()
	// 除了最底层的每层会调用
	setExistent()
}

// InheritManagerForRecur 由于继承管理器需要处理本实体和本实体继承了的父实体，因此需要知道父实体对应管理器的某些方法来递归调用
type InheritManagerForRecur interface {
	getParentManager() InheritManagerForRecur
	getAllTables() []string
	recurAddInQueue(layer EntityForInheritManager)
	recurAfterFind(layer entityForRecur)
	recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer entityForRecur)
	generateGoenId() int
}

// 以下是Manager扩展的方法用于基实体
func (p *manager[T, PT]) getParentManager() InheritManagerForRecur {
	return nil
}
func (p *manager[T, PT]) getAllTables() []string {
	return []string{p.tableName}
}
func (p *manager[T, PT]) recurAddInQueue(layer EntityForInheritManager) {
	p.addInQueue(layer.(EntityForInheritManager))
}
func (p *manager[T, PT]) recurAfterFind(layer entityForRecur) {
	layer.afterFind()
}
func (p *manager[T, PT]) recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer entityForRecur) {
	layer.inheritAfterNew(goenId, inheritType)
}

// 以下是InheritManager的实现
//type inheritManagerTypeParam[T any] interface {
//	*T
//	EntityForInheritManager
//}

type inheritManager[T any, PT any] struct {

	// 如果该类是基类，可以借用之前实现的管理器的方法
	*manager[T, PT]
	// 如果该类不是基类，可以借用父辈管理器的方法
	parentManager InheritManagerForRecur

	GoenInheritType
}

func NewInheritManager[T any, PT any](tableName string, parentManager InheritManagerForRecur, inheritType GoenInheritType) (*inheritManager[T, PT], error) {
	_, ok := (any(new(T))).(PT)
	if !ok {
		return nil, errors.New("the type value T does not implement PT ")
	}
	_, ok = (any(new(T))).(EntityForInheritManager)
	if !ok {
		return nil, errors.New("the type value T does not implement EntityForInheritManager ")
	}
	manager := &manager[T, PT]{}
	manager.tableName = tableName
	ret := &inheritManager[T, PT]{
		GoenInheritType: inheritType,
		parentManager:   parentManager,
		manager:         manager,
	}
	return ret, nil
}

func (p *inheritManager[T, PT]) getInterface(e *T) EntityForInheritManager {
	return (any(e)).(EntityForInheritManager)
}

func (p *inheritManager[T, PT]) getPT(ei EntityForInheritManager) PT {
	return (any(ei)).(PT)
}

func (p *inheritManager[T, PT]) New() PT {
	e := p.getInterface(new(T))
	p.recurInheritAfterNew(p.generateGoenId(), p.GoenInheritType, e)
	p.recurAddInQueue(e)
	return p.getPT(e)
}
func (p *inheritManager[T, PT]) generateGoenId() int {
	return p.getParentManager().generateGoenId()
}

func (p *inheritManager[T, PT]) recurAddInQueue(layer EntityForInheritManager) {
	// 顺序要按照父实体在前子实体在后，否则容易出现先insert子实体但是没有对应父实体从而插入失败的情况
	p.getParentManager().recurAddInQueue(layer.GetParentEntity())
	p.addInQueue(layer)
}
func (p *inheritManager[T, PT]) recurAfterFind(layer entityForRecur) {
	p.getParentManager().recurAfterFind(layer.GetParentEntity())
	// 这里只是要初始化该layer的fieldChange
	layer.setExistent()
}
func (p *inheritManager[T, PT]) recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer entityForRecur) {
	p.getParentManager().recurInheritAfterNew(goenId, inheritType, layer.GetParentEntity())
	layer.setCreated()
}

func (p *inheritManager[T, PT]) getAllTables() []string {
	return append(p.getParentManager().getAllTables(), p.tableName)
}

func (p *inheritManager[T, PT]) getParentManager() InheritManagerForRecur {
	return p.parentManager
}

func (p *inheritManager[T, PT]) Get(goenId int) (PT, error) {
	tables := p.getAllTables()
	e := p.getInterface(new(T))
	//query := fmt.Sprintf("select * from %s  where goen_id=? and goen_in_all_instance = true", tables[0])
	for _, table := range tables {
		query := fmt.Sprintf("select * from %s  where goen_id=?", table)
		err := Db.Get(e, query, goenId)
		var nilPT PT
		if err != nil {
			if err == sql.ErrNoRows {
				return nilPT, nil
			}
			return nilPT, err
		}
	}

	//e.setExistent()
	p.recurAfterFind(e)
	p.recurAddInQueue(e)

	return p.getPT(e), nil
}

func (p *inheritManager[T, PT]) GetFromAllInstanceBy(field string, value any) (PT, error) {
	e := p.getInterface(new(T))
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true %s", p.getTablesQuery(), field, p.getJoinQuery())
	err := Db.Get(e, query, value)
	var nilPT PT
	if err != nil {
		if err == sql.ErrNoRows {
			return nilPT, nil
		}
		return nilPT, err
	}
	//e.setExistent()
	p.recurAfterFind(e)
	p.recurAddInQueue(e)
	return p.getPT(e), nil
}

func (p *inheritManager[T, PT]) getTablesQuery() string {
	tables := p.getAllTables()
	tablesQuery := tables[0]
	for _, table := range tables[1:] {
		tablesQuery += fmt.Sprintf(", %s", table)
	}
	return tablesQuery
}

// form: table1.goen_id = table2.goen_id and table2.goen_id = table3.goen_id
func (p *inheritManager[T, PT]) getJoinQuery() string {
	tables := p.getAllTables()
	joinQuery := fmt.Sprintf("and %s.goen_id = ", tables[0])
	for _, table := range tables[1 : len(tables)-1] {
		joinQuery += fmt.Sprintf("%s.goen_id and %s.goen_id = ", table, table)
	}
	joinQuery += fmt.Sprintf("%s.goen_id", tables[len(tables)-1])
	return joinQuery
}

func (p *inheritManager[T, PT]) FindFromAllInstanceBy(field string, value any) ([]PT, error) {
	var entityArr []*T
	var interfaceArr []PT
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true %s", p.getTablesQuery(), field, p.getJoinQuery())
	err := Db.Get(entityArr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		ei := p.getInterface(e)
		//e.setExistent()
		p.recurAfterFind(ei)
		p.recurAddInQueue(ei)
		interfaceArr = append(interfaceArr, p.getPT(ei))
	}
	return interfaceArr, nil
}

func (p *inheritManager[T, PT]) FindFromMultiAssTable(assTableName string, ownerId int) ([]PT, error) {
	var entityArr []*T
	var interfaceArr []PT
	tables := p.getAllTables()
	query := fmt.Sprintf("select tmp.* from %s as ass, %s where ass.owner_goen_id = ? and ass.possession_goen_id = %s.goen_id %s ",
		assTableName, p.getTablesQuery(), tables[0], p.getJoinQuery())
	if err := Db.Select(&entityArr, query, ownerId); err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		ei := p.getInterface(e)
		//e.setExistent()
		p.recurAfterFind(ei)
		p.addInQueue(ei)
		interfaceArr = append(interfaceArr, p.getPT(ei))
	}
	return interfaceArr, nil
}
