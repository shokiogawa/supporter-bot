
-- +migrate Up
CREATE TABLE IF NOT EXISTS costs(
    id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    public_cost_id VARCHAR(1024) NOT NULL,
    user_id MEDIUMINT NOT NULL ,
    title VARCHAR(64) NOT NULL ,
    outcome INT NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
    );
ALTER TABLE costs ADD INDEX index2(user_id, created_at);
-- +migrate Down

DROP TABLE IF EXISTS costs;