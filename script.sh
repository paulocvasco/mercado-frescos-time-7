#!/usr/bin/env bash

function create_database {
    mysql --user=${USER} --password=${PASS} --execute="DROP DATABASE IF EXISTS mercado_fresco_db;"
    mysql --user=${USER} --password=${PASS} --execute="CREATE DATABASE mercado_fresco_db;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="SET SQL_MODE = \"NO_AUTO_VALUE_ON_ZERO\";"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="START TRANSACTION;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="SET time_zone = \"-03:00\";"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE logs (id int(11) NOT NULL,  message varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL, time varchar(255) COLLATE utf8mb4_unicode_ci);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE logs  ADD PRIMARY KEY (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE logs  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE buyers (id int(11) NOT NULL,  id_card_number varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL, first_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL, last_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE carriers (  id int(11) NOT NULL,  cid varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  company_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  address varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  telephone varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  locality_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE countries (  Id int(11) NOT NULL,  country_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE employees (  id int(11) NOT NULL,  id_card_number varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  first_name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  last_name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  warehouse_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE inbound_orders (  id int(11) NOT NULL,  order_date datetime(6) NOT NULL,  order_number varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  employe_id int(11) NOT NULL,  product_batch_id int(11) NOT NULL,  warehouse_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE localities (  id int(11) NOT NULL,  locality_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  province_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE order_details (  id int(11) NOT NULL,  clean_liness_status varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  quantity int(11) NOT NULL,  temperature decimal(19,2) NOT NULL,  product_record int(11) NOT NULL,  purchase_order int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE order_status (  id int(11) NOT NULL,  description varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE products (  id int(11) NOT NULL,  description varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  expiration_rate decimal(19,2) DEFAULT NULL,  freezing_rate decimal(19,2) DEFAULT NULL,  height decimal(19,2) DEFAULT NULL,  length decimal(19,2) DEFAULT NULL,  net_weight decimal(19,2) DEFAULT NULL,  product_code varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  recommended_freezing_temperature decimal(19,2) DEFAULT NULL,  width decimal(19,2) DEFAULT NULL,  product_type_id int(11) NOT NULL,  seller_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE products_batches (  id int(11) UNIQUE NOT NULL AUTO_INCREMENT,  batch_number int(11) UNIQUE NOT NULL ,  current_quantity int(11) default NULL,  current_tempertature decimal(19,2) DEFAULT NULL,  due_date datetime(6) DEFAULT NULL,  initial_quantity int(11) DEFAULT NULL,  manufacturing_date datetime(6) NOT NULL,  manufacturing_hour int(11) default NULL,  minimum_temperature decimal(19,2) DEFAULT NULL,  product_id int(11) NOT NULL,  section_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE products_types (  Id int(11) NOT NULL,  description varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE product_records (  id int(11) NOT NULL,  last_update_date datetime(6) NOT NULL,  purchase_price decimal(19,2) NOT NULL,  sale_price decimal(19,2) NOT NULL,  product_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE provinces (  id int(11) NOT NULL,  province_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  id_country_fk int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE purchase_orders (  id int(11) NOT NULL,  order_number varchar(255) UNIQUE COLLATE utf8mb4_unicode_ci DEFAULT NULL,  order_date datetime(6) DEFAULT NULL,  tracking_code varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  buyer_id int(11) NOT NULL,  carrier_id int(11) NOT NULL,  order_status_id int(11) NOT NULL,  warehouse_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE rol (  id int(11) NOT NULL,  description varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  rol_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE sections (  id int(11) UNIQUE NOT NULL auto_increment,  section_number int(11) NOT NULL UNIQUE,  current_capacity int(11) NOT NULL,  current_temperature int(11) NOT NULL,  maximum_capacity int(11) NOT NULL,  minimum_capacity int(11) NOT NULL,  minimum_temperature int(11) NOT NULL,  product_type_id int(11) NOT NULL,  warehouse_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE sellers (  id int(11) NOT NULL,  cid varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  company_name varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  address varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  telephone varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  locality_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE users (  id int(11) NOT NULL,  password varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  username varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE user_rol (  usuario_id int(11) DEFAULT NULL,  rol_id int(11) DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="CREATE TABLE warehouse (  id int(11) NOT NULL,  address varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  telephone varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  warehouse_code varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  minimum_capacity int(11) NOT NULL,  minimum_temperature int(11) NOT NULL,  locality_id int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE buyers  ADD PRIMARY KEY (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE carriers  ADD PRIMARY KEY (id),  ADD KEY fk_carriers_locality (locality_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE countries  ADD PRIMARY KEY (Id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE employees  ADD PRIMARY KEY (id),  ADD KEY fk_employees_warehouse (warehouse_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE inbound_orders  ADD PRIMARY KEY (id),  ADD KEY fk_inbound_orders_employee (employe_id),  ADD KEY fk_inbound_orders_product_bash (product_batch_id),  ADD KEY fk_inbound_orders_warehouse (warehouse_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE localities  ADD PRIMARY KEY (id),  ADD KEY fk_provincie (province_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE order_details  ADD PRIMARY KEY (id),  ADD KEY fk_order_details_product_record (product_record),  ADD KEY fk_order_details_purchase_order (purchase_order);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE order_status  ADD PRIMARY KEY (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products  ADD PRIMARY KEY (id),  ADD KEY fk_products_products_types (product_type_id),  ADD KEY fk_products_sellers (seller_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products_batches  ADD PRIMARY KEY (id),  ADD KEY fk_products_batches_products (product_id),  ADD KEY fk_products_batches_section (section_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products_types  ADD PRIMARY KEY (Id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE product_records  ADD PRIMARY KEY (id),  ADD KEY fk_product_records_products (product_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE provinces  ADD PRIMARY KEY (id),  ADD KEY fk_provinces_countrie (id_country_fk);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE purchase_orders  ADD PRIMARY KEY (id),  ADD KEY fk_purchase_orders_buyer (buyer_id),  ADD KEY fk_purchase_orders_carrier (carrier_id),  ADD KEY fk_purchase_orders_order_status (order_status_id),  ADD KEY fk_purchase_orders_wirehouse (warehouse_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE rol  ADD PRIMARY KEY (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sections  ADD PRIMARY KEY (id),  ADD KEY fk_sections_warehouse (warehouse_id),  ADD KEY fk_sections_product_types (product_type_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sellers  ADD PRIMARY KEY (id),  ADD KEY fk_locality (locality_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE users  ADD PRIMARY KEY (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE user_rol  ADD KEY fk_user_rol_users (usuario_id),  ADD KEY fk_user_rol_rol (rol_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE warehouse  ADD PRIMARY KEY (id),  ADD KEY fk_warehouse_locality (locality_id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE buyers  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE carriers  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE countries  MODIFY Id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE employees  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE inbound_orders  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE localities  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE order_details  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE order_status  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products_batches  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products_types  MODIFY Id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE product_records  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE provinces  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE purchase_orders  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE rol  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sections  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sellers  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE users  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE warehouse  MODIFY id int(11) NOT NULL AUTO_INCREMENT;"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE carriers  ADD CONSTRAINT fk_carriers_locality FOREIGN KEY (locality_id) REFERENCES localities (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE employees  ADD CONSTRAINT fk_employees_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouse (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE inbound_orders  ADD CONSTRAINT fk_inbound_orders_employee FOREIGN KEY (employe_id) REFERENCES employees (id),  ADD CONSTRAINT fk_inbound_orders_product_bash FOREIGN KEY (product_batch_id) REFERENCES products_batches (id),  ADD CONSTRAINT fk_inbound_orders_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouse (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE localities  ADD CONSTRAINT fk_provincie FOREIGN KEY (province_id) REFERENCES provinces (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE order_details  ADD CONSTRAINT fk_order_details_product_record FOREIGN KEY (product_record) REFERENCES product_records (id),  ADD CONSTRAINT fk_order_details_purchase_order FOREIGN KEY (purchase_order) REFERENCES purchase_orders (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products  ADD CONSTRAINT fk_products_products_types FOREIGN KEY (product_type_id) REFERENCES products_types (Id),  ADD CONSTRAINT fk_products_sellers FOREIGN KEY (seller_id) REFERENCES sellers (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products_batches  ADD CONSTRAINT fk_products_batches_products FOREIGN KEY (product_id) REFERENCES products (id),  ADD CONSTRAINT fk_products_batches_section FOREIGN KEY (section_id) REFERENCES sections (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE product_records  ADD CONSTRAINT fk_product_records_products FOREIGN KEY (product_id) REFERENCES products (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE provinces  ADD CONSTRAINT fk_provinces_countrie FOREIGN KEY (id_country_fk) REFERENCES countries (Id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE purchase_orders  ADD CONSTRAINT fk_purchase_orders_buyer FOREIGN KEY (buyer_id) REFERENCES buyers (id),  ADD CONSTRAINT fk_purchase_orders_carrier FOREIGN KEY (carrier_id) REFERENCES carriers (id),  ADD CONSTRAINT fk_purchase_orders_order_status FOREIGN KEY (order_status_id) REFERENCES order_status (id),  ADD CONSTRAINT fk_purchase_orders_wirehouse FOREIGN KEY (warehouse_id) REFERENCES warehouse (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sections  ADD CONSTRAINT fk_sections_product_types FOREIGN KEY (product_type_id) REFERENCES products_types (Id),  ADD CONSTRAINT fk_sections_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouse (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sellers  ADD CONSTRAINT fk_locality FOREIGN KEY (locality_id) REFERENCES localities (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE user_rol  ADD CONSTRAINT fk_user_rol_rol FOREIGN KEY (rol_id) REFERENCES rol (id),  ADD CONSTRAINT fk_user_rol_users FOREIGN KEY (usuario_id) REFERENCES users (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE warehouse  ADD CONSTRAINT fk_warehouse_locality FOREIGN KEY (locality_id) REFERENCES localities (id);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE inbound_orders ADD UNIQUE(order_number);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE products ADD CONSTRAINT UNIQUE(product_code);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE sellers ADD UNIQUE(cid);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE employees ADD UNIQUE(id_card_number);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="ALTER TABLE inbound_orders RENAME COLUMN employe_id TO employee_id;"
}

function create_dummy_values {
    # Employees models
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO countries VALUES (1, 'Brasil');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO countries VALUES (2, 'Argentina');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO provinces VALUES (1, 'Sao Paulo', '1');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO localities VALUES (1, 'Sao Paulo', '1');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO warehouse VALUES (1, 'Av Cantareira', '1198765421', '23546', 100, 25, 1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO employees VALUES (1, '34562', 'Vitoria', 'Souza', 1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO sellers VALUES (1, '34562', 'Teste', 'Av Cantareira', '1198765342', 1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO products_types  VALUES (null, 'Teste');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO products VALUES (null, 'SEl', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1');"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO sections VALUES (null, 123, 5, 35, 5, 3, 28, 1, 1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO products_batches VALUES (null, 123, 1, 35, '2022-06-09', 1, '2022-09-09', 10, 23, 1, 1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO inbound_orders VALUES (null, '2022-09-08', '5674', 1, 1, 1);"

    # Products models
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO mercado_fresco_db.countries(country_name)VALUES('Brasil');INSERT INTO mercado_fresco_db.provinces(province_name,id_country_fk)VALUES('São Paulo',1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO mercado_fresco_db.localities(locality_name,province_id)VALUES('São Paulo',1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO mercado_fresco_db.sellers(cid,company_name,address,telephone,locality_id)VALUES('abc','extra','av. paulista','0011223344',1);"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO mercado_fresco_db.products_types(description)VALUES('frutas');alter table products add constraint unique(product_code);INSERT INTO mercado_fresco_db.products(id,description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id)VALUES(10,'carambola',2.1,2.1,2.1,2.1,2.1,'teste',2.1,2.1,1,1);"

    # Warehouse models
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO countries(country_name) VALUES ('Brasil')"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO provinces (province_name, id_country_fk) VALUES ('Sao Paulo', 1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO localities(locality_name, province_id) VALUES ('local 1', 1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO localities(locality_name, province_id) VALUES ('local 3', 1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO localities(locality_name, province_id) VALUES ('local 4', 1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO carriers(cid, company_name, address, locality_id) VALUES (23, 'meli', 'rua 1', 1)"

    # Buyers models
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO buyers(id_card_number,first_name,last_name) VALUES ('order#1', 'Pedro', 'Augusto')"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO countries(country_name) VALUES ('Brasil')"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO provinces(province_name,id_country_fk) VALUES ('São Paulo', 1) "
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO localities(locality_name,province_id) VALUES ('São Paulo',1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO carriers(cid,company_name,address,telephone,locality_id) VALUES ('order#1', 'Meli01', 'Rua Meli 01','(11) 33333333',1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO warehouse(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES ('Rua Melli', '(11) 3333-4444', 'Code#1', 1,1,1)"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO order_status(description) VALUES ('Done')"
    mysql --user=${USER} --password=${PASS} --database="mercado_fresco_db" --execute="INSERT INTO purchase_orders(order_number,order_date,tracking_code,buyer_id,carrier_id,order_status_id,warehouse_id) VALUES ('order#1', '2021-04-04', 'abscf123', 1, 1, 1, 1)"

}

function show_usage {
    echo "Invalid command. Usage:"
    echo "script.sh -c|-p --user USERNAME --password PASSWORD"
}

function check_args {
    if [ ${#ARGS[@]} -ne 5 ]
    then
        show_usage
        exit 1
    fi

    if [[ "${ARGS[0]}" = "-c" || "${ARGS[0]}" = "-p" ]]
    then
        OPTION=("${ARGS[0]}")
    else
          show_usage
        exit 1
    fi

    if [[ "${ARGS[1]}" = "--user" ]]
    then
        USER=${ARGS[2]}
    else
         show_usage
        exit 1
    fi

    if [[ "${ARGS[3]}" = "--password" ]]
    then
        PASS=${ARGS[4]}
        if [[ "$PASS" = "nopass" ]]
        then
            PASS=
        fi
    else
        show_usage
        exit 1
    fi
}


USER=""
PASS=""
OPTION=""

ARGS=("$@")
check_args
if [[ "$OPTION" = "-c" ]]
then
    create_database
else
    create_dummy_values
fi
