CREATE TABLE user_statistics (
    user_id               BIGINT PRIMARY KEY,
    first_request_time    TIMESTAMP WITH TIME ZONE,
    total_requests        INTEGER,
    last_request_time     TIMESTAMP WITH TIME ZONE
);