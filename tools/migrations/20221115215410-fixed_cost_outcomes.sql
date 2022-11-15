
-- +migrate Up
CREATE TABLE IF NOT EXISTS fixed_cost_outocomes(
    id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    public_fixed_cost_outcome_id VARCHAR(1024) NOT NULL,
    fixed_cost_id MEDIUMINT NOT NULL ,
    outcome INT NOT NULL ,
    outocome_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(fixed_cost_id) REFERENCES fixed_costs(id)
    );

-- +migrate Down
DROP TABLE IF EXISTS fixed_cost_outocomes;