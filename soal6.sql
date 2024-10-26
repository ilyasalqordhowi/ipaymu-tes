-- SELECT 
--     c.customer_name, 
--     p.product_name, 
--     SUM(o.total_price) AS total_spent 
-- FROM 
--     Customers c
-- JOIN 
--     Orders o ON c.customer_id = o.customer_id
-- JOIN 
--     Products p ON o.product_id = p.product_id
-- WHERE 
--     c.city = 'New York'
-- GROUP BY 
--     c.customer_name, p.product_name;

--perintah rendering
CREATE INDEX idx_customers_customer_id ON Customers(customer_id);
CREATE INDEX idx_customers_city ON Customers(city);
CREATE INDEX idx_orders_customer_id ON Orders(customer_id);
CREATE INDEX idx_orders_product_id ON Orders(product_id);
CREATE INDEX idx_products_product_id ON Products(product_id);

SELECT 
    c.customer_name, 
    p.product_name, 
    SUM(o.total_price) AS total_spent 
FROM 
    Customers c
JOIN 
    Orders o ON c.customer_id = o.customer_id
JOIN 
    Products p ON o.product_id = p.product_id
WHERE 
    c.city = 'New York'
GROUP BY 
    c.customer_name, p.product_name;

--menggunakan explain plan
    EXPLAIN ANALYZE 
SELECT 
    c.customer_name, 
    p.product_name, 
    SUM(o.total_price) AS total_spent 
FROM 
    Customers c
JOIN 
    Orders o ON c.customer_id = o.customer_id
JOIN 
    Products p ON o.product_id = p.product_id
WHERE 
    c.city = 'New York'
GROUP BY 
    c.customer_name, p.product_name;

