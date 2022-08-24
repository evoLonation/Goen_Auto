package entity

import (
	"Auto/entityManager"
	"time"
)

var cardPaymentManager entityManager.ManagerForEntity[CardPayment]
var CardPaymentManager entityManager.InheritManagerForOther[CardPayment]

type CardPayment interface {
	Payment
	GetCardAccountNumber() string
	GetExpiryDate() time.Time
	SetCardAccountNumber(cardAccountNumber string)
	SetExpiryDate(expiryDate time.Time)
}

type CardPaymentEntity struct {
	PaymentEntity
	entityManager.FieldChange

	CardAccountNumber string    `db:"card_account_number"`
	ExpiryDate        time.Time `db:"expiry_date"`
}

func (p *CardPaymentEntity) GetParentEntity() entityManager.EntityForInheritManager {
	return &p.PaymentEntity
}

func (p *CardPaymentEntity) GetCardAccountNumber() string {
	return p.CardAccountNumber
}
func (p *CardPaymentEntity) GetExpiryDate() time.Time {
	return p.ExpiryDate
}
func (p *CardPaymentEntity) SetCardAccountNumber(cardAccountNumber string) {
	p.CardAccountNumber = cardAccountNumber
	p.AddBasicFieldChange("card_account_number")
}
func (p *CardPaymentEntity) SetExpiryDate(expiryDate time.Time) {
	p.ExpiryDate = expiryDate
	p.AddBasicFieldChange("expiry_date")
}
