# Тестовый бот
Бот позволяет получать прогноз погоды и курсы валют. Написан с использованием clean architecture.
Данные о погоде и о курсах кэшируются в базе данных с временем жизни 2 часа.
Для работы с телеграммом использовал свою библиотеку, которую писал последние несколько дней под свои задачи.Было бы оптимальнее использовать готовое решение, но для тестового бота решил юзать свою.

# TODO
Тесты

Подгрузка сообщений с файлов, а не хардкод.

# Запуск

Ставим токен бота в docker-compose.yml

Запуск осуществляется при помощи docker compose.
Docker-compose файл включает в себя базу данных,и само приложение.

```
docker-compose up -d
```

# Описание архитектуры

├── cmd                     
│   └── main.go                  //Точка входа
├── configs                      //Конфиги
│   └── config.yml
├── domain                       //Сущности
│   ├── currency.go
│   ├── errors.go
│   ├── statistics.go
│   ├── types.go
│   └── weather.go
├── handler                      //Обработчики для тг
│   ├── currecny_handler.go
│   ├── handler.go
│   ├── stat_handler.go
│   └── weather_handler.go
├── migrations                   //Миграции
│   ├── 000001_init.down.sql
│   └── 000001_init.up.sql
├── repository                   //Репозиторий постгрес
│   ├── currency_postgres.go
│   ├── postgres.go
│   ├── repository.go
│   ├── statistics_postgres.go
│   └── weather_postgres.go
├── usecase                     //Логика
│   ├── currency_usecase.go
│   ├── statistics_ucase.go
│   ├── usecase.go
│   └── weather_ucase.go
├── utils                       //Утилиты для парсинга погоды и курсов
│   ├── currency
│   └── weather
└── wait-for-postgres.sh        //Ожидания запуска постгреса для докера