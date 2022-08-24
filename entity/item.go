package entity

import (
	"Auto/entityManager"
)

var itemManager entityManager.ManagerForEntity[Item]
var ItemManager entityManager.ManagerForOther[Item]

type Item interface {
	GetBarcode() int
	GetName() string
	GetPrice() float64
	GetStockNumber() int
	GetOrderPrice() float64
	GetBelongedCatalog() ProductCatalog
	SetBarcode(barcode int)
	SetName(name string)
	SetPrice(price float64)
	SetStockNumber(stockNumber int)
	SetOrderPrice(orderPrice float64)
	SetBelongedCatalog(productCatalog ProductCatalog)
}

type ItemEntity struct {
	entityManager.Entity

	Barcode               int     `db:"barcode"`
	Name                  string  `db:"name"`
	Price                 float64 `db:"price"`
	StockNumber           int     `db:"stock_number"`
	OrderPrice            float64 `db:"order_price"`
	BelongedCatalogGoenId *int    `db:"belonged_catalog_goen_id"`
}

func (p *ItemEntity) GetBarcode() int {
	return p.Barcode
}
func (p *ItemEntity) GetName() string {
	return p.Name
}
func (p *ItemEntity) GetPrice() float64 {
	return p.Price
}
func (p *ItemEntity) GetStockNumber() int {
	return p.StockNumber
}
func (p *ItemEntity) GetOrderPrice() float64 {
	return p.OrderPrice
}
func (p *ItemEntity) GetBelongedCatalog() ProductCatalog {
	if p.BelongedCatalogGoenId == nil {
		return nil
	} else {
		ret, _ := productCatalogManager.Get(*p.BelongedCatalogGoenId)
		return ret
	}
}
func (p *ItemEntity) SetBarcode(barcode int) {
	p.Barcode = barcode
	p.AddBasicFieldChange("barcode")
}
func (p *ItemEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *ItemEntity) SetPrice(price float64) {
	p.Price = price
	p.AddBasicFieldChange("price")
}
func (p *ItemEntity) SetStockNumber(stockNumber int) {
	p.StockNumber = stockNumber
	p.AddBasicFieldChange("stock_number")
}
func (p *ItemEntity) SetOrderPrice(orderPrice float64) {
	p.OrderPrice = orderPrice
	p.AddBasicFieldChange("order_price")
}
func (p *ItemEntity) SetBelongedCatalog(productCatalog ProductCatalog) {
	id := productCatalogManager.GetGoenId(productCatalog)
	p.BelongedCatalogGoenId = &id
	p.AddAssFieldChange("belonged_catalog_goen_id")
}
