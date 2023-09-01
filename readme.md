# Avito segment task
Решение тестового задания avito go backend-trainee-assignment-2023
# Инструкция по запуску
```bash
docker-compose up --build
```
# API
Данный микросервис имеет четыре метода. Примеры запросов можно найти в файле Avito-segment-test-task.postman_collection
## Описание методов в swaggert
### http://localhost:8080/swagger/index.html
## Метод добавления пользователя в базу
### Post "/user/"
Данный метод принимает id(userId) пользователя и добавляет в базу имеющихся пользователей. Ответом метода является уведомление об успешности операции или ошибка.
### Пример
```bash
curl --request POST \
  --url http://localhost:8080/user/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: Insomnia/2023.5.7' \
  --data '{
  "userId": 31
}'
```
## Метод добавления пользователю сегментов
### Patch "/user/"
Данный метод принимает три параметра: список имён сегментов на добавление(segToAdd), список имён сегментов на удаление(segTosegToDelete) и id пользователя для которого совершаются операции. Ответом метода является уведомление об успешности операции или ошибка.
### Пример
```bash
curl --request PATCH \
  --url http://localhost:8080/user/ \
  --header 'Content-Type: application/json' \
  --data '{
  "segToAdd": [
    "sale20"
  ],
  "segToDelete": [
    "sale10"
  ],
  "userId": 31
}'
```
## Метод получения сегментов пользователя
### Get "/user/"
Данный метод принимает id(userId) пользователя. Ответом метода является список имён сегментов которым принадлежит пользователь.
### Пример
```bash
curl --request GET \
  --url http://localhost:8080/user/ \
  --header 'Content-Type: application/json' \
  --data '{
	"userId": 31
}'
```
## Метод создания сегмента
### Post "/segment/"
Данный метод принимает slug(segName) сегмента. Ответом метода является id добавленного в базу сегмента.
### Пример
```bash
curl --request POST \
  --url http://localhost:8080/segment \
  --header 'Content-Type: application/json' \
  --data '{
	"segName": "sale10"
}'
```
## Метод удаления сегмента
Данный метод принимает slug(segName) сегмента. Ответом метода является сообщение о успешности удаления.
### Delete "/segment/"
### Пример
```bash
curl --request DELETE \
  --url http://localhost:8080/segment \
  --header 'Content-Type: application/json' \
  --data '{
	"segName": "sale10"
}'
```



