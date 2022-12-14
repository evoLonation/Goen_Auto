package service

import (
	"Auto/entity"
	"Auto/entityManager"
	"time"
)

var CurrentSaleLine entity.SalesLineItem
var CurrentSale entity.Sale
var CurrentPaymentMethod entity.PaymentMethod

func MakeNewSale() (result bool, retErr error) {
	defer func() {
		if err := entityManager.Saver.Save(); err != nil {
			retErr = NewErrPostCondition(err)
			return
		}
	}()

	//precondition
	if !((CurrentCashDesk == nil) == false &&
		CurrentCashDesk.GetIsOpened() == true &&
		((CurrentSale == nil) == true ||
			((CurrentSale == nil) == false &&
				CurrentSale.GetIsComplete() == true))) {
		retErr = ErrPreConditionUnsatisfied
		return
	}

	// post condition
	s := entity.SaleManager.New()
	s.SetBelongedCashDesk(CurrentCashDesk)
	CurrentCashDesk.AddContainedSales(s)
	s.SetIsComplete(false)
	s.SetIsReadytoPay(false)
	entity.SaleManager.AddInAllInstance(s)
	CurrentSale = s
	result = true
	return
}

func EnterItem(barcode int, quantity int) (bool, error) {
	// definition
	item, err := entity.ItemManager.GetFromAllInstanceBy("barcode", barcode)
	if err != nil {
		return false, nil
	}

	// precondition
	if !((CurrentSale == nil) == false &&
		CurrentSale.GetIsComplete() == false &&
		(item == nil) == false &&
		item.GetStockNumber() > 0) {
		return false, ErrPreConditionUnsatisfied
	}

	// post condition
	sli := entity.SalesLineItemManager.New()
	CurrentSaleLine = sli
	sli.SetBelongedSale(CurrentSale)
	CurrentSale.AddContainedSalesLine(sli)
	sli.SetQuantity(quantity)
	sli.SetBelongedItem(item)
	item.SetStockNumber(item.GetStockNumber() - quantity)
	sli.SetSubamount(item.GetPrice() * float64(quantity))
	entity.SalesLineItemManager.AddInAllInstance(sli)

	if err := entityManager.Saver.Save(); err != nil {
		return false, NewErrPostCondition(err)
	}
	return true, nil
}

func EndSale() (float64, error) {

	// definition
	var sls []entity.SalesLineItem = CurrentSale.GetContainedSalesLine()
	var sub []float64
	for _, sl := range sls {
		sub = append(sub, sl.GetSubamount())
	}

	// precondition
	if !((CurrentSale == nil) == false &&
		CurrentSale.GetIsComplete() == false &&
		CurrentSale.GetIsReadytoPay() == false) {
		return 0, ErrPreConditionUnsatisfied
	}

	// post condition
	CurrentSale.SetAmount(Sum(sub))
	CurrentSale.SetIsReadytoPay(true)

	if err := entityManager.Saver.Save(); err != nil {
		return 0, NewErrPostCondition(err)
	}
	return CurrentSale.GetAmount(), nil
}
func Collect[FT any, ET any](e ET, field string) []FT {

	return []FT{}
}

type Addable interface {
	int | float32 | float64 | string
}

func Sum[T Addable](arr []T) T {
	var a T
	for i := 0; i <= len(arr); i++ {
		a = arr[i] + a
	}
	return a
}

func MakeCashPayment(amount float64) (bool, error) {
	// precondition
	if !((CurrentSale == nil) == false &&
		CurrentSale.GetIsComplete() == false &&
		CurrentSale.GetIsReadytoPay() == true &&
		amount >= CurrentSale.GetAmount()) {
		return false, ErrPreConditionUnsatisfied
	}

	// post condition
	var cp entity.CashPayment
	cp = entity.CashPaymentManager.New()
	cp.SetAmountTendered(amount)
	cp.SetBelongedSale(CurrentSale)
	CurrentSale.SetAssoicatedPayment(cp)
	CurrentSale.SetBelongedstore(CurrentStore)
	CurrentStore.AddSales(CurrentSale)
	CurrentSale.SetIsComplete(true)
	CurrentSale.SetTime(time.Now())
	cp.SetBalance(amount - CurrentSale.GetAmount())
	entity.CashPaymentManager.AddInAllInstance(cp)

	if err := entityManager.Saver.Save(); err != nil {
		return false, NewErrPostCondition(err)
	}
	return true, nil
}

func MakeCardPayment(cardAccountNumber string, expiryDate time.Time, fee float64) (bool, error) {
	// precondition
	if !((CurrentSale == nil) == false &&
		CurrentSale.GetIsComplete() == false &&
		CurrentSale.GetIsReadytoPay() == true) {
		return false, ErrPreConditionUnsatisfied
	}

	// post condition
	var cdp entity.CardPayment
	cdp = entity.CardPaymentManager.New()
	cdp.SetAmountTendered(fee)
	cdp.SetBelongedSale(CurrentSale)
	CurrentSale.SetAssoicatedPayment(cdp)
	cdp.SetCardAccountNumber(cardAccountNumber)
	cdp.SetExpiryDate(expiryDate)
	entity.CardPaymentManager.AddInAllInstance(cdp)
	CurrentSale.SetBelongedstore(CurrentStore)
	CurrentStore.AddSales(CurrentSale)
	CurrentSale.SetIsComplete(true)
	CurrentSale.SetTime(time.Now())

	if err := entityManager.Saver.Save(); err != nil {
		return false, NewErrPostCondition(err)
	}
	return true, nil
}
