CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL, 
    password TEXT NOT NULL,
    mobile TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    refresh_token TEXT,
    last_active timestamptz NOT NULL DEFAULT NOW(),
    UNIQUE (mobile, email)
);

INSERT INTO "users" ("id", "username", "password", "mobile", "email", "refresh_token") VALUES
(1, 'vlad', '7465737485d5c8dcb7db39358916efb7e6f840502f547b2c', '+380500106072', 'vlad1k.zhuchkov@gmail.com', '91d85ce5222682752ac16cdcb0e13098b4fd89c45ea717da8126153a99907ff4');


CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY UNIQUE,
    name VARCHAR(255)
);

INSERT INTO category(name) VALUES
('phone'),
('tablet'),
('notebook'),
('tv');

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY UNIQUE,
    category_id INT NOT NULL REFERENCES category,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    count INT NOT NULL DEFAULT 1,
    amount NUMERIC(15,6) NOT NULL DEFAULT 0.00,
    avaibility BOOLEAN NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()      
);

INSERT INTO items(category_id, name, description, count, amount, avaibility, created_at, updated_at) VALUES
(1, 'Samsung A32', 'New Samsung A32', 4, 25000.53, true, NOW(), NOW()),
(1, 'Apple iPhone 11', 'New Apple iPhone 11', 10, 24500.00, true, NOW(), NOW()),
(2, 'Lenovo YOGA 2', 'New Lenovo YOGA 2', 3, 12700.75, true, NOW(), NOW()),
(2, 'iPad Pro 2022', 'New iPad Pro 2022', 5, 50800.00, false, NOW(), NOW()),
(4, 'Xiaomi TV 2022', 'New Xiaomi TV 2022', 1, 70000.00, true, NOW(), NOW()),
(3, 'Lenovo Z510', 'New Lenovo Z510', 1, 8000.50, true, NOW(), NOW());


CREATE TABLE IF NOT EXISTS status (
    id SERIAL PRIMARY KEY UNIQUE,
    name VARCHAR(255) NOT NULL
);

INSERT INTO status(name) VALUES
('done'),
('new'),
('in progress'),
('invalid');


CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY UNIQUE,
    user_id INT REFERENCES users,
    item_id INT REFERENCES items,
    status_id INT REFERENCES status,
    count INT NOT NULL DEFAULT 1,
    amount NUMERIC(15,6) NOT NULL DEFAULT 0.00,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW() 
);
