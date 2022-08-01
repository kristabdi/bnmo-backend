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
Database single object digunakan terus menerus sehingga cocok menggunakan design pattern singleton. Inisialisasi client database dilakukan pada awal, dan digunakan terus menerus.
2. Composite
Models database berbentuk komposit seperti misalnya RateInfo dan QueryInfo yang terdapat pada model Converter dan User yang terdapat pada beberapa model lain seperti Transaction dan Request.
```