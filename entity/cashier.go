package entity

import(
"Auto/entityManager"
)
var cashierManager entityManager.ManagerForEntity[Cashier]
var CashierManager entityManager.ManagerForOther[Cashier]


type Cashier interface{
	GetId () int 
	GetName () string 
	GetWorkedStore () Store 
	SetId (id int) 
	SetName (name string) 
	SetWorkedStore (store Store) 
}

type CashierEntity struct{
	entityManager.Entity
	
	Id int `db:"id"`
	Name string `db:"name"`
	WorkedStoreGoenId *int `db:"worked_store_goen_id"`
}
func (p *CashierEntity) GetId () int  {
	return p.Id 
}
func (p *CashierEntity) GetName () string  {
	return p.Name 
}
func (p *CashierEntity) GetWorkedStore () Store  {
	if p.WorkedStoreGoenId == nil {
		return nil
	} else {
		ret, _ := storeManager.Get(*p.WorkedStoreGoenId)
		return ret
	}
}
func (p *CashierEntity) SetId (id int)  {
	p.Id = id 
	p.AddBasicFieldChange("id")
}
func (p *CashierEntity) SetName (name string)  {
	p.Name = name 
	p.AddBasicFieldChange("name")
}
func (p *CashierEntity) SetWorkedStore (store Store)  {
	id := storeManager.GetGoenId(store)
	p.WorkedStoreGoenId = &id
	p.AddAssFieldChange("worked_store_goen_id")
}
