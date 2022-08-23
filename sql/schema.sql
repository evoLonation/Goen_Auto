create table card_payment
(
	goen_id                  int primary key,
	card_account_number varchar(255) not null default(''),
	expiry_date datetime not null default('0001-01-01 00:00:00')		
);

create table cash_desk
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	name varchar(255) not null default(''),
	is_opened boolean not null default(false),
	belonged_store_goen_id int		
);

create table cash_desk_contained_sales
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table cash_payment
(
	goen_id                  int primary key,
	balance double not null default(0)		
);

create table cashier
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	name varchar(255) not null default(''),
	worked_store_goen_id int		
);

create table item
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	barcode int not null default(0),
	name varchar(255) not null default(''),
	price double not null default(0),
	stock_number int not null default(0),
	order_price double not null default(0),
	belonged_catalog_goen_id int		
);

create table order_entry
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	quantity int not null default(0),
	sub_amount double not null default(0),
	item_goen_id int		
);

create table order_product
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	time datetime not null default('0001-01-01 00:00:00'),
	amount double not null default(0),
	order_status int not null default (0), #枚举类型 
	supplier_goen_id int		
);

create table order_product_contained_entries
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table payment
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	goen_inherit_type    int   not null default (0),
	amount_tendered double not null default(0),
	belonged_sale_goen_id int		
);

create table product_catalog
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	name varchar(255) not null default('')		
);

create table product_catalog_contained_items
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table sale
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	time datetime not null default('0001-01-01 00:00:00'),
	is_complete boolean not null default(false),
	amount double not null default(0),
	is_readyto_pay boolean not null default(false),
	belongedstore_goen_id int, 
	belonged_cash_desk_goen_id int, 
	assoicated_payment_goen_id int		
);

create table sale_contained_sales_line
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table sales_line_item
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	quantity int not null default(0),
	subamount double not null default(0),
	belonged_sale_goen_id int, 
	belonged_item_goen_id int		
);

create table store
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	name varchar(255) not null default(''),
	address varchar(255) not null default(''),
	is_opened boolean not null default(false)		
);

create table store_association_cashdeskes
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table store_productcatalogs
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table store_items
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table store_cashiers
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table store_sales
(
	owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);

create table supplier
(
	goen_id                  int primary key,
	goen_in_all_instance     bool not null default (false),
	id int not null default(0),
	name varchar(255) not null default('')		
);

alter table card_payment
add constraint foreign key (goen_id) references payment (goen_id) on delete cascade;

alter table cash_desk
add constraint foreign key (belonged_store_goen_id) references store(goen_id) on delete set null;

alter table cash_desk_contained_sales
	add constraint foreign key (owner_goen_id) references cash_desk (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references sale (goen_id) on delete cascade;

alter table cash_payment
add constraint foreign key (goen_id) references payment (goen_id) on delete cascade;

alter table cashier
add constraint foreign key (worked_store_goen_id) references store(goen_id) on delete set null;

alter table item
add constraint foreign key (belonged_catalog_goen_id) references product_catalog(goen_id) on delete set null;

alter table order_entry
add constraint foreign key (item_goen_id) references item(goen_id) on delete set null;

alter table order_product
add constraint foreign key (supplier_goen_id) references supplier(goen_id) on delete set null;

alter table order_product_contained_entries
	add constraint foreign key (owner_goen_id) references order_product (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references order_entry (goen_id) on delete cascade;

alter table payment
add constraint foreign key (belonged_sale_goen_id) references sale(goen_id) on delete set null;


alter table product_catalog_contained_items
	add constraint foreign key (owner_goen_id) references product_catalog (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references item (goen_id) on delete cascade;

alter table sale
add constraint foreign key (belongedstore_goen_id) references store(goen_id) on delete set null,
add constraint foreign key (belonged_cash_desk_goen_id) references cash_desk(goen_id) on delete set null,
add constraint foreign key (assoicated_payment_goen_id) references payment(goen_id) on delete set null;

alter table sale_contained_sales_line
	add constraint foreign key (owner_goen_id) references sale (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references sales_line_item (goen_id) on delete cascade;

alter table sales_line_item
add constraint foreign key (belonged_sale_goen_id) references sale(goen_id) on delete set null,
add constraint foreign key (belonged_item_goen_id) references item(goen_id) on delete set null;


alter table store_association_cashdeskes
	add constraint foreign key (owner_goen_id) references store (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references cash_desk (goen_id) on delete cascade;

alter table store_productcatalogs
	add constraint foreign key (owner_goen_id) references store (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references product_catalog (goen_id) on delete cascade;

alter table store_items
	add constraint foreign key (owner_goen_id) references store (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references item (goen_id) on delete cascade;

alter table store_cashiers
	add constraint foreign key (owner_goen_id) references store (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references cashier (goen_id) on delete cascade;

alter table store_sales
	add constraint foreign key (owner_goen_id) references store (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references sale (goen_id) on delete cascade;


