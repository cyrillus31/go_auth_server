# go_auth_server
gRPC auth server


- [Youtube](https://www.youtube.com/watch?v=EURjTg5fw-E)
- [web page](https://selectel.ru/blog/tutorials/go-grcp/?utm_source=youtube.com&utm_medium=referral&utm_campaign=help_tgbot-grcp_181123_tuzov_paid)


## Protobuf code repo
[https://github.com/cyrillus31/go_auth_server_proto](https://github.com/cyrillus31/go_auth_server_proto)

## Структура проекта
```
sso
├── cmd.............. Команды для запуска приложения и утилит
│	├── migrator.... Утилита для миграций базы данных
│	└── sso......... Основная точка входа в сервис SSO
├── config........... Конфигурационные yaml-файлы
├── internal......... Внутренности проекта
│	├── app.......... Код для запуска различных компонентов приложения
│	│	└── grpc.... Запуск gRPC-сервера
│	├── config....... Загрузка конфигурации
│	├── domain
│	│	└── models.. Структуры данных и модели домена
│	├── grpc
│	│	└── auth.... gRPC-хэндлеры сервиса Auth
│	├── lib.......... Общие вспомогательные утилиты и функции
│	├── services..... Сервисный слой (бизнес-логика)
│	│	├── auth
│	│	└── permissions
│	└── storage...... Слой хранения данных
│	└── sqlite.. Реализация на SQLite
├── migrations....... Миграции для базы данных
├── storage.......... Файлы хранилища, например SQLite базы данных
└── tests............ Функциональные тесты
```
