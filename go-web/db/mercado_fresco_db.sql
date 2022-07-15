
DROP DATABASE IF EXISTS mercado_fresco_db;
CREATE DATABASE mercado_fresco_db;
USE mercado_fresco_db;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

-- --------------------------------------------------------

--
-- Estrutura da tabela `buyers`
--

CREATE TABLE `buyers` (
  `id` int(11) NOT NULL,
  `id_card_number` varchar(255) UNIQUE COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `carriers`
--

CREATE TABLE `carriers` (
  `id` int(11) NOT NULL,
  `cid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `company_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `telephone` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `locality_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `countries`
--

CREATE TABLE `countries` (
  `Id` int(11) NOT NULL,
  `country_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `employees`
--

CREATE TABLE `employees` (
  `id` int(11) NOT NULL,
  `id_card_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `warehouse_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `inbound_orders`
--

CREATE TABLE `inbound_orders` (
  `id` int(11) NOT NULL,
  `order_date` datetime(6) NOT NULL,
  `order_number` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `employe_id` int(11) NOT NULL,
  `product_batch_id` int(11) NOT NULL,
  `warehouse_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `localities`
--

CREATE TABLE `localities` (
  `id` int(11) NOT NULL,
  `locality_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `province_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `order_details`
--

CREATE TABLE `order_details` (
  `id` int(11) NOT NULL,
  `clean_liness_status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `quantity` int(11) NOT NULL,
  `temperature` decimal(19,2) NOT NULL,
  `product_record` int(11) NOT NULL,
  `purchase_order` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `order_status`
--

CREATE TABLE `order_status` (
  `id` int(11) NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `expiration_rate` decimal(19,2) DEFAULT NULL,
  `freezing_rate` decimal(19,2) DEFAULT NULL,
  `height` decimal(19,2) DEFAULT NULL,
  `length` decimal(19,2) DEFAULT NULL,
  `net_weight` decimal(19,2) DEFAULT NULL,
  `product_code` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `recommended_freezing_temperature` decimal(19,2) DEFAULT NULL,
  `width` decimal(19,2) DEFAULT NULL,
  `product_type_id` int(11) NOT NULL,
  `seller_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `products_batches`
--

CREATE TABLE `products_batches` (
  `id` int(11) NOT NULL,
  `batch_number` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `current_number` int(11) DEFAULT NULL,
  `current_tempertature` decimal(19,2) DEFAULT NULL,
  `due_date` datetime(6) DEFAULT NULL,
  `initial_quantity` int(11) DEFAULT NULL,
  `manufacturing_date` datetime(6) NOT NULL,
  `manufacturing_hour` datetime(6) NOT NULL,
  `minimun_temperature` decimal(19,2) DEFAULT NULL,
  `product_id` int(11) NOT NULL,
  `section_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `products_types`
--

CREATE TABLE `products_types` (
  `Id` int(11) NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `product_records`
--

CREATE TABLE `product_records` (
  `id` int(11) NOT NULL,
  `last_update_date` datetime(6) NOT NULL,
  `purchase_prince` decimal(19,2) NOT NULL,
  `sale_price` decimal(19,2) NOT NULL,
  `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `provinces`
--

CREATE TABLE `provinces` (
  `id` int(11) NOT NULL,
  `province_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `id_country_fk` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `purchase_orders`
--

CREATE TABLE `purchase_orders` (
  `id` int(11) NOT NULL,
  `order_number` varchar(255) UNIQUE COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `order_date` datetime(6) DEFAULT NULL,
  `tracking_code` varchar(255) UNIQUE COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `buyer_id` int(11) NOT NULL,
  `carrier_id` int(11) NOT NULL,
  `order_status_id` int(11) NOT NULL,
  `warehouse_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `rol`
--

CREATE TABLE `rol` (
  `id` int(11) NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rol_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `sections`
--

CREATE TABLE `sections` (
  `id` int(11) NOT NULL,
  `section_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `current_capacity` int(11) NOT NULL,
  `current_temperature` decimal(19,2) NOT NULL,
  `maximum_capacity` int(11) NOT NULL,
  `minimum_capacity` int(11) NOT NULL,
  `minimum_temperature` decimal(19,2) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `warehouse_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `sellers`
--

CREATE TABLE `sellers` (
  `id` int(11) NOT NULL,
  `cid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `company_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `telephone` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `locality_id` int(11) NOT NULL,
   UNIQUE(`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `user_rol`
--

CREATE TABLE `user_rol` (
  `usuario_id` int(11) DEFAULT NULL,
  `rol_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `warehouse`
--

CREATE TABLE `warehouse` (
  `id` int(11) NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `telephone` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `warehouse_code` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `locality_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `buyers`
--
ALTER TABLE `buyers`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `carriers`
--
ALTER TABLE `carriers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_carriers_locality` (`locality_id`);

--
-- Índices para tabela `countries`
--
ALTER TABLE `countries`
  ADD PRIMARY KEY (`Id`);

--
-- Índices para tabela `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_employees_warehouse` (`warehouse_id`);

--
-- Índices para tabela `inbound_orders`
--
ALTER TABLE `inbound_orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_inbound_orders_employee` (`employe_id`),
  ADD KEY `fk_inbound_orders_product_bash` (`product_batch_id`),
  ADD KEY `fk_inbound_orders_warehouse` (`warehouse_id`);

--
-- Índices para tabela `localities`
--
ALTER TABLE `localities`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_provincie` (`province_id`);

--
-- Índices para tabela `order_details`
--
ALTER TABLE `order_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_order_details_product_record` (`product_record`),
  ADD KEY `fk_order_details_purchase_order` (`purchase_order`);

--
-- Índices para tabela `order_status`
--
ALTER TABLE `order_status`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_products_products_types` (`product_type_id`),
  ADD KEY `fk_products_sellers` (`seller_id`);

--
-- Índices para tabela `products_batches`
--
ALTER TABLE `products_batches`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_products_batches_products` (`product_id`),
  ADD KEY `fk_products_batches_section` (`section_id`);

--
-- Índices para tabela `products_types`
--
ALTER TABLE `products_types`
  ADD PRIMARY KEY (`Id`);

--
-- Índices para tabela `product_records`
--
ALTER TABLE `product_records`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_product_records_products` (`product_id`);

--
-- Índices para tabela `provinces`
--
ALTER TABLE `provinces`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_provinces_countrie` (`id_country_fk`);

--
-- Índices para tabela `purchase_orders`
--
ALTER TABLE `purchase_orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_purchase_orders_buyer` (`buyer_id`),
  ADD KEY `fk_purchase_orders_carrier` (`carrier_id`),
  ADD KEY `fk_purchase_orders_order_status` (`order_status_id`),
  ADD KEY `fk_purchase_orders_wirehouse` (`warehouse_id`);

--
-- Índices para tabela `rol`
--
ALTER TABLE `rol`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `sections`
--
ALTER TABLE `sections`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_sections_warehouse` (`warehouse_id`),
  ADD KEY `fk_sections_product_types` (`product_type_id`);

--
-- Índices para tabela `sellers`
--
ALTER TABLE `sellers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_locality` (`locality_id`);

--
-- Índices para tabela `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `user_rol`
--
ALTER TABLE `user_rol`
  ADD KEY `fk_user_rol_users` (`usuario_id`),
  ADD KEY `fk_user_rol_rol` (`rol_id`);

--
-- Índices para tabela `warehouse`
--
ALTER TABLE `warehouse`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_warehouse_locality` (`locality_id`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `buyers`
--
ALTER TABLE `buyers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `carriers`
--
ALTER TABLE `carriers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `countries`
--
ALTER TABLE `countries`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `employees`
--
ALTER TABLE `employees`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `inbound_orders`
--
ALTER TABLE `inbound_orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `localities`
--
ALTER TABLE `localities`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `order_details`
--
ALTER TABLE `order_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `order_status`
--
ALTER TABLE `order_status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `products_batches`
--
ALTER TABLE `products_batches`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `products_types`
--
ALTER TABLE `products_types`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `product_records`
--
ALTER TABLE `product_records`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `provinces`
--
ALTER TABLE `provinces`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `purchase_orders`
--
ALTER TABLE `purchase_orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `rol`
--
ALTER TABLE `rol`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `sections`
--
ALTER TABLE `sections`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `sellers`
--
ALTER TABLE `sellers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `warehouse`
--
ALTER TABLE `warehouse`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Restrições para despejos de tabelas
--

--
-- Limitadores para a tabela `carriers`
--
ALTER TABLE `carriers`
  ADD CONSTRAINT `fk_carriers_locality` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

--
-- Limitadores para a tabela `employees`
--
ALTER TABLE `employees`
  ADD CONSTRAINT `fk_employees_warehouse` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

--
-- Limitadores para a tabela `inbound_orders`
--
ALTER TABLE `inbound_orders`
  ADD CONSTRAINT `fk_inbound_orders_employee` FOREIGN KEY (`employe_id`) REFERENCES `employees` (`id`),
  ADD CONSTRAINT `fk_inbound_orders_product_bash` FOREIGN KEY (`product_batch_id`) REFERENCES `products_batches` (`id`),
  ADD CONSTRAINT `fk_inbound_orders_warehouse` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

--
-- Limitadores para a tabela `localities`
--
ALTER TABLE `localities`
  ADD CONSTRAINT `fk_province` FOREIGN KEY (`province_id`) REFERENCES `provinces` (`id`);

--
-- Limitadores para a tabela `order_details`
--
ALTER TABLE `order_details`
  ADD CONSTRAINT `fk_order_details_product_record` FOREIGN KEY (`product_record`) REFERENCES `product_records` (`id`),
  ADD CONSTRAINT `fk_order_details_purchase_order` FOREIGN KEY (`purchase_order`) REFERENCES `purchase_orders` (`id`);

--
-- Limitadores para a tabela `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_products_products_types` FOREIGN KEY (`product_type_id`) REFERENCES `products_types` (`Id`),
  ADD CONSTRAINT `fk_products_sellers` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`);

--
-- Limitadores para a tabela `products_batches`
--
ALTER TABLE `products_batches`
  ADD CONSTRAINT `fk_products_batches_products` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_products_batches_section` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`);

--
-- Limitadores para a tabela `product_records`
--
ALTER TABLE `product_records`
  ADD CONSTRAINT `fk_product_records_products` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Limitadores para a tabela `provinces`
--
ALTER TABLE `provinces`
  ADD CONSTRAINT `fk_provinces_countrie` FOREIGN KEY (`id_country_fk`) REFERENCES `countries` (`Id`);

--
-- Limitadores para a tabela `purchase_orders`
--
ALTER TABLE `purchase_orders`
  ADD CONSTRAINT `fk_purchase_orders_buyer` FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`),
  ADD CONSTRAINT `fk_purchase_orders_carrier` FOREIGN KEY (`carrier_id`) REFERENCES `carriers` (`id`),
  ADD CONSTRAINT `fk_purchase_orders_order_status` FOREIGN KEY (`order_status_id`) REFERENCES `order_status` (`id`),
  ADD CONSTRAINT `fk_purchase_orders_wirehouse` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

--
-- Limitadores para a tabela `sections`
--
ALTER TABLE `sections`
  ADD CONSTRAINT `fk_sections_product_types` FOREIGN KEY (`product_type_id`) REFERENCES `products_types` (`Id`),
  ADD CONSTRAINT `fk_sections_warehouse` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

--
-- Limitadores para a tabela `sellers`
--
ALTER TABLE `sellers`
  ADD CONSTRAINT `fk_locality` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

--
-- Limitadores para a tabela `user_rol`
--
ALTER TABLE `user_rol`
  ADD CONSTRAINT `fk_user_rol_rol` FOREIGN KEY (`rol_id`) REFERENCES `rol` (`id`),
  ADD CONSTRAINT `fk_user_rol_users` FOREIGN KEY (`usuario_id`) REFERENCES `users` (`id`);

--
-- Limitadores para a tabela `warehouse`
--
ALTER TABLE `warehouse`
  ADD CONSTRAINT `fk_warehouse_locality` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);
COMMIT;


INSERT INTO buyers(id_card_number,first_name,last_name) 
	VALUES ('order#1', 'Pedro', 'Augusto');

  INSERT INTO countries(country_name) 
	VALUES ('Brasil') ; 

INSERT INTO provinces(province_name,id_country_fk) 
	VALUES ('São Paulo', 1) ; 
    
 INSERT INTO localities(locality_name,province_id) 
	VALUES ('São Paulo',1) ; 

INSERT INTO carriers(cid,company_name,address,telephone,locality_id) 
 	VALUES ('order#1', 'Meli01', 'Rua Meli 01','(11) 33333333',1) ; 

INSERT INTO warehouse(address,telephone,warehouse_code,locality_id) 
	VALUES ('Rua Melli', '(11) 3333-4444', 'Code#1', 1);

INSERT INTO order_status(`description`) 
	VALUES ('Done');


INSERT INTO purchase_orders(order_number,order_date,tracking_code,buyer_id,carrier_id,order_status_id,warehouse_id) 
	VALUES ('order#1', '2021-04-04', 'abscf123', 1, 1, 1, 1)



