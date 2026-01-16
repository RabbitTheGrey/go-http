## Инструкция по развертыванию ##
1. Создать и заполнить файл `.env`, пример в `.env.dist`
```
touch .env
```
2. создать в корне проекта файл sqlite.db
3. сборка и запуск проекта 
```
sudo docker compose build && sudo docker compose up -d
```
4. для пересборки использовать `./cmd/sh/rebuild.sh`
5. выполнить миграции 
```
sudo docker exec -it go-http-app01 migrate
```
6. Создать пользователя через админ панель
```
sudo docker exec -it go-http-app01 admin-cli
```
7. Сервис доступен на localhost(или ваш ip):8080(или другой порт из .env)

По обратной связи пишите, пожалуйста, в телеграм @rabbitthegrey.

Это поможет будущим проектам стать еще лучше
