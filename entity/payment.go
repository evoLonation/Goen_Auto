package entity

import(
"Auto/entityManager"
)
var paymentManager entityManager.ManagerForEntity[Payment]
var PaymentManager entityManager.ManagerForOther[Payment]


type Payment interface{
	GetAmountTendered () float64 
	GetBelongedSale () Sale 
	SetAmountTendered (amountTendered float64) 
	SetBelongedSale (sale Sale) 
}

type PaymentEntity struct{
	entityManager.BasicEntity
	
	AmountTendered float64 `db:"amount_tendered"`
	BelongedSaleGoenId *int `db:"belonged_sale_goen_id"`
}
func (p *PaymentEntity) GetAmountTendered () float64  {
	return p.AmountTendered 
}
func (p *PaymentEntity) GetBelongedSale () Sale  {
	if p.BelongedSaleGoenId == nil {
		return nil
	} else {
		ret, _ := saleManager.Get(*p.BelongedSaleGoenId)
		return ret
	}
}
func (p *PaymentEntity) SetAmountTendered (amountTendered float64)  {
	p.AmountTendered = amountTendered 
	p.AddBasicFieldChange("amount_tendered")
}
func (p *PaymentEntity) SetBelongedSale (sale Sale)  {
	id := saleManager.GetGoenId(sale)
	p.BelongedSaleGoenId = &id
	p.AddAssFieldChange("belonged_sale_goen_id")
}
