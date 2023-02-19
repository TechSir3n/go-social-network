# Small Social Network 


# [Images (must-view)](https://github.com/Kl1ck9r/go-social-network/tree/main/screenshots)

---
# Requirements
* **Go:** `1.19`
* **PostgreSQL:**  `15.2`
* **Memcached:** `1.6.17`
* **Redis:** `6.0.17`
* **Google** and **Github** keys

## docker-compose
Server is ready immediately after containers are up
```shell
 make compose        #docker-compose up
```

## Dockerfile 
```shell
make docker-build   #docker build -t social-network .
make docker-run     #docker run social-network
```

---

## Setup OAuth

### - Github
1. Go to the **[Developer settings](https://github.com/settings/apps)**
2. Create **[Application](https://github.com/settings/apps)**
3. Enable `User permissions` -> `Email addresses` -> `Read Only` in the **[Permissions](https://github.com/settings/apps/permissions)**
4. Generate **secret token**
5. Enter the `URIs` that are allowed to be redirect-URIs (e.g. `https://localhost:3000/oauth/github/callback`)
6. Paste both `Client ID` and `Client Secret` to the `github_secret.json`

### - Google
1. Go to the **[Google Cloud Console](https://console.cloud.google.com/projectselector2/apis/credentials)**
2. Create project (add content to the consent screen like title and logo) or use existing
3. `Credentials` -> `Create Credentials` -> `Oauth Client ID`
4. Choose the `Web Application` type and give it a name
5. Enter the `URIs` that are allowed to be redirect-URIs (e.g. `https://localhost:3000/oauth/google/callback`)
6. Paste both `Client ID` and `Client Secret` to the `google_secret.json`

---

- ### Redis
    **Port** `:6379` **Password** `password` **Database** `0`
- ### PostgreSQL
    **Port** `:5431` **User** `postgres` **Password** `pg_password` **Database** `social-network`

It is possible to additionally configure the app using environment variables
```yaml
environment:
  POSTGRES_IP: 127.0.0.1 # connect to local database
  HOST_PORT: 8082 # change server port
```
---
## Build / Run

```shell
git clone https://github.com/Kl1ck9r/go-social-network.git 
cd social-network
make build
make run 
```
## Run arguments 
```shell
  server -http # or https
```

---
## Setup Postgresql 
```shell
migrate -database ${POSTGRESQL_URL} -path migrate/ up
```

### Down 
```shell
migrate -database ${POSTGRESQL_URL} -path migrate/ down
```
---

## Tests
* [x] Mocks
* [x] Unit 

```shell
make test 
```

---

## Logs
```yaml
# File (default logger.txt)
[ERROR]: 16002-216-216 216:216:216 - [memcache: cache miss Failed to get value of the memcached]
[INFO]: 12002-212-212 212:212:212 - [database connected successfully]
[INFO]: 16002-216-216 216:216:216 - protocol[http method GET path /favicon.ico Duration  769.2µs Status Code  200 Status text OK]
[INFO]: 16002-216-216 216:216:216 - protocol[http method GET path /login Duration  512.6µs Status Code  200 Status text OK]
[FATAL]: 16002-216-216 216:216:216 - [listen tcp :3000: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.]
```


## Endpoints

### - Main Page `/home`

### - Register  `/registration`

### - Login `/login`

### - Logout `/logout`

### - Mail verify `/verify`

### - Reset Password `/reset/password`

### - Access Admin `/access/admin`

#### - Github OAuth `/login/github`  `/github/callback` 

#### - Google OAuth `/login/google`  `/google/callback` 