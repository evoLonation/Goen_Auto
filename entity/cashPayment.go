package entity

import(
"Auto/entityManager"
)
var cashPaymentManager entityManager.ManagerForEntity[CashPayment]
var CashPaymentManager entityManager.ManagerForOther[CashPayment]


type CashPayment interface{
	Payment
	GetBalance () float64 
	SetBalance (balance float64) 
}

type CashPaymentEntity struct{
	PaymentEntity
	entityManager.FieldChange
	
	Balance float64 `db:"balance"`
}
func (p *CashPaymentEntity) GetParentEntity() entityManager.EntityForInheritManager {
	return &p.PaymentEntity
}

func (p *CashPaymentEntity) GetBalance () float64  {
	return p.Balance 
}
func (p *CashPaymentEntity) SetBalance (balance float64)  {
	p.Balance = balance 
	p.AddBasicFieldChange("balance")
}
