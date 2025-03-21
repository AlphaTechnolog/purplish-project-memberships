CREATE TABLE IF NOT EXISTS memberships (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    scopes VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS company_memberships (
    company_id VARCHAR(36) NOT NULL,
    membership_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (company_id, membership_id)
);