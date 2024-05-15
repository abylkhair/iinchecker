# Kaspi REST
## Тестовое задание
Нужно реализовать Web-сервис, со следующими роутами:

1. Проверка ИИН
   Нужно проверить контрольную сумму ИИН, определить пол человека и дату рождения. Вся инфа есть в википедии.
```
/iin_check/{номер_ИИН} - метод GET, ответ в виде JSON:
   {
   "correct": true/false,
   "sex": "female"/"male",
   "date_of_birth": "DD.MM.YYYY"
   }
```

2. Заполнение БД
   Реализовать запись в БД информации о человеке. БД можно использовать любую, то есть даже JSON-файл.
```
/people/info - метод POST, в теле запроса JSON:
{
"name": "Мыркымбаев Мыркымбай Мыркымбайулы",
"iin": "123123123123",
"phone": "+71231231234"
}
```
   ответ на запрос в виде JSON:
```
{
"success": true/false,
"errors": "ошибки если есть"
}
```
   Проверять на корректность ИИН-а функцией из первого пункта. 
   При ошибке возвращать код 500, "success" поле false, и ошибки в поле "errors".

3. Получении инфо из БД.
3.1. Получение ранее сохраненных данных о человеке по ИИН.
```
/people/info/iin/{iin} - метод GET, ответ на запрос в виде JSON (структура данных):
{
"name": "Мыркымбаев Мыркымбай Мыркымбайулы",
"iin": "123123123123",
"phone": "+71231231234"
}
```
Проверять на корректность ИИН-а функцией из первого пункта. При ошибке возвращать код 500, "success" поле false, и ошибки в поле "errors". 
Возвращать код 404, если не найдено.
3.2. Получение ранее сохраненных данных о человеке по части имени:
```
/people/info/phone/{часть_имени_или_фамилии} - метод GET, ответ на запрос в виде JSON (массив структур данных):
[
{
"name": "Мыркымбаев Мыркымбай Мыркымбайулы",
"iin": "123123123123",
"phone": "+71231231234"
},
{
"name": "Мыркымбаев Кырыкбай Сайлаубекулы",
"iin": "123123123123",
"phone": "+71231231234"
}
]
```
Возвращать пустой массив ("[]") если данные не найдены.