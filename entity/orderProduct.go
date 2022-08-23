package entity

import(
"Auto/entityManager"
"time"
)
var orderProductManager entityManager.ManagerForEntity[OrderProduct]
var OrderProductManager entityManager.ManagerForOther[OrderProduct]

type OrderStatus int

const (
	OrderStatusNEW OrderStatus = iota
	OrderStatusRECEIVED
	OrderStatusREQUESTED
)

type OrderProduct interface{
	GetId () int 
	GetTime () time.Time 
	GetAmount () float64 
	GetOrderStatus () OrderStatus 
	GetSupplier () Supplier 
	GetContainedEntries () []OrderEntry 
	SetId (id int) 
	SetTime (time time.Time) 
	SetAmount (amount float64) 
	SetOrderStatus (orderStatus OrderStatus) 
	SetSupplier (supplier Supplier) 
	AddContainedEntries (orderEntry OrderEntry) 
}

type OrderProductEntity struct{
	entityManager.Entity
	
	Id int `db:"id"`
	Time time.Time `db:"time"`
	Amount float64 `db:"amount"`
	OrderStatus OrderStatus `db:"order_status"`
	SupplierGoenId *int `db:"supplier_goen_id"`
}
func (p *OrderProductEntity) GetId () int  {
	return p.Id 
}
func (p *OrderProductEntity) GetTime () time.Time  {
	return p.Time 
}
func (p *OrderProductEntity) GetAmount () float64  {
	return p.Amount 
}
func (p *OrderProductEntity) GetOrderStatus () OrderStatus  {
	return p.OrderStatus 
}
func (p *OrderProductEntity) GetSupplier () Supplier  {
	if p.SupplierGoenId == nil {
		return nil
	} else {
		ret, _ := supplierManager.Get(*p.SupplierGoenId)
		return ret
	}
}
func (p *OrderProductEntity) GetContainedEntries () []OrderEntry  {
	ret, _ := orderEntryManager.FindFromMultiAssTable("order_product_contained_entries", p.GoenId)
	return ret 
}
func (p *OrderProductEntity) SetId (id int)  {
	p.Id = id 
	p.AddBasicFieldChange("id")
}
func (p *OrderProductEntity) SetTime (time time.Time)  {
	p.Time = time 
	p.AddBasicFieldChange("time")
}
func (p *OrderProductEntity) SetAmount (amount float64)  {
	p.Amount = amount 
	p.AddBasicFieldChange("amount")
}
func (p *OrderProductEntity) SetOrderStatus (orderStatus OrderStatus)  {
	p.OrderStatus = orderStatus 
	p.AddBasicFieldChange("order_status")
}
func (p *OrderProductEntity) SetSupplier (supplier Supplier)  {
	id := supplierManager.GetGoenId(supplier)
	p.SupplierGoenId = &id
	p.AddAssFieldChange("supplier_goen_id")
}
func (p *OrderProductEntity) AddContainedEntries (orderEntry OrderEntry)  {
	p.AddMultiAssChange(entityManager.Include, "order_product_contained_entries", orderEntryManager.GetGoenId(orderEntry))
}
