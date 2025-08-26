-- ==========================
-- Users Table
-- ==========================
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INT,
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
    CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users(id)
);

-- ==========================
-- Products Table
-- ==========================
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INT REFERENCES users(id)
);

-- ==========================
-- Orders Table
-- ==========================
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id),
    user_id INT NOT NULL REFERENCES users(id),
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INT REFERENCES users(id)
);

-- ==========================
-- Seed Users
-- ==========================
INSERT INTO users (username, password)
VALUES 
  ('admin', 'admin123'),
  ('john', 'john123'),
  ('jane', 'jane123')
ON CONFLICT DO NOTHING;

-- ==========================
-- Seed Products
-- ==========================
INSERT INTO products (name, quantity, created_by, updated_by)
VALUES
  ('Laptop', 10, 1, 1),
  ('Mouse', 50, 1, 1),
  ('Keyboard', 30, 1, 1)
ON CONFLICT DO NOTHING;

-- ==========================
-- Seed Orders
-- ==========================
INSERT INTO orders (product_id, user_id, status, created_by, updated_by)
VALUES
  (1, 2, 'PENDING', 1, 1),
  (2, 3, 'PENDING', 1, 1)
ON CONFLICT DO NOTHING;
