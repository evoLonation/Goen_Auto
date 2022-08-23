package entityManager

import (
	"database/sql"
	"errors"
	"fmt"
)

type EntityForManager interface {
	afterNew(int)
	afterFind()
	afterBasicSave()
	afterAssSave()
	getEntityStatus() entityStatus
	setGoenInAllInstance(bool)
	getBasicFieldChange() []string
	getAssFieldChange() []string
	getMultiAssChange() []multiAssInfo
	GetGoenId() int
}

//type managerTypeParam[T any] interface {
//	*T
//	EntityForManager
//}

type manager[T any, PT any] struct {
	tableName string
	maxGoenId int
}

func NewManager[T any, PT any](tableName string) (*manager[T, PT], error) {
	_, ok := (any(new(T))).(PT)
	if !ok {
		return nil, errors.New("the type value T does not implement PT ")
	}
	_, ok = (any(new(T))).(EntityForManager)
	if !ok {
		return nil, errors.New("the type value T does not implement EntityForManager ")
	}
	manager := &manager[T, PT]{}
	manager.tableName = tableName
	query := fmt.Sprintf("select goen_id from %s order by goen_id DESC limit 1", manager.tableName)
	err := Db.Get(&manager.maxGoenId, query)
	if err != nil && err != sql.ErrNoRows {
		return manager, err
	}
	return manager, nil
}

func (p *manager[T, PT]) getInterface(e *T) EntityForManager {
	return (any(e)).(EntityForManager)
}

func (p *manager[T, PT]) getPT(ei EntityForManager) PT {
	return (any(ei)).(PT)
}

func (p *manager[T, PT]) GetGoenId(e PT) int {
	return (any(e)).(EntityForManager).GetGoenId()
}

func (p *manager[T, PT]) generateGoenId() int {
	p.maxGoenId = p.maxGoenId + 1
	return p.maxGoenId
}

func (p *manager[T, PT]) New() PT {
	e := p.getInterface(new(T))
	e.afterNew(p.generateGoenId())
	p.addInQueue(e)
	return p.getPT(e)
}

// Get if no rows, return nil, nil
func (p *manager[T, PT]) Get(goenId int) (PT, error) {
	e := p.getInterface(new(T))
	//query := fmt.Sprintf("select * from %s where goen_id=? and goen_in_all_instance = true", p.tableName)
	query := fmt.Sprintf("select * from %s where goen_id=?", p.tableName)
	err := Db.Get(e, query, goenId)
	if err != nil {
		var nilPT PT
		if err == sql.ErrNoRows {
			return nilPT, nil
		}
		return nilPT, err
	}
	e.afterFind()
	p.addInQueue(e)
	return p.getPT(e), nil
}

func (p *manager[T, PT]) GetFromAllInstanceBy(member string, value any) (PT, error) {
	e := p.getInterface(new(T))
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Get(e, query, value)
	if err != nil {
		var nilPT PT
		if err == sql.ErrNoRows {
			return nilPT, nil
		}
		return nilPT, err
	}
	e.afterFind()
	p.addInQueue(e)
	return p.getPT(e), nil
}

func (p *manager[T, PT]) FindFromAllInstanceBy(member string, value any) ([]PT, error) {
	var entityArr []*T
	var interfaceArr []PT
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Select(&entityArr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		ei := p.getInterface(e)
		ei.afterFind()
		p.addInQueue(ei)
		interfaceArr = append(interfaceArr, p.getPT(ei))
	}
	return interfaceArr, nil
}

func (p *manager[T, PT]) AddInAllInstance(e PT) {
	(any(e)).(EntityForManager).setGoenInAllInstance(true)
}

func (p *manager[T, PT]) RemoveFromAllInstance(e PT) {
	(any(e)).(EntityForManager).setGoenInAllInstance(false)
}

func (p *manager[T, PT]) FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error) {
	var entityArr []*T
	var interfaceArr []PT
	query := fmt.Sprintf("select tmp.* from %s as ass, %s as tmp where ass.owner_goen_id = ? and ass.possession_goen_id = tmp.goen_id ",
		tableName, p.tableName)
	if err := Db.Select(&entityArr, query, ownerId); err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		ei := p.getInterface(e)
		ei.afterFind()
		p.addInQueue(ei)
		interfaceArr = append(interfaceArr, p.getPT(ei))
	}
	return interfaceArr, nil
}

func (p *manager[T, PT]) addInQueue(e EntityForManager) {
	Saver.addInBasicSaveQueue(func() error {
		return p.saveBasic(e)
	})
	Saver.addInAssSaveQueue(func() error {
		return p.saveAss(e)
	})
}

// the length of changedField must > 0
func (p *manager[T, PT]) getUpdateQuery(changedField []string) string {
	query := fmt.Sprintf("update %s set ", p.tableName)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf("%s= :%s,", field, field)
	}
	field := changedField[len(changedField)-1]
	query += fmt.Sprintf("%s = :%s", field, field)
	query += fmt.Sprintf(" where goen_id = :goen_id")
	print(query)
	return query
}

// the length of changedField must > 0
func (p *manager[T, PT]) getInsertQuery(changedField []string) string {
	lastField := changedField[len(changedField)-1]
	query := fmt.Sprintf("insert into %s(goen_id, ", p.tableName)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf("%s, ", field)
	}
	query += fmt.Sprintf("%s) values(:goen_id, ", lastField)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf(":%s ,", field)
	}
	query += fmt.Sprintf(":%s )", lastField)
	print(query)
	return query
}

func (p *manager[T, PT]) getMultiAssInsertQuery(tableName string) string {
	query := fmt.Sprintf("insert into %s (owner_goen_id, possession_goen_id) values (?, ?)", tableName)
	return query
}
func (p *manager[T, PT]) getMultiAssDeleteQuery(tableName string) string {
	query := fmt.Sprintf("delete from %s where owner_goen_id=? and possession_goen_id=?", tableName)
	return query
}

func (p *manager[T, PT]) saveBasic(e EntityForManager) error {
	if len(e.getBasicFieldChange()) != 0 {
		if e.getEntityStatus() == Created {
			if _, err := Db.NamedExec(p.getInsertQuery(e.getBasicFieldChange()), e); err != nil {
				return err
			}
		} else {
			if _, err := Db.NamedExec(p.getUpdateQuery(e.getBasicFieldChange()), e); err != nil {
				return err
			}
		}
	}
	e.afterBasicSave()
	return nil
}

// saveAss e 's entityStatus must be Existence
func (p *manager[T, PT]) saveAss(e EntityForManager) error {

	if e.getEntityStatus() == Created {
		return errors.New("entityStatus must be Existence")
	}
	if len(e.getAssFieldChange()) != 0 {
		if _, err := Db.NamedExec(p.getUpdateQuery(e.getAssFieldChange()), e); err != nil {
			return err
		}
	}
	for _, info := range e.getMultiAssChange() {
		var query string
		if info.typ == Include {
			query = p.getMultiAssInsertQuery(info.tableName)
		} else {
			query = p.getMultiAssDeleteQuery(info.tableName)
		}
		if _, err := Db.Exec(query, e.GetGoenId(), info.targetId); err != nil {
			return err
		}
	}
	e.afterAssSave()
	return nil
}
