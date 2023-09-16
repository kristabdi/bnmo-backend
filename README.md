# bnmo-backend
## Tech Stack
```
Golang 1.18
PostgreSQL 14
Redis 7.0
```

## How-to-Run
You can build the server side by `docker compose up` in CLI.\
Endpoint : http://localhost:3001

## Endpoint
```
[POST] /auth/register
- username
- name
- password
- photo

[POST] /auth/login
- username
- password

[POST] /admin/verify/user?username=xxx

[POST] /admin/verify/req?id=xxx

[GET] /admin/list/user

[GET] /admin/list/history

[GET] /customer/info

[GET] /customer/history/request?page=1&page_size=5 OR /customer/history/transaction?page=1&page_size=5

[POST] /customer/withdraw
- amount
- currency

[POST] /customer/deposit
- amount
- currency

[POST] /customer/transaction
- amount
- username_to
- currency_from
```

## Design Pattern
```
1. Singleton
Single object databases are used continuously so they are suitable for using the singleton design pattern. Initialization of the client database is carried out at the beginning, and is used continuously.
2. Composite
Database models are in the form of composites, such as RateInfo and QueryInfo which are found in the Converter and User models which are found in several other models such as Transaction and Request.
```
