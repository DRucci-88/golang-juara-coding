/*
 Navicat Premium Data Transfer

 Source Server         : juaracoding
 Source Server Type    : PostgreSQL
 Source Server Version : 180003 (180003)
 Source Host           : localhost:5434
 Source Catalog        : ecommercedb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 180003 (180003)
 File Encoding         : 65001

 Date: 18/06/2026 19:44:36
*/


-- ----------------------------
-- Sequence structure for categories_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."categories_id_seq";
CREATE SEQUENCE "public"."categories_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for orders_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."orders_id_seq";
CREATE SEQUENCE "public"."orders_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for products_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."products_id_seq";
CREATE SEQUENCE "public"."products_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_profiles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_profiles_id_seq";
CREATE SEQUENCE "public"."user_profiles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS "public"."categories";
CREATE TABLE "public"."categories" (
  "id" int4 NOT NULL DEFAULT nextval('categories_id_seq'::regclass),
  "name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO "public"."categories" VALUES (1, 'Otomotif');
INSERT INTO "public"."categories" VALUES (2, 'Elektronik');
INSERT INTO "public"."categories" VALUES (3, 'Pakaian');
INSERT INTO "public"."categories" VALUES (4, 'Makanan');
INSERT INTO "public"."categories" VALUES (5, 'Buku');

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS "public"."order_items";
CREATE TABLE "public"."order_items" (
  "order_id" int4 NOT NULL,
  "product_id" int4 NOT NULL,
  "quantity" int4 NOT NULL,
  "price" numeric(12,2) NOT NULL
)
;

-- ----------------------------
-- Records of order_items
-- ----------------------------

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "public"."orders";
CREATE TABLE "public"."orders" (
  "id" int4 NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
  "order_number" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "total_amount" numeric(12,2) NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;

-- ----------------------------
-- Records of orders
-- ----------------------------

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "public"."products";
CREATE TABLE "public"."products" (
  "id" int4 NOT NULL DEFAULT nextval('products_id_seq'::regclass),
  "sku" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "category_id" int4 NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "price" numeric(12,2) NOT NULL,
  "stock" int4 NOT NULL DEFAULT 0,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO "public"."products" VALUES (1, 'SKU-5fc1-11', 1, 'Oli Mesin Shell Helix HX8', 150000.00, 47, '2026-06-18 12:37:40.914248');
INSERT INTO "public"."products" VALUES (2, 'SKU-d971-12', 1, 'Helm Full Face KYT Vendetta', 600000.00, 71, '2026-06-18 12:37:40.928312');
INSERT INTO "public"."products" VALUES (3, 'SKU-d8bd-13', 1, 'Kampas Rem Brembo Original', 150000.00, 42, '2026-06-18 12:37:40.932752');
INSERT INTO "public"."products" VALUES (4, 'SKU-d16d-14', 1, 'Ban Motor Maxxis Extramaxx', 45000.00, 82, '2026-06-18 12:37:40.937296');
INSERT INTO "public"."products" VALUES (5, 'SKU-e525-15', 1, 'Parfum Mobil Little Trees', 105000.00, 87, '2026-06-18 12:37:40.941461');
INSERT INTO "public"."products" VALUES (6, 'SKU-5ae9-16', 1, 'Lap Microfiber Pembersih', 510000.00, 27, '2026-06-18 12:37:40.946884');
INSERT INTO "public"."products" VALUES (7, 'SKU-cefc-17', 1, 'Busi Denso Iridium', 180000.00, 87, '2026-06-18 12:37:40.95386');
INSERT INTO "public"."products" VALUES (8, 'SKU-d707-18', 1, 'Jas Hujan Axio Rubber', 540000.00, 83, '2026-06-18 12:37:40.974086');
INSERT INTO "public"."products" VALUES (9, 'SKU-3c67-19', 1, 'Car Charger Anker Dual Port', 45000.00, 49, '2026-06-18 12:37:40.986234');
INSERT INTO "public"."products" VALUES (10, 'SKU-9f6c-110', 1, 'Cairan Pembersih Jamur Kaca', 180000.00, 62, '2026-06-18 12:37:40.991157');
INSERT INTO "public"."products" VALUES (11, 'SKU-e1fb-21', 2, 'Smartphone iPhone 15 Pro', 435000.00, 86, '2026-06-18 12:37:41.00111');
INSERT INTO "public"."products" VALUES (12, 'SKU-d4e4-22', 2, 'Samsung Galaxy S24 Ultra', 165000.00, 66, '2026-06-18 12:37:41.00641');
INSERT INTO "public"."products" VALUES (13, 'SKU-b972-23', 2, 'Laptop ASUS ROG Zephyrus', 285000.00, 76, '2026-06-18 12:37:41.011926');
INSERT INTO "public"."products" VALUES (14, 'SKU-4c7a-24', 2, 'TWS Sony WF-1000XM5', 225000.00, 72, '2026-06-18 12:37:41.019835');
INSERT INTO "public"."products" VALUES (15, 'SKU-13df-25', 2, 'Smart TV LG OLED 55 Inch', 270000.00, 38, '2026-06-18 12:37:41.02521');
INSERT INTO "public"."products" VALUES (16, 'SKU-d39f-26', 2, 'iPad Air M2', 345000.00, 58, '2026-06-18 12:37:41.031028');
INSERT INTO "public"."products" VALUES (17, 'SKU-c992-27', 2, 'Keyboard Mechanical Keychron K2', 135000.00, 74, '2026-06-18 12:37:41.035881');
INSERT INTO "public"."products" VALUES (18, 'SKU-9606-28', 2, 'Mouse Logistics MX Master 3S', 690000.00, 36, '2026-06-18 12:37:41.04056');
INSERT INTO "public"."products" VALUES (19, 'SKU-f2d4-29', 2, 'Monitor Dell UltraSharp 27', 255000.00, 65, '2026-06-18 12:37:41.046871');
INSERT INTO "public"."products" VALUES (20, 'SKU-2443-210', 2, 'PlayStation 5 Slim', 720000.00, 32, '2026-06-18 12:37:41.053254');
INSERT INTO "public"."products" VALUES (21, 'SKU-0d2c-31', 3, 'Kemeja Flanel Uniqlo', 585000.00, 83, '2026-06-18 12:37:41.064211');
INSERT INTO "public"."products" VALUES (22, 'SKU-8135-32', 3, 'Jaket Denim Levi''s 501', 390000.00, 76, '2026-06-18 12:37:41.069517');
INSERT INTO "public"."products" VALUES (23, 'SKU-431c-33', 3, 'Celana Chino Erigo', 390000.00, 80, '2026-06-18 12:37:41.074734');
INSERT INTO "public"."products" VALUES (24, 'SKU-812d-34', 3, 'Kaos Polos Cotton Combed 30s', 510000.00, 83, '2026-06-18 12:37:41.080347');
INSERT INTO "public"."products" VALUES (25, 'SKU-e4ee-35', 3, 'Sweater Hoodie H&M', 600000.00, 32, '2026-06-18 12:37:41.086431');
INSERT INTO "public"."products" VALUES (26, 'SKU-418c-36', 3, 'Sepatu Sneakers Nike Air Jordan', 15000.00, 47, '2026-06-18 12:37:41.091289');
INSERT INTO "public"."products" VALUES (27, 'SKU-9928-37', 3, 'Celana Cargo Tactical', 525000.00, 22, '2026-06-18 12:37:41.096321');
INSERT INTO "public"."products" VALUES (28, 'SKU-e6da-38', 3, 'Rok Plisket Panjang', 120000.00, 13, '2026-06-18 12:37:41.100984');
INSERT INTO "public"."products" VALUES (29, 'SKU-0377-39', 3, 'Batik Pria Lengan Panjang', 720000.00, 82, '2026-06-18 12:37:41.10608');
INSERT INTO "public"."products" VALUES (30, 'SKU-0c5f-310', 3, 'Sepatu Pantofel Kulit', 435000.00, 46, '2026-06-18 12:37:41.111512');
INSERT INTO "public"."products" VALUES (31, 'SKU-87ee-41', 4, 'Indomie Goreng Aceh', 585000.00, 84, '2026-06-18 12:37:41.122383');
INSERT INTO "public"."products" VALUES (32, 'SKU-9aa9-42', 4, 'Keripik Singkong Maicih', 105000.00, 67, '2026-06-18 12:37:41.127556');
INSERT INTO "public"."products" VALUES (33, 'SKU-a1fa-43', 4, 'Cokelat Silverqueen Almond', 345000.00, 21, '2026-06-18 12:37:41.132384');
INSERT INTO "public"."products" VALUES (34, 'SKU-c784-44', 4, 'Roti Tawar Sari Roti', 735000.00, 82, '2026-06-18 12:37:41.137659');
INSERT INTO "public"."products" VALUES (35, 'SKU-0a6b-45', 4, 'Susu UHT Ultra Milk 1L', 540000.00, 12, '2026-06-18 12:37:41.142833');
INSERT INTO "public"."products" VALUES (36, 'SKU-b771-46', 4, 'Kopi Kenangan Mantan Bottle', 420000.00, 55, '2026-06-18 12:37:41.149482');
INSERT INTO "public"."products" VALUES (37, 'SKU-103a-47', 4, 'Biskuit Oreo Vanilla', 195000.00, 72, '2026-06-18 12:37:41.155986');
INSERT INTO "public"."products" VALUES (38, 'SKU-9cb0-48', 4, 'Sereal Kellogg''s Corn Flakes', 330000.00, 58, '2026-06-18 12:37:41.161231');
INSERT INTO "public"."products" VALUES (39, 'SKU-f37a-49', 4, 'Samyang Buldak Noodles', 705000.00, 43, '2026-06-18 12:37:41.166308');
INSERT INTO "public"."products" VALUES (40, 'SKU-956b-410', 4, 'Selai Nutella 350g', 375000.00, 27, '2026-06-18 12:37:41.171276');
INSERT INTO "public"."products" VALUES (41, 'SKU-6241-51', 5, 'Buku Atomic Habits - James Clear', 225000.00, 73, '2026-06-18 12:37:41.18472');
INSERT INTO "public"."products" VALUES (42, 'SKU-c076-52', 5, 'Buku Filosofi Teras - Henry Manampiring', 300000.00, 72, '2026-06-18 12:37:41.190122');
INSERT INTO "public"."products" VALUES (43, 'SKU-e748-53', 5, 'Novel Bumi - Tere Liye', 660000.00, 77, '2026-06-18 12:37:41.195258');
INSERT INTO "public"."products" VALUES (44, 'SKU-d9e8-54', 5, 'Buku Berani Tidak Disukai', 285000.00, 36, '2026-06-18 12:37:41.200153');
INSERT INTO "public"."products" VALUES (45, 'SKU-9c50-55', 5, 'Buku Laut Bercerita - Leila S. Chudori', 585000.00, 81, '2026-06-18 12:37:41.204655');
INSERT INTO "public"."products" VALUES (46, 'SKU-a7c0-56', 5, 'Buku Rich Dad Poor Dad', 105000.00, 44, '2026-06-18 12:37:41.209277');
INSERT INTO "public"."products" VALUES (47, 'SKU-394a-57', 5, 'Novel Gadis Kretek', 750000.00, 73, '2026-06-18 12:37:41.215794');
INSERT INTO "public"."products" VALUES (48, 'SKU-10b5-58', 5, 'Buku Sebuah Seni untuk Bersikap Bodo Amat', 585000.00, 57, '2026-06-18 12:37:41.221558');
INSERT INTO "public"."products" VALUES (49, 'SKU-1841-59', 5, 'Buku Psikologi Uang', 615000.00, 70, '2026-06-18 12:37:41.227206');
INSERT INTO "public"."products" VALUES (50, 'SKU-1343-510', 5, 'Buku Sapiens - Yuval Noah Harari', 735000.00, 87, '2026-06-18 12:37:41.232591');

-- ----------------------------
-- Table structure for user_profiles
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_profiles";
CREATE TABLE "public"."user_profiles" (
  "id" int4 NOT NULL DEFAULT nextval('user_profiles_id_seq'::regclass),
  "user_id" int4 NOT NULL,
  "full_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_number" varchar(15) COLLATE "pg_catalog"."default",
  "address" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of user_profiles
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "email" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;

-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."categories_id_seq"
OWNED BY "public"."categories"."id";
SELECT setval('"public"."categories_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."orders_id_seq"
OWNED BY "public"."orders"."id";
SELECT setval('"public"."orders_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."products_id_seq"
OWNED BY "public"."products"."id";
SELECT setval('"public"."products_id_seq"', 50, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_profiles_id_seq"
OWNED BY "public"."user_profiles"."id";
SELECT setval('"public"."user_profiles_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 1, false);

-- ----------------------------
-- Uniques structure for table categories
-- ----------------------------
ALTER TABLE "public"."categories" ADD CONSTRAINT "categories_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table categories
-- ----------------------------
ALTER TABLE "public"."categories" ADD CONSTRAINT "categories_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Checks structure for table order_items
-- ----------------------------
ALTER TABLE "public"."order_items" ADD CONSTRAINT "order_items_quantity_check" CHECK (quantity > 0);

-- ----------------------------
-- Primary Key structure for table order_items
-- ----------------------------
ALTER TABLE "public"."order_items" ADD CONSTRAINT "order_items_pkey" PRIMARY KEY ("order_id", "product_id");

-- ----------------------------
-- Uniques structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "orders_order_number_key" UNIQUE ("order_number");

-- ----------------------------
-- Primary Key structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "orders_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "products_sku_key" UNIQUE ("sku");

-- ----------------------------
-- Checks structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "products_price_check" CHECK (price >= 0::numeric);
ALTER TABLE "public"."products" ADD CONSTRAINT "products_stock_check" CHECK (stock >= 0);

-- ----------------------------
-- Primary Key structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "products_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table user_profiles
-- ----------------------------
ALTER TABLE "public"."user_profiles" ADD CONSTRAINT "user_profiles_user_id_key" UNIQUE ("user_id");

-- ----------------------------
-- Primary Key structure for table user_profiles
-- ----------------------------
ALTER TABLE "public"."user_profiles" ADD CONSTRAINT "user_profiles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_email_key" UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table order_items
-- ----------------------------
ALTER TABLE "public"."order_items" ADD CONSTRAINT "fk_order" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "public"."order_items" ADD CONSTRAINT "fk_product" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "fk_product_category" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table user_profiles
-- ----------------------------
ALTER TABLE "public"."user_profiles" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
