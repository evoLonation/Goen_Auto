package entity

import(
"Auto/entityManager"
)
var salesLineItemManager entityManager.ManagerForEntity[SalesLineItem]
var SalesLineItemManager entityManager.ManagerForOther[SalesLineItem]


type SalesLineItem interface{
	GetQuantity () int 
	GetSubamount () float64 
	GetBelongedSale () Sale 
	GetBelongedItem () Item 
	SetQuantity (quantity int) 
	SetSubamount (subamount float64) 
	SetBelongedSale (sale Sale) 
	SetBelongedItem (item Item) 
}

type SalesLineItemEntity struct{
	entityManager.Entity
	
	Quantity int `db:"quantity"`
	Subamount float64 `db:"subamount"`
	BelongedSaleGoenId *int `db:"belonged_sale_goen_id"`
	BelongedItemGoenId *int `db:"belonged_item_goen_id"`
}
func (p *SalesLineItemEntity) GetQuantity () int  {
	return p.Quantity 
}
func (p *SalesLineItemEntity) GetSubamount () float64  {
	return p.Subamount 
}
func (p *SalesLineItemEntity) GetBelongedSale () Sale  {
	if p.BelongedSaleGoenId == nil {
		return nil
	} else {
		ret, _ := saleManager.Get(*p.BelongedSaleGoenId)
		return ret
	}
}
func (p *SalesLineItemEntity) GetBelongedItem () Item  {
	if p.BelongedItemGoenId == nil {
		return nil
	} else {
		ret, _ := itemManager.Get(*p.BelongedItemGoenId)
		return ret
	}
}
func (p *SalesLineItemEntity) SetQuantity (quantity int)  {
	p.Quantity = quantity 
	p.AddBasicFieldChange("quantity")
}
func (p *SalesLineItemEntity) SetSubamount (subamount float64)  {
	p.Subamount = subamount 
	p.AddBasicFieldChange("subamount")
}
func (p *SalesLineItemEntity) SetBelongedSale (sale Sale)  {
	id := saleManager.GetGoenId(sale)
	p.BelongedSaleGoenId = &id
	p.AddAssFieldChange("belonged_sale_goen_id")
}
func (p *SalesLineItemEntity) SetBelongedItem (item Item)  {
	id := itemManager.GetGoenId(item)
	p.BelongedItemGoenId = &id
	p.AddAssFieldChange("belonged_item_goen_id")
}
