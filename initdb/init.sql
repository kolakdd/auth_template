CREATE TABLE users (     
    guid UUID PRIMARY KEY,	
    name VARCHAR(16) NOT NULL, 
    deactivated BOOLEAN NOT NULL,
    last_login_at TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE refresh_tokens (     
    token_hash TEXT PRIMARY KEY,
    user_guid UUID NOT NULL REFERENCES users(guid),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE invalid_access_tokens (
    guid UUID PRIMARY KEY,
    user_guid UUID NOT NULL REFERENCES users(guid),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);