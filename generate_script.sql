create or replace function random_integer(int, int) returns int language sql as 
$$
select round($1 + (random() * ($2 - $1)))::int;
$$;

create or replace function random_price(int, int) returns numeric language sql as 
$$
select round(($1 + (random() * ($2 - $1)))::numeric, 2);
$$;

create or replace function random_product_name(int) returns varchar language sql as 
$$
select 'laptop_' || 
case floor(random() * 4)::int 
	when 0 then 'dell' 
	when 1 then 'lenovo'
	when 2 then 'apple'
	when 3 then 'samsung'
end || '_' || $1;
$$;

create or replace function random_search_timestamp() returns timestamp language sql as 
$$
select (date '2016-01-01' + floor(random() * 365 * 5)::int)::timestamp;
$$;

create or replace function random_varchar(int) returns varchar language sql as 
$$
select upper(substr(md5(random()::text), 0, $1 + 1));
$$;

create or replace function random_order_status() returns order_status language sql as 
$$
select os from  unnest(enum_range(null::order_status)) os order by random() limit 1;
$$;


do $$
declare
	USER_COUNT int := 1000;
	PRODUCT_COUNT int := 100000;
	SEARCH_COUNT int := 100000;
	ORDER_COUNT int := 10000;
begin

insert into
	users (email)
select
	'user' || i || '@gmail.com'
from
	generate_series(1, USER_COUNT) as seq(i);

insert into
	products (name, price, quantity)
select
	random_product_name(i) as name,
	random_price(0,1000) as price,
	random_integer(1,10) as quantity
from
	generate_series(1, PRODUCT_COUNT) as seq(i);

insert into
	products (name, price, quantity)
values 
	('laptop', 10, 100);


insert into 
	attributes (name)
values
	('size'),
	('color'),
	('weight');


insert into 
	attribute_values (attribute_id, value)
values
	(1, 'small'),
	(1, 'medium'),
	(1, 'large');

insert into
	product_attributes (product_id, attribute_id, attribute_value_id)
select	
	random_integer(1, PRODUCT_COUNT) as product_id,
	1 as attribute_id,
	random_integer(1, 3) as attribute_value_id
from
	generate_series(1, PRODUCT_COUNT/2)
on 
	conflict (product_id, attribute_id) do nothing;

insert into 
	attribute_values (attribute_id, value)
values
	(2, 'red'),
	(2, 'green'),
	(2, 'blue'),
	(2, 'yellow'),
	(2, 'black'),
	(2, 'brown'),
	(2, 'white'),
	(2, 'pink');

insert into
	product_attributes (product_id, attribute_id, attribute_value_id)
select	
	random_integer(1, PRODUCT_COUNT) as product_id,
	2 as attribute_id,
	random_integer(4, 11) as attribute_value_id
from
	generate_series(1, PRODUCT_COUNT/2)
on 
	conflict (product_id, attribute_id) do nothing;


insert into 
	attribute_values (attribute_id, value)
select 
	3 as attribute_id,
	random_price(0,100)::varchar as value
from
	generate_series(1, PRODUCT_COUNT) as seq(i);

insert into
	product_attributes (product_id, attribute_id, attribute_value_id)
select	
	random_integer(1, PRODUCT_COUNT) as product_id,
	3 as attribute_id,
	random_integer(12, 12+PRODUCT_COUNT-1) as attribute_value_id
from
	generate_series(1, PRODUCT_COUNT/2)
on 
	conflict (product_id, attribute_id) do nothing;


insert into
	basket_items (user_id, product_id, quantity)
select
	random_integer(1, USER_COUNT) as user_id,
	random_integer(1, PRODUCT_COUNT) as product_id,
	random_integer(1, 10) as quantity
from
	generate_series(1, PRODUCT_COUNT * 2) as seq(i)
on 
	conflict (user_id, product_id) do nothing;


insert into
	search_log (query)
values
	('dell'),
	('apple'),
	('potato'),
	('laptop'); 


insert into
	search_log_timestamps (search_log_id, searched_at)
select
	random_integer(1,4) as search_log_id,
	random_search_timestamp() as searched_at
from
	generate_series(1, SEARCH_COUNT);


insert into
	orders (user_id, address, total, status)
select
	random_integer(1,USER_COUNT) as user_id,
	random_varchar(10) as address,
	0 as total,
	random_order_status() as status
	-- 'pending' as status
from
	generate_series(1, ORDER_COUNT);

insert into
	order_items (order_id, product_id, quantity)
select
	random_integer(1,ORDER_COUNT) as order_id,
	random_integer(1,PRODUCT_COUNT) as product_id,
	random_integer(1,10) as quantity
from
	generate_series(1, ORDER_COUNT * 3)
on 
	conflict (order_id, product_id) do nothing;


update orders set total = (
	select coalesce(sum(oi.quantity * p.price), 0) + 3 as order_sum from order_items oi
	join products p on oi.product_id = p.id
	where oi.order_id = i
) 
from 
	generate_series(1, ORDER_COUNT)  as seq(i)
where 
	orders.id = i;

end $$;
