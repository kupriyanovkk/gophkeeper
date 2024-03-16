CREATE TABLE private (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    type INT NOT NULL,
    content BYTEA NOT NULL,
    updated TIMESTAMP DEFAULT NOW(),
    deleted BOOLEAN DEFAULT FALSE,
    constraint fk_user_id foreign key (user_id) references users (id) on delete cascade
);