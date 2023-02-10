DROP TABLE IF EXISTS Users

DROP TABLE IF EXISTS Session 

DROP TABLE IF EXISTS GitHubResponse

DROP TABLE IF EXISTS GitHubUserData

DROP TABLE IF EXISTS GoogleContentUser


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
    UUID BIGINT NOT NULL UNIQUE,
    REFRESH_TOKEN VARCHAR(245) NOT NULL,
    ExpiresRefresh BIGINT NOT NULL,
)


CREATE TABLE GitHubResponse(
    AcessToken VARCHAR(244) NOT NULL,
    TokenType VARCHAR(244) NOT NULL, 
    Scope VARCHAR(244) NOT NULL,
)


CREATE TABLE GitHubUserData(
     ID SERIAL PRIMARY KEY,
    Name VARCHAR(244) NOT NULL UNIQUE,
    Login VARCHAR(244) NOT NULL UNIQUE,
    Location VARCHAR(244) NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTMAP,
)

CREATE TABLE GoogleContentUser(
    ID VARCHAR(244) NOT NULL,
    Email VARCHAR(244) NOT NULL UNIQUE,
)

