DROP TABLE IF EXISTS Users

DROP TABLE IF EXISTS Session 

CREATE TABLE Users(
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(249) NOT NULL UNIQUE,
    Email VARCHAR(249) NOT NULL UNIQUE,
    Password BIGINT NOT NULL,
    Created_at TIMESTAMP,
    Update_at TIMESTAMP ,
)


CREATE TABLE Session(
    userID SERIAL PRIMARY KEY,
    REFRESH_TOKEN BIGINT NOT NULL UNIQUE,
    ExpiresIn INT NOT NULL,
    FingerPrint varchar(244) NOT NULL,
)


