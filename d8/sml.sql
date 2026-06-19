CREATE TABLE orders (

id SERIAL PRIMARY KEY,

order_number VARCHAR(50) UNIQUE NOT NULL,

user_id INT NOT NULL,

total_amount NUMERIC(12, 2) NOT NULL,

created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE

RESTRICT

);

-- Tabel Junction/Associative Table

CREATE TABLE order_items (

order_id INT NOT NULL,

product_id INT NOT NULL,

quantity INT NOT NULL CHECK (quantity > 0),

price NUMERIC(12, 2) NOT NULL,

PRIMARY KEY (order_id, product_id), -- Composite Primary Key

CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE

CASCADE,

CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE

RESTRICT

);

-------------------
-- 1. Membuat tabel Kategori (Tetap sama)

CREATE TABLE categories (

    id SERIAL PRIMARY KEY,

    name VARCHAR(50) NOT NULL UNIQUE

);

 

-- 2. Membuat tabel Produk dengan ID sebagai PK dan SKU sebagai Unique

CREATE TABLE products (

    id SERIAL PRIMARY KEY,                         -- ID Baru sebagai Primary Key

    sku VARCHAR(50) NOT NULL UNIQUE,               -- SKU diubah menjadi UNIQUE

    category_id INT NOT NULL,

    name VARCHAR(100) NOT NULL,

    price NUMERIC(12, 2) NOT NULL CHECK (price >= 0),

    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    

    CONSTRAINT fk_product_category

        FOREIGN KEY (category_id)

        REFERENCES categories(id)

        ON DELETE RESTRICT

);