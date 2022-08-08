
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    public_user_id VARCHAR(1024) NOT NULL,
    line_user_id VARCHAR(64) NOT NULL ,
    image VARCHAR (1024) NOT NULL ,
    name VARCHAR(32) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down

DROP TABLE IF EXISTS users;
