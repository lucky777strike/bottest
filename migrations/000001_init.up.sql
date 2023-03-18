CREATE TABLE user_statistics (
    user_id               BIGINT PRIMARY KEY,
    first_request_time    TIMESTAMP WITH TIME ZONE,
    total_requests        INTEGER,
    last_request_time     TIMESTAMP WITH TIME ZONE
);
CREATE TABLE weather (
    city      VARCHAR(255) NOT NULL,
    temp      INTEGER NOT NULL,
    condition VARCHAR(255) NOT NULL,
    last_upd  TIMESTAMP NOT NULL,

    PRIMARY KEY (city)
);

CREATE TABLE currency (
    name TEXT PRIMARY KEY,
    value FLOAT NOT NULL,
    last_updated  TIMESTAMP NOT NULL
);