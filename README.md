# Тестовое задание avitoTech

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [Endpoints](#Endpoints)
1. [Запуск](#Запуск)
1. [Примеры](#Примеры)
<!-- ToC end -->

# Описание задачи

Разработать микросервис для работы с балансом пользователей (баланс, зачисление/списание/перевод средств). 
Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.
Дополнительно реализовать методы конвертации баланса, приобретения услуг и получения списка транзакций.
Полное описание в [TASK](https://github.com/avito-tech/job-backend-trainee-assignment/).
# Реализация

- Следование дизайну REST API.
- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с фреймворком [gin-gonic/gin](https://github.com/gin-gonic/gin).
- Работа с СУБД Postgres с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Запуск бд из Docker.
**Структура проекта:**
```
.
├── pkg
│   ├── handler     // обработчики запросов
│   ├── service     // бизнес-логика
│   └── repository  // взаимодействие с БД
├── cmd             // точка входа в приложение
├── schema          // SQL файлы с миграциями
├── configs         // файлы конфигурации
```

# Endpoints

- GET /api/balance/id - получение баланса пользователя по id
- GET /api/usdBalance/id - получение баланса пользователя в долларах США по id
- PUT /api/balance/id - пополнение баланса/списание из баланса пользователя id
    - Тело запроса:
        - netto - сумма перевода.
        - cashflow - направление денежного потока.
- GET /api/transactions/id - получение списка транзакций пользователя, отсортированного по времени транзакции по id
- GET /api/trans_by_summ/id - получение списка транзакций пользователя, отсортированного по убыванию суммы по id
- PUT /api/send/id - перевод средств на баланс другого пользователя со счета id
    - Тело запроса:
        - netto - сумма перевода в RUB.
        - reacherid - идентификатор пользователя, на баланс которого начисляются средства.
- PUT /api/purchase/id - приобретение одной из предоставляемых услуг пользователем id.
    - Тело запроса:
        - service - приобретаемая услуга.
# Запуск

```
go run cmd/main.go
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
docker pull postgres
docker run --name=balance-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
go mod tidy
```


# Примеры

Запросы сгенерированы из Postman.

### 1. GET  /balance/1

**Тело ответа:**
```
{
    "balance": 1000
}
```

### 2. GET /usdBalance/1
**Запрос:**

**Тело ответа:**
```
{
    "usdBalance": 14.28
}
```

### 3. GET /trans_by_summ/2


**Тело ответа:**
```
{
    "transactions": [
        {
            "id": 13,
            "user_id": 2,
            "netto": 1000,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:49:11.362805Z"
        },
        {
            "id": 2,
            "user_id": 2,
            "netto": 400,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:11:05.400642Z"
        },
        {
            "id": 1,
            "user_id": 2,
            "netto": 400,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:08:54.206476Z"
        },
        {
            "id": 4,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:13:18.163215Z"
        },
        {
            "id": 5,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:14:05.728221Z"
        },
        {
            "id": 6,
            "user_id": 2,
            "netto": 100,
            "cashflow": false,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:19:23.398114Z"
        },
        {
            "id": 7,
            "user_id": 2,
            "netto": 100,
            "cashflow": false,
            "source_or_purpose": "массаж",
            "transtime": "2023-01-11T20:20:53.601119Z"
        },
        {
            "id": 3,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:11:22.814428Z"
        },
        {
            "id": 9,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:23:17.591597Z"
        },
        {
            "id": 11,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:23:45.77365Z"
        },
        {
            "id": 8,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:21:03.874242Z"
        }
    ]
}
```

### 4. GET /transactions/2

**Тело ответа:**
```
{
    "transactions": [
        {
            "id": 1,
            "user_id": 2,
            "netto": 400,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:08:54.206476Z"
        },
        {
            "id": 2,
            "user_id": 2,
            "netto": 400,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:11:05.400642Z"
        },
        {
            "id": 3,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:11:22.814428Z"
        },
        {
            "id": 4,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:13:18.163215Z"
        },
        {
            "id": 5,
            "user_id": 2,
            "netto": 100,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:14:05.728221Z"
        },
        {
            "id": 6,
            "user_id": 2,
            "netto": 100,
            "cashflow": false,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:19:23.398114Z"
        },
        {
            "id": 7,
            "user_id": 2,
            "netto": 100,
            "cashflow": false,
            "source_or_purpose": "массаж",
            "transtime": "2023-01-11T20:20:53.601119Z"
        },
        {
            "id": 8,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:21:03.874242Z"
        },
        {
            "id": 9,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:23:17.591597Z"
        },
        {
            "id": 11,
            "user_id": 2,
            "netto": 20,
            "cashflow": false,
            "source_or_purpose": "1",
            "transtime": "2023-01-11T20:23:45.77365Z"
        },
        {
            "id": 13,
            "user_id": 2,
            "netto": 1000,
            "cashflow": true,
            "source_or_purpose": "developer",
            "transtime": "2023-01-11T20:49:11.362805Z"
        }
    ]
}
```

### 5. PUT /balance/1

**Тело запроса:**
```
{
    "netto" : 1000,
    "cashflow" : true
}
```
**Тело ответа:**
```
{
    "status": ok
}
```

### 6. PUT /send/2

**Тело запроса:**
```
{
    "netto" : 20,
    "reacherid" : 1
}
```
**Тело ответа:**
```
{
    "status": ok
}
```

### 7. PUT /purchase/2

**Тело запроса:**
```
{
    "services" : "массаж"
}
```
**Тело ответа:**
```
{
    "status": ok
}
```
