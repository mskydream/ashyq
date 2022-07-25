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

-- ALTER TABLE user_profile    OWNER TO postgres;
-- ALTER TABLE real_estate     OWNER TO postgres;
-- ALTER TABLE visit           OWNER TO postgres;
-- COMMIT;



ALTER TABLE user_profile    OWNER TO study;
ALTER TABLE real_estate     OWNER TO study;
ALTER TABLE visit           OWNER TO study;
COMMIT;

-- INSERT INTO user_profile(username,surname,born, status,phone_number,iin,gender,residential_address,password,created_at)
-- VALUES 
-- ('Aidyn','Omarov','2000-11-29 22:00:44','green',87474224540,001129500582,'Male','per, Lokomotiv 25','Aidyn','2022-07-07 22:00:44'),

-- --  create users
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Hedvige', 'Moffat', '2015-04-15', 'green', '4382945196', '62', 'Female', '32802 Mariners Cove Place', 'a6MaJhzj06', '2022-05-16');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Ileana', 'Mease', '2014-06-09', 'green', '1837968708', '6780', 'Male', '8596 Browning Park', 'psG51APsz', '2022-06-22');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Merrill', 'Greensite', '2011-12-08', 'green', '8175941012', '01022', 'Male', '7 Spaight Trail', 'mhPD7lEJ6V', '2021-09-03');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Micky', 'Hauger', '2004-10-04', 'green', '4391714623', '5', 'Male', '34 Declaration Place', 'prXVPjXc', '2021-11-17');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Lissa', 'Sturgess', '2004-06-07', 'red', '4422677689', '3759', 'Female', '5126 La Follette Alley', 'Q0NwTEyUqQY', '2022-03-14');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Paige', 'Pllu', '2003-09-23', 'green', '8157069163', '37691', 'Male', '13876 Hovde Trail', 'iMIEkD0GA', '2022-04-23');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Hillyer', 'Vatcher', '2005-10-25', 'green', '8614292550', '8814', 'Male', '40 Arkansas Terrace', 'dJaoc9rSN8', '2022-05-21');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Kellsie', 'Bensusan', '2003-12-19', 'green', '3062763148', '3106', 'Female', '4085 Arapahoe Court', '6VJxh500', '2021-09-27');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Gina', 'Plaid', '2013-08-11', 'yellow', '9389308947', '36326', 'Female', '43024 Hooker Place', 'E7dSnYCEal', '2022-06-29');
-- insert into user_profile (username, surname, born, status, phone_number, iin, gender, residential_address, password, created_at) values ('Mallory', 'Paske', '2009-02-03', 'yellow', '7778998170', '99807', 'Female', '97916 School Point', 'JTFWOh0v', '2021-09-17');


-- -- create real-estate
-- insert into real_estate(user_profile_id, address, qr_code,created_at) 
-- values(1,'улица Жибек Жолы 135, Almaty 050000','ashyq.kz zhibek zholy 135',now())

-- -- create 

-- insert into visit(real_estate_id,user_profile_id,created_at)
-- values (1,1,now())