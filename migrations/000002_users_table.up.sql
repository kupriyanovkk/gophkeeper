create TABLE users (
    id UUID DEFAULT gen_random_uuid() not null unique,
    login TEXT not null unique,
    password TEXT not null
);

create index index_login_users on users (login);