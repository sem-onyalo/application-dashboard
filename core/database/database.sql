create table endpoint (
  id serial primary key,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null,
  deleted_at timestamp,
  name varchar(127) not null,
  url varchar(255) not null
);

create table endpoint_test (
  id serial primary key,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null,
  deleted_at timestamp,
  endpoint_id int references endpoint(id) on delete cascade,
  response_status varchar(16) not null,
  time_elapsed float not null
);
