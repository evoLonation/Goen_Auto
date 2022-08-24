package entity

import (
	"Auto/entityManager"
)

var productCatalogManager entityManager.ManagerForEntity[ProductCatalog]
var ProductCatalogManager entityManager.ManagerForOther[ProductCatalog]

type ProductCatalog interface {
	GetId() int
	GetName() string
	GetContainedItems() []Item
	SetId(id int)
	SetName(name string)
	AddContainedItems(item Item)
}

type ProductCatalogEntity struct {
	entityManager.Entity

	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (p *ProductCatalogEntity) GetId() int {
	return p.Id
}
func (p *ProductCatalogEntity) GetName() string {
	return p.Name
}
func (p *ProductCatalogEntity) GetContainedItems() []Item {
	ret, _ := itemManager.FindFromMultiAssTable("product_catalog_contained_items", p.GoenId)
	return ret
}
func (p *ProductCatalogEntity) SetId(id int) {
	p.Id = id
	p.AddBasicFieldChange("id")
}
func (p *ProductCatalogEntity) SetName(name string) {
	p.Name = name
	p.AddBasicFieldChange("name")
}
func (p *ProductCatalogEntity) AddContainedItems(item Item) {
	p.AddMultiAssChange(entityManager.Include, "product_catalog_contained_items", itemManager.GetGoenId(item))
}
