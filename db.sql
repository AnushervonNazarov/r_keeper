-- Active: 1724854584276@@localhost@5433@postgres

INSERT INTO menus (id, name, price, category, created_at, updated_at) 
VALUES (3, 'Menu Item 1', 50.00, 'Category 1', NOW(), NOW()), 
       (5, 'Menu Item 2', 50.50, 'Category 2', NOW(), NOW());

INSERT INTO tables (id, table_number, capacity, created_at, updated_at) VALUES (1, 10, 4, NOW(), NOW());

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
