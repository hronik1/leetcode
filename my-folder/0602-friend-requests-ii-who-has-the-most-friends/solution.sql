# Write your MySQL query statement below

with A as (
select 
    requester_id as person, 
    sum(1) as num_friends 
from 
    request_accepted ra 
group by 1),

B as (
select 
  accepter_id as person, 
  sum(1) as num_friends
from 
    request_accepted ra 
group by 1 
),

comb as (
select * from a 

union all 

select * from b)

select 
    person as id,
    sum(num_friends) as num
from comb 
group by 1 
order by 2 desc 
limit 1;

