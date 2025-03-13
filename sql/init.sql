\c warehouse

create table users (
	id serial primary key,
	email varchar unique not null,
    password varchar not null,
    role varchar not null check (role in ('admin', 'staff', 'guest')),
	created_at timestamp not null default (now()),
  	updated_at timestamp not null default (now()),
   	-- deleted_at timestamp
);

create table warehouse_locations (
    id serial primary key,
    name varchar not null,
    capacity integer not null check (capacity >= 0),
	created_at timestamp not null default (now()),
  	updated_at timestamp not null default (now()),
   	-- deleted_at timestamp
);

create table products (
    id serial primary key,
    name varchar not null,
    sku varchar unique not null,
    quantity integer not null check (quantity >= 0),
    location_id integer not null references warehouse_locations(id),
	created_at timestamp not null default (now()),
  	updated_at timestamp not null default (now()),
   	-- deleted_at timestamp
);

CREATE TABLE orders (
    id serial primary key,
    product_id integer not null references products(id),
    quantity integer not null check (quantity > 0),
    type varchar not null check (type in ('receive', 'ship')),
	created_at timestamp not null default (now()),
  	updated_at timestamp not null default (now()),
   	-- deleted_at timestamp
);