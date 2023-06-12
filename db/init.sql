\c postgres;
CREATE TABLE messages (
    key        SERIAL PRIMARY KEY,
    chat       VARCHAR(100) NOT NULL,
    message    VARCHAR(500) NOT NULL,
    sender     VARCHAR(50) NOT NULL,
    created_at BIGINT NOT NULL
);