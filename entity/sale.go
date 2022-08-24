package entity

import (
	"Auto/entityManager"
	"time"
)

var saleManager entityManager.ManagerForEntity[Sale]
var SaleManager entityManager.ManagerForOther[Sale]

type Sale interface {
	GetTime() time.Time
	GetIsComplete() bool
	GetAmount() float64
	GetIsReadytoPay() bool
	GetBelongedstore() Store
	GetBelongedCashDesk() CashDesk
	GetAssoicatedPayment() Payment
	GetContainedSalesLine() []SalesLineItem
	SetTime(time time.Time)
	SetIsComplete(isComplete bool)
	SetAmount(amount float64)
	SetIsReadytoPay(isReadytoPay bool)
	SetBelongedstore(store Store)
	SetBelongedCashDesk(cashDesk CashDesk)
	SetAssoicatedPayment(payment Payment)
	AddContainedSalesLine(salesLineItem SalesLineItem)
}

type SaleEntity struct {
	entityManager.Entity

	Time                    time.Time `db:"time"`
	IsComplete              bool      `db:"is_complete"`
	Amount                  float64   `db:"amount"`
	IsReadytoPay            bool      `db:"is_readyto_pay"`
	BelongedstoreGoenId     *int      `db:"belongedstore_goen_id"`
	BelongedCashDeskGoenId  *int      `db:"belonged_cash_desk_goen_id"`
	AssoicatedPaymentGoenId *int      `db:"assoicated_payment_goen_id"`
}

func (p *SaleEntity) GetTime() time.Time {
	return p.Time
}
func (p *SaleEntity) GetIsComplete() bool {
	return p.IsComplete
}
func (p *SaleEntity) GetAmount() float64 {
	return p.Amount
}
func (p *SaleEntity) GetIsReadytoPay() bool {
	return p.IsReadytoPay
}
func (p *SaleEntity) GetBelongedstore() Store {
	if p.BelongedstoreGoenId == nil {
		return nil
	} else {
		ret, _ := storeManager.Get(*p.BelongedstoreGoenId)
		return ret
	}
}
func (p *SaleEntity) GetBelongedCashDesk() CashDesk {
	if p.BelongedCashDeskGoenId == nil {
		return nil
	} else {
		ret, _ := cashDeskManager.Get(*p.BelongedCashDeskGoenId)
		return ret
	}
}
func (p *SaleEntity) GetAssoicatedPayment() Payment {
	if p.AssoicatedPaymentGoenId == nil {
		return nil
	} else {
		ret, _ := paymentManager.Get(*p.AssoicatedPaymentGoenId)
		return ret
	}
}
func (p *SaleEntity) GetContainedSalesLine() []SalesLineItem {
	ret, _ := salesLineItemManager.FindFromMultiAssTable("sale_contained_sales_line", p.GoenId)
	return ret
}
func (p *SaleEntity) SetTime(time time.Time) {
	p.Time = time
	p.AddBasicFieldChange("time")
}
func (p *SaleEntity) SetIsComplete(isComplete bool) {
	p.IsComplete = isComplete
	p.AddBasicFieldChange("is_complete")
}
func (p *SaleEntity) SetAmount(amount float64) {
	p.Amount = amount
	p.AddBasicFieldChange("amount")
}
func (p *SaleEntity) SetIsReadytoPay(isReadytoPay bool) {
	p.IsReadytoPay = isReadytoPay
	p.AddBasicFieldChange("is_readyto_pay")
}
func (p *SaleEntity) SetBelongedstore(store Store) {
	id := storeManager.GetGoenId(store)
	p.BelongedstoreGoenId = &id
	p.AddAssFieldChange("belongedstore_goen_id")
}
func (p *SaleEntity) SetBelongedCashDesk(cashDesk CashDesk) {
	id := cashDeskManager.GetGoenId(cashDesk)
	p.BelongedCashDeskGoenId = &id
	p.AddAssFieldChange("belonged_cash_desk_goen_id")
}
func (p *SaleEntity) SetAssoicatedPayment(payment Payment) {
	id := paymentManager.GetGoenId(payment)
	p.AssoicatedPaymentGoenId = &id
	p.AddAssFieldChange("assoicated_payment_goen_id")
}
func (p *SaleEntity) AddContainedSalesLine(salesLineItem SalesLineItem) {
	p.AddMultiAssChange(entityManager.Include, "sale_contained_sales_line", salesLineItemManager.GetGoenId(salesLineItem))
}
