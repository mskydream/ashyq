BEGIN;
CREATE TABLE IF NOT EXISTS user_profile(
    id                      serial                                              PRIMARY KEY,
    name                    VARCHAR(255)                                        NOT NULL,
    surname                 VARCHAR(255)                                        NOT NULL,
    born                    TIMESTAMP                                           NOT NULL,
    status                  TEXT CHECK(status IN ('green','red', 'yellow'))     NOT NULL,
    phone_number            VARCHAR(255) UNIQUE                                 NOT NULL,
    iin                     VARCHAR(255) UNIQUE                                 NOT NULL,
    gender                  TEXT CHECK(gender IN ('Male','Female'))             NOT NULL,
    residential_address     TEXT                                                NOT NULL,
    password                TEXT                                                NOT NULL,
    created_at              TIMESTAMP                                           NOT NULL
);

CREATE TABLE IF NOT EXISTS real_estate(
    id                      serial                  PRIMARY KEY,
    user_profile_id         INT                     NOT NULL REFERENCES user_profile(id),
    address                 TEXT UNIQUE             NOT NULL,
    qr_code                 TEXT UNIQUE             NOT NULL,
    created_at              TIMESTAMP               NOT NULL
);

CREATE TABLE IF NOT EXISTS visit(
    id                      serial                  PRIMARY KEY,
    real_estate_id          INT NOT NULL            REFERENCES real_estate(id),
    user_profile_id         INT NOT NULL            REFERENCES user_profile(id),
    created_at              TIMESTAMP               NOT NULL
);


ALTER TABLE user_profile    OWNER TO study;
ALTER TABLE real_estate     OWNER TO study;
ALTER TABLE visit           OWNER TO study;
COMMIT;