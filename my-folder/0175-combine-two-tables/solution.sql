# Write your MySQL query statement below
select 
    p.FirstName,
    p.LastName,
    City, 
    State
from Person p 
left join Address a     
    on p.PersonId = a.PersonId
    
