# PR-TEST-3

## Что это

Выполнение тестового задания:
1. Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
2. Реализовать gRPC сервис на основе proto файла на Go
3. Для хранения данных использовать PostgreSQL
4. на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редиса
5. При добавлении пользователя делать лог в clickHouse
6. Добавление логов в clickHouse делать через очередь Kafka


## Структура папок

- cmd
  - **pr-test-3:** сервер (proto файл в папке users/delivery)
  - **pr-test-3-client:** тестовый клиент (proto файл в папке proto)
- init
  - **clickhouse.sql:** команды создания базы в clickhouse и настройки автоматической выгрузки из kafka
- scripts: пара скриптов в помощь
- third_party: protoc 

## Нюансы

Базы, темы и т.д. в БД и redis проверяются и при необходимости создаются в самой программе сервера. Настройка clickhouse внешняя.
Конфигурация берется из .env файла в корне папки сервера, если аналогичных установок нет в окружении.


## Стэк

- Go 1.18
  - GORM
  - gRPC
  - Testing (в легкой форме)
- Postgres 14.1
- Redis 5.0
- Kafka 3.2.1
- Clickhouse 22.2.2.1
