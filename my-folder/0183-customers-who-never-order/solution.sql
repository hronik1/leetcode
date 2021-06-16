# Write your MySQL query statement below
select Name as Customers from Customers LEFT OUTER JOIN Orders on Customers.ID = Orders.CustomerID where Orders.ID is NULL
