-- Active: 1724854584276@@localhost@5433@postgres

INSERT INTO menus (id, name, price, category, created_at, updated_at) 
VALUES (3, 'Menu Item 1', 50.00, 'Category 1', NOW(), NOW()), 
       (5, 'Menu Item 2', 50.50, 'Category 2', NOW(), NOW());

INSERT INTO tables (id, table_number, capacity, created_at, updated_at) VALUES (1, 10, 4, NOW(), NOW());

INSERT INTO roles (id, name) VALUES (1, 'admin');
INSERT INTO roles (id, name) VALUES (2, 'waiter');

INSERT INTO users (username, password, role_id)
VALUES ('saybrhon', 'q123', 1);  -- 1 should be an existing role_id


DELETE FROM order_items WHERE order_id = 1;
ON DELETE CASCADE

SELECT 
    conname AS constraint_name,
    confrelid::regclass AS referenced_table
FROM 
    pg_constraint
WHERE 
    conname LIKE '%fk_orders_items%';


ALTER TABLE order_items
DROP CONSTRAINT fk_orders_items;

ALTER TABLE order_items
ADD CONSTRAINT fk_orders_items
FOREIGN KEY (order_id) REFERENCES orders(id)
ON DELETE CASCADE;


SELECT * FROM roles WHERE id = 2;

ALTER TABLE users ALTER COLUMN role_id DROP NOT NULL;

SELECT * FROM checks

INSERT INTO orders (table_id, total_amount) 
VALUES (1, 100.00); -- Ошибка, если нет записи с id = 1 в таблице tables

SELECT * FROM orders;

INSERT INTO tables (id, name) VALUES (1, 'Table 1');

SELECT * FROM tables

SELECT * FROM tables WHERE id = 1;