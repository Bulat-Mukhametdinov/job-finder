create table if not exists users (
    id integer primary key autoincrement,
    username varchar(50) unique not null,
    password_hash char(128) not null,
    created_at timestamp default current_timestamp
);

create table if not exists favourites (
    id text primary key autoincrement,
    comments text,
    created_at timestamp default current_timestamp,
    user_id integer not null,
    foreign key (user_id) references users(id) on delete cascade
);

create table if not exists sessions (
    token char(16) primary key,
    expires_at timestamp not null,
    created_at timestamp default current_timestamp,
    user_id integer not null,
    foreign key (user_id) references users(id) on delete cascade
);