-- Tabela de Convites
CREATE TABLE IF NOT EXISTS invites (
    invite_id SERIAL PRIMARY KEY,
    requester VARCHAR(14), -- school
    school VARCHAR(100) NOT NULL, --name_school
    email_school VARCHAR(100) NOT NULL,
    guest VARCHAR(14), -- driver
    driver VARCHAR(100) NOT NULL, --name_driver
    email_driver VARCHAR(100) NOT NULL,
    status TEXT NOT NULL,
);