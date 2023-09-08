# Прототип информационной системы для школы

## Примеры запросов:

### Отправление всех запросов осуществляется через программу Postman

#### POST

**Добавление нового пользователя**

Запрос на адрес localhost:80/new

Тело запроса:

```
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Ответ:

```
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Запрос (добавление пользователя с уже существующим ID):

```
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Ответ:

```
{
    "kid": {
        "id": 0,
        "name": "Losev",
        "surname": "Sergey",
        "age": 15,
        "grade": 9
    },
    "message": "this ID already exists"
}
```

#### GET

**Просмотр всех существующих пользователей**

Запрос на адрес localhost:80/show

Ответ:

```
[
    {
        "id": 0,
        "name": "Losev",
        "surname": "Sergey",
        "age": 15,
        "grade": 9
    },
    {
        "id": 1,
        "name": "Roman",
        "surname": "Mamoev",
        "age": 16,
        "grade": 9
    },
    {
        "id": 2,
        "name": "Andrey",
        "surname": "Domnin",
        "age": 16,
        "grade": 10
    },
    {
        "id": 3,
        "name": "Egor",
        "surname": "Skorn",
        "age": 17,
        "grade": 11
    },
    {
        "id": 4,
        "name": "Alexey",
        "surname": "Kurnikov",
        "age": 15,
        "grade": 9
    },
    {
        "id": 5,
        "name": "Ivan",
        "surname": "Anikin",
        "age": 7,
        "grade": 1
    },
    {
        "id": 6,
        "name": "Matvey",
        "surname": "Alexeenko",
        "age": 11,
        "grade": 5
    },
    {
        "id": 7,
        "name": "Alina",
        "surname": "Kozlova",
        "age": 8,
        "grade": 1
    },
    {
        "id": 8,
        "name": "Andrey",
        "surname": "Ivanov",
        "age": 11,
        "grade": 5
    },
    {
        "id": 9,
        "name": "Max",
        "surname": "Zuev",
        "age": 10,
        "grade": 4
    },
    {
        "id": 10,
        "name": "Denis",
        "surname": "Danchishin",
        "age": 11,
        "grade": 5
    }
]
```

**Просмотр всех пользователей из выбранного класса**

Запрос на адрес localhost:80/show/class/5

Ответ:

```
[
    {
        "id": 6,
        "name": "Matvey",
        "surname": "Alexeenko",
        "age": 11,
        "grade": 5
    },
    {
        "id": 8,
        "name": "Andrey",
        "surname": "Ivanov",
        "age": 11,
        "grade": 5
    },
    {
        "id": 10,
        "name": "Denis",
        "surname": "Danchishin",
        "age": 11,
        "grade": 5
    }
]
```

Запрос на адрес localhost:80/show/class/111 (несуществующий класс)

Ответ:

```
{
    "message": "Class 111 is empty"
}
```

**Просмотр выбранного пользователя по ID**

Запрос на адрес localhost:80/show/0

Ответ:

```
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Запрос на адрес localhost:80/show/111 (несуществующий ID)

```
{
    "message": "there is no kid with that id"
}
```

#### DELETE

**Удаление пользователя по ID**

Запрос на адрес localhost:80/delete/0

Ответ:

```
{
    "message": "kid deleted succsessfully!"
}
```

Запрос на адрес localhost:80/delete/111

Ответ:

```
{
    "message": "there is no kids with this id"
}
```

#### PATCH

**Обновление пользователя по ID**

Запрос на адрес localhost:80/update/0

*При отсутствии явного указания полей будут оставаться старые*

*id указывать не нужно, в примере указан для наглядности*

Запрос:

```
{
    "id": 200,
    "name": "Sergey",
    "surname": "Losev",
    "age": 21,
    "grade": 11
}
```

Ответ:

```
{
    "id": 0,
    "name": "Sergey",
    "surname": "Losev",
    "age": 21,
    "grade": 11
}
```

Запрос на адрес localhost:80/update/111

Запрос:

```
{
    "id": 200,
    "name": "Sergey",
    "surname": "Losev",
    "age": 21,
    "grade": 11
}
```

Ответ:

```
{
    "message": "there is no kid with this ID"
}
```