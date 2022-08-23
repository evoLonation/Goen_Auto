package entity

import(
"Auto/entityManager"
)
var storeManager entityManager.ManagerForEntity[Store]
var StoreManager entityManager.ManagerForOther[Store]


type Store interface{
	GetId () int 
	GetName () string 
	GetAddress () string 
	GetIsOpened () bool 
	GetAssociationCashdeskes () []CashDesk 
	GetProductcatalogs () []ProductCatalog 
	GetItems () []Item 
	GetCashiers () []Cashier 
	GetSales () []Sale 
	SetId (id int) 
	SetName (name string) 
	SetAddress (address string) 
	SetIsOpened (isOpened bool) 
	AddAssociationCashdeskes (cashDesk CashDesk) 
	AddProductcatalogs (productCatalog ProductCatalog) 
	AddItems (item Item) 
	AddCashiers (cashier Cashier) 
	AddSales (sale Sale) 
}

type StoreEntity struct{
	entityManager.Entity
	
	Id int `db:"id"`
	Name string `db:"name"`
	Address string `db:"address"`
	IsOpened bool `db:"is_opened"`
}
func (p *StoreEntity) GetId () int  {
	return p.Id 
}
func (p *StoreEntity) GetName () string  {
	return p.Name 
}
func (p *StoreEntity) GetAddress () string  {
	return p.Address 
}
func (p *StoreEntity) GetIsOpened () bool  {
	return p.IsOpened 
}
func (p *StoreEntity) GetAssociationCashdeskes () []CashDesk  {
	ret, _ := cashDeskManager.FindFromMultiAssTable("store_association_cashdeskes", p.GoenId)
	return ret 
}
func (p *StoreEntity) GetProductcatalogs () []ProductCatalog  {
	ret, _ := productCatalogManager.FindFromMultiAssTable("store_productcatalogs", p.GoenId)
	return ret 
}
func (p *StoreEntity) GetItems () []Item  {
	ret, _ := itemManager.FindFromMultiAssTable("store_items", p.GoenId)
	return ret 
}
func (p *StoreEntity) GetCashiers () []Cashier  {
	ret, _ := cashierManager.FindFromMultiAssTable("store_cashiers", p.GoenId)
	return ret 
}
func (p *StoreEntity) GetSales () []Sale  {
	ret, _ := saleManager.FindFromMultiAssTable("store_sales", p.GoenId)
	return ret 
}
func (p *StoreEntity) SetId (id int)  {
	p.Id = id 
	p.AddBasicFieldChange("id")
}
func (p *StoreEntity) SetName (name string)  {
	p.Name = name 
	p.AddBasicFieldChange("name")
}
func (p *StoreEntity) SetAddress (address string)  {
	p.Address = address 
	p.AddBasicFieldChange("address")
}
func (p *StoreEntity) SetIsOpened (isOpened bool)  {
	p.IsOpened = isOpened 
	p.AddBasicFieldChange("is_opened")
}
func (p *StoreEntity) AddAssociationCashdeskes (cashDesk CashDesk)  {
	p.AddMultiAssChange(entityManager.Include, "store_association_cashdeskes", cashDeskManager.GetGoenId(cashDesk))
}
func (p *StoreEntity) AddProductcatalogs (productCatalog ProductCatalog)  {
	p.AddMultiAssChange(entityManager.Include, "store_productcatalogs", productCatalogManager.GetGoenId(productCatalog))
}
func (p *StoreEntity) AddItems (item Item)  {
	p.AddMultiAssChange(entityManager.Include, "store_items", itemManager.GetGoenId(item))
}
func (p *StoreEntity) AddCashiers (cashier Cashier)  {
	p.AddMultiAssChange(entityManager.Include, "store_cashiers", cashierManager.GetGoenId(cashier))
}
func (p *StoreEntity) AddSales (sale Sale)  {
	p.AddMultiAssChange(entityManager.Include, "store_sales", saleManager.GetGoenId(sale))
}
