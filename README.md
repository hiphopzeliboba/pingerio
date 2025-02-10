### Клонирование и запуск (не запустится)
- `git clone https://github.com/hiphopzeliboba/pingerio.git` - клонирование репозитория
- `cd pingerio` - переход в директорию pingerio/
- `docker-compose up -d --buld` - запуск проекта в docker в фоновом режиме

не успел отладить проект, поэтому он не запускается как надо, и написать человеческие миграции))

### Структура проекта
```
pingerio/
│── backend/
│   │── cmd/
│   │   └── main.go                 # Точка входа в приложение
│   │
│   │── internal/
│   │   ├── api/
│   │   │   ├── handler/
│   │   │   │   └── container.go   # Обработчик API контейнеров
│   │   │   ├── router/
│   │   │   │   └── container.go   # Настройка маршрутов
│   │   ├── db/
│   │   │   └── postgres.go        # Коннектор к базе данных
│   │   ├── model/
│   │   │   └── container.go       # Определение моделей данных
│   │   ├── repository/
│   │   │   ├── containers/
│   │   │   │   └── repository.go  # Логика работы с БД (имплиментация)
│   │   │   └── repository.go      # Логика работы с БД (интерфейсы)
│   │   ├── service/
│   │   │   ├── containers/
│   │   │   │   └── service.go     # Бизнес-логика приложения (имплиментация)
│   │   │   └── service.go         # Бизнес-логика приложения (интерфейсы)
│   │
│   │── migrations/                # Миграции базы данных
│   │── .env                       # Файл конфигурации окружения
│   │── Dockerfile                 # Docker-контейнер для backend
│   │── go.mod                      # Файл зависимостей Go
│   │── makefile                    # Makefile для автоматизации задач
│
│── frontend/
│   │── public/
│   │   └── index.html              # Главный HTML-файл
│   │
│   │── src/
│   │   ├── api.ts                  # Запросы к API
│   │   ├── App.css                 # Стили
│   │   ├── App.tsx                 # Основной компонент приложения
│   │   ├── index.tsx               # Точка входа в React-приложение
│   │
│   │── app.jsx                      # Альтернативный вариант главного компонента
│   │── Dockerfile                   # Docker-контейнер для frontend
│   │── nginx.conf                   # Конфигурация nginx
│   │── package.json                 # Зависимости проекта
│   │── tsconfig.json                 # Конфигурация TypeScript
│
│── pinger/
│   │── cmd/
│   │   └── main.go                 # Точка входа в приложение
│   │
│   │── Dockerfile                 # Docker-контейнер для pinger
│   │── go.mod                      # Файл зависимостей Go
```
