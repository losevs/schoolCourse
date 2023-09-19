# Прототип информационной системы для школы

## Примеры запросов:

### Отправление всех запросов осуществляется через программу Postman

## Директор:

#### POST

**Добавление нового директора**

Запрос на адрес localhost:80/dir/add

Тело запроса: 

``` json
{
    "name": "max",
    "surname": "zuev"
}
```

Ответ:

``` json
{
    "message": "new director updated!"
}
```

Запрос (директор уже существует):

Тело запроса: 

``` json
{
    "name": "someone",
    "surname": "xxx"
}
```

Ответ:

``` json
{
    "director": {
        "name": "max",
        "surname": "zuev"
    },
    "message": "director already exists"
}
```

#### GET

**Просмотр существующего директора**

Запрос на адрес localhost:80/dir/show

Ответ:

``` json
{
    "name": "max",
    "surname": "zuev"
}
```

Запрос (директора нет):

Ответ:

``` json
{
    "message": "there is no director"
}
```

#### DELETE

**Удаление существующего директора**

Запрос на адрес localhost:80/dir/del

Ответ:

``` json
{
    "message": "director deleted successfully"
}
```

Запрос (директора нет):

``` json
{
    "message": "there is no director"
}
```

## Предметы:

#### POST

**Добавление нового предмета**

Запрос на адрес localhost:80/subs/add

Тело запроса:

``` json
{
    "subj": "Biology"
}
```

Ответ: 

``` json
{
    "subj": "Biology"
}
```

Запрос (предмет уже существует):

Тело запроса:

``` json
{
    "subj": "Biology"
}
```

Ответ:

``` json
{
    "message": "this subject already exists"
}
```

#### GET

**Просмотр всех доступных предметов**

Запрос на адрес localhost:80/subs/show

Ответ: 

``` json
[
    {
        "subj": "English"
    },
    {
        "subj": "Russian"
    },
    {
        "subj": "Maths"
    },
    {
        "subj": "Biology"
    }
]
```

Запрос (предметов нет):

Ответ:

``` json
{
    "message": "there are no subjects"
}
```

#### DELETE

**Удаление предмета по имени**

Запрос на адрес localhost:80/subs/del

Тело запроса:

``` json
{
    "subj": "Russian"
}
```

Ответ: 

``` json
{
    "message": "subject deleted successfully"
}
```

Запрос (несуществующий предмет):

Тело запроса:

``` json
{
    "subj": "not a subject"
}
```

Ответ:

``` json
{
    "message": "subject not a subject does not exists"
}
```

## Преподаватели:

#### POST

**Добавление нового преподавателя**

Запрос на адрес localhost:80/teach/new

Тело запроса:

``` json
{
    "id": 1,
    "name": "Valery",
    "surname": "Treyster",
    "grade": 9,
    "subject": "Biology"
}
```

Ответ:

``` json
{
    "id": 1,
    "name": "Valery",
    "surname": "Treyster",
    "grade": 9,
    "subject": "Biology"
}
```

Запрос (добавление пользователя с уже существующим ID):

``` json
{
    "id": 1,
    "name": "x",
    "surname": "x",
    "grade": 0,
    "subject": "x"
}
```

Ответ:

``` json
{
    "message": "teacher with this ID already exists",
    "teacher": {
        "id": 1,
        "name": "Valery",
        "surname": "Treyster",
        "grade": 9,
        "subject": "Biology"
    }
}
```

Запрос (добавление пользователя с недействительным названием предмета):

``` json
{
    "id": 10,
    "name": "x",
    "surname": "x",
    "grade": 0,
    "subject": "subject_False"
}
```

Ответ:

``` json
{
    "message": "subject subject_False does not exists in our school"
}
```

#### GET

**Просмотр всех существующих преподавателей**

Запрос на адрес localhost:80/teach/show

Ответ:

``` json
[
    {
        "id": 0,
        "name": "Alexey",
        "surname": "Korot",
        "grade": 7,
        "subject": "Russian"
    },
    {
        "id": 1,
        "name": "Valery",
        "surname": "Treyster",
        "grade": 9,
        "subject": "Biology"
    },
    {
        "id": 2,
        "name": "someone",
        "surname": "noone",
        "grade": 11,
        "subject": "Russian"
    }
]
```

**Просмотр преподавателя по ID**

Запрос на адрес localhost:80/teach/show/exact/1

Ответ:

``` json
{
    "id": 1,
    "name": "Valery",
    "surname": "Treyster",
    "grade": 9,
    "subject": "Biology"
}
```

Запрос (несуществующее ID):

Ответ:

``` json
{
    "message": "teacher with id=999 does not exist"
}
```

**Просмотр преподавателей по предмету**

Запрос на адрес localhost:80/teach/show/sub

Тело запроса:

``` json
{
    "subj": "Russian"
}
```

Ответ:

``` json
[
    {
        "id": 0,
        "name": "Alexey",
        "surname": "Korot",
        "grade": 7,
        "subject": "Russian"
    },
    {
        "id": 2,
        "name": "someone",
        "surname": "noone",
        "grade": 11,
        "subject": "Russian"
    }
]
```

Запрос (несуществующий предмет / предмет, который никто не ведет):

Тело запроса:

``` json
{
    "subj": "subj"
}
```

Ответ:

``` json
{
    "message": "there is noone studying subj"
}
```

#### DELETE

**Удаление преподавателя по ID**

Запрос на адрес localhost:80/teach/delete/1

Ответ:

``` json
{
    "message": "teacher deleted successfully"
}
```

Запрос Запрос на адрес localhost:80/teach/delete/999 (несуществующий ID):

``` json
{
    "message": "there is noone with id=999"
}
```

#### PATCH

**Обновление преподавателя по ID**

Запрос на адрес localhost:80/teach/update/0

*При отсутствии явного указания полей будут оставаться старые*

*id указывать не нужно, в примере указан для наглядности*

Тело запроса:

``` json
{
        "id": 100,
        "name": "name",
        "surname": "surname",
        "grade": 1,
        "subject": "Maths"
    }
```

Ответ:

``` json
{
    "id": 0,
    "name": "name",
    "surname": "surname",
    "grade": 1,
    "subject": "Maths"
}
```

Запрос на адрес localhost:80/teach/update/999 (несуществующий ID):

Тело запроса:

``` json
{
        "id": 100,
        "name": "name",
        "surname": "surname",
        "grade": 1,
        "subject": "Maths"
    }
```

Ответ:

``` json
{
    "message": "there is noone with id=999"
}
```

Запрос на адрес localhost:80/teach/update/0 (несуществующий предмет):

Тело запроса:

``` json
{
        "subject": "false_subject"
}
```

Ответ:

``` json
{
    "message": "subject false_subject does not exists in our school",
    "new teacher": {
        "id": 0,
        "name": "name",
        "surname": "surname",
        "grade": 1,
        "subject": "Maths"
    }
}
```

## Обучающиеся:

#### POST

**Добавление нового пользователя**

Запрос на адрес localhost:80/new

Тело запроса:

``` json
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Ответ:

``` json
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Запрос (добавление пользователя с уже существующим ID):

``` json
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Ответ:

``` json
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

``` json
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

``` json
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

``` json
{
    "message": "Class 111 is empty"
}
```

**Просмотр выбранного пользователя по ID**

Запрос на адрес localhost:80/show/0

Ответ:

``` json
{
    "id": 0,
    "name": "Losev",
    "surname": "Sergey",
    "age": 15,
    "grade": 9
}
```

Запрос на адрес localhost:80/show/111 (несуществующий ID)

``` json
{
    "message": "there is no kid with that id"
}
```

#### DELETE

**Удаление пользователя по ID**

Запрос на адрес localhost:80/delete/0

Ответ:

``` json
{
    "message": "kid deleted succsessfully!"
}
```

Запрос на адрес localhost:80/delete/111

Ответ:

``` json
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

``` json
{
    "id": 200,
    "name": "Sergey",
    "surname": "Losev",
    "age": 21,
    "grade": 11
}
```

Ответ:

``` json
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

``` json
{
    "id": 200,
    "name": "Sergey",
    "surname": "Losev",
    "age": 21,
    "grade": 11
}
```

Ответ:

``` json
{
    "message": "there is no kid with this ID"
}
```