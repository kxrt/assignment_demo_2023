\c postgres;
CREATE TABLE messages (
    id         VARCHAR(100) NOT NULL,
    message    VARCHAR(500) NOT NULL,
    sender     VARCHAR(50) NOT NULL,
    created_at BIGINT NOT NULL,
    PRIMARY KEY (id, message, sender, created_at)
);