-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS bookings
(
    id           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    course       VARCHAR(255) NOT NULL,
    class        VARCHAR(255) NOT NULL,
    lab          VARCHAR(255) NOT NULL,
    teacher      VARCHAR(255) NOT NULL,
    booking_date DATE         NOT NULL,
    time_start   TIME         NOT NULL,
    time_end     TIME         NOT NULL,
    description  TEXT         NULL,
    created_at   TIMESTAMP    NOT NULL,
    updated_at   TIMESTAMP    NULL,
    deleted_at   TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS bookings;