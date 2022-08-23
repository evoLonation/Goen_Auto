package entity

import (
	"Auto/entityManager"
	"log"
)

const (
	CardPaymentInheritType entityManager.GoenInheritType = iota + 1
	CashPaymentInheritType
)

func init() {
	tmpPaymentManager, err := entityManager.NewManager[PaymentEntity, Payment]("payment")			
	if err != nil {
		log.Fatal(err)
	}
	paymentManager = tmpPaymentManager
	PaymentManager = tmpPaymentManager
	tmpCardPaymentManager, err := entityManager.NewInheritManager[CardPaymentEntity, CardPayment]("card_payment", tmpPaymentManager, CardPaymentInheritType)			
	if err != nil {
		log.Fatal(err)
	}
	cardPaymentManager = tmpCardPaymentManager
	CardPaymentManager = tmpCardPaymentManager
	tmpCashDeskManager, err := entityManager.NewManager[CashDeskEntity, CashDesk]("cash_desk")			
	if err != nil {
		log.Fatal(err)
	}
	cashDeskManager = tmpCashDeskManager
	CashDeskManager = tmpCashDeskManager
	tmpCashPaymentManager, err := entityManager.NewInheritManager[CashPaymentEntity, CashPayment]("cash_payment", tmpPaymentManager, CashPaymentInheritType)			
	if err != nil {
		log.Fatal(err)
	}
	cashPaymentManager = tmpCashPaymentManager
	CashPaymentManager = tmpCashPaymentManager
	tmpCashierManager, err := entityManager.NewManager[CashierEntity, Cashier]("cashier")			
	if err != nil {
		log.Fatal(err)
	}
	cashierManager = tmpCashierManager
	CashierManager = tmpCashierManager
	tmpItemManager, err := entityManager.NewManager[ItemEntity, Item]("item")			
	if err != nil {
		log.Fatal(err)
	}
	itemManager = tmpItemManager
	ItemManager = tmpItemManager
	tmpOrderEntryManager, err := entityManager.NewManager[OrderEntryEntity, OrderEntry]("order_entry")			
	if err != nil {
		log.Fatal(err)
	}
	orderEntryManager = tmpOrderEntryManager
	OrderEntryManager = tmpOrderEntryManager
	tmpOrderProductManager, err := entityManager.NewManager[OrderProductEntity, OrderProduct]("order_product")			
	if err != nil {
		log.Fatal(err)
	}
	orderProductManager = tmpOrderProductManager
	OrderProductManager = tmpOrderProductManager
	tmpProductCatalogManager, err := entityManager.NewManager[ProductCatalogEntity, ProductCatalog]("product_catalog")			
	if err != nil {
		log.Fatal(err)
	}
	productCatalogManager = tmpProductCatalogManager
	ProductCatalogManager = tmpProductCatalogManager
	tmpSaleManager, err := entityManager.NewManager[SaleEntity, Sale]("sale")			
	if err != nil {
		log.Fatal(err)
	}
	saleManager = tmpSaleManager
	SaleManager = tmpSaleManager
	tmpSalesLineItemManager, err := entityManager.NewManager[SalesLineItemEntity, SalesLineItem]("sales_line_item")			
	if err != nil {
		log.Fatal(err)
	}
	salesLineItemManager = tmpSalesLineItemManager
	SalesLineItemManager = tmpSalesLineItemManager
	tmpStoreManager, err := entityManager.NewManager[StoreEntity, Store]("store")			
	if err != nil {
		log.Fatal(err)
	}
	storeManager = tmpStoreManager
	StoreManager = tmpStoreManager
	tmpSupplierManager, err := entityManager.NewManager[SupplierEntity, Supplier]("supplier")			
	if err != nil {
		log.Fatal(err)
	}
	supplierManager = tmpSupplierManager
	SupplierManager = tmpSupplierManager
	
}
