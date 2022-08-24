package entity

import (
	"Auto/entityManager"
)

var cashDeskManager entityManager.ManagerForEntity[CashDesk]
var CashDeskManager entityManager.ManagerForOther[CashDesk]

type CashDesk interface {
	GetId() int
	GetName() string
	GetIsOpened() bool
	GetBelongedStore() Store
	GetContainedSales() []Sale
	SetId(id int)
	SetName(name string)
	SetIsOpened(isOpened bool)
	SetBelongedStore(store Store)
	AddContainedSales(sale Sale)
}

type CashDeskEntity struct {
	entityManager.Entity

	Id                  int    `db:"id"`
	Name                string `db:"name"`
	IsOpened            bool   `db:"is_opened"`
	BelongedStoreGoenId *int   `db:"belonged_store_goen_id"`
}

func (p *CashDeskEntity) GetId() int {
	return p.Id
}
func (p *CashDeskEntity) GetName() string {
	return p.Name
}
func (p *CashDeskEntity) GetIsOpened() bool {
	return p.IsOpened
}
func (p *CashDeskEntity) GetBelongedStore() Store {
	if p.BelongedStoreGoenId == nil {
		return nil
	} else {
		ret, _ := storeManager.Get(*p.BelongedStoreGoenId)
		return ret
	}
}
func (p *CashDeskEntity) GetContainedSales() []Sale {
	ret, _ := saleManager.FindFromMultiAssTable("cash_desk_contained_sales", p.GoenId)
	return ret
}
func (p *CashDeskEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *CashDeskEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *CashDeskEntity) SetIsOpened(isOpened bool) {
	p.IsOpened = isOpened
	p.AddBasicFieldChange("is_opened")
}
func (p *CashDeskEntity) SetBelongedStore(store Store) {
	id := storeManager.GetGoenId(store)
	p.BelongedStoreGoenId = &id
	p.AddAssFieldChange("belonged_store_goen_id")
}
func (p *CashDeskEntity) AddContainedSales(sale Sale) {
	p.AddMultiAssChange(entityManager.Include, "cash_desk_contained_sales", saleManager.GetGoenId(sale))
}
