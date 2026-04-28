# 🚀 Go Backend Portfolio — Nikita Solokha

Привет! Я — Никита Солоха, студент-бэкендер, изучаю Go и развиваю навыки в построении распределённых backend-систем.

Этот репозиторий — моё портфолио проектов, где я применяю:
- микросервисную архитектуру
- REST / gRPC
- очереди сообщений
- Docker и Kubernetes
- работу с PostgreSQL и Redis

---

# 📌 🚖 Ride-Sharing Microservices Platform (Основной проект)

Распределённая система для сервиса такси / райдшеринга, построенная на микросервисной архитектуре.

## 🔧 Стек
Go, gRPC, REST API, RabbitMQ, PostgreSQL, MongoDB, Redis, Docker, Kubernetes, Web (Next.js), OpenTelemetry

## 🧱 Архитектура
Проект разделён на микросервисы:

- API Gateway — точка входа для клиентов
- Trip Service — управление поездками
- Driver Service — работа с водителями
- Payment Service — обработка платежей
- Web Client — frontend (Next.js)
- Shared — общие библиотеки, proto, contracts

## ⚙️ Инфраструктура
- Docker для локального запуска
- Kubernetes манифесты (dev / prod)
- RabbitMQ для событий между сервисами
- gRPC для внутренних вызовов
- OpenTelemetry / tracing

## 💡 Функциональность
- создание и управление поездками
- назначение водителей
- обработка статусов поездки
- платежные сценарии
- потоковая коммуникация через события

👉 Демонстрирует навыки работы с реальной distributed системой

---

# 📌 Snippetbox

Веб-приложение для хранения и управления текстовыми сниппетами (заметками, фрагментами кода).

## 🔧 Стек
Go, PostgreSQL, HTML/CSS, Bootstrap, Git, Linux

## ✅ Функционал
- CRUD-операции
- работа с PostgreSQL через `database/sql`
- шаблонизация `html/template`
- кэширование шаблонов
- структура с разделением слоёв (handler / service / repository)
- защита от XSS и SQL-инъекций

---

# 📌 JWT Auth API

REST API с аутентификацией на основе JWT.

## 🔧 Стек
Go, JWT, REST API, JSON, middleware

## ✅ Функционал
- регистрация и логин пользователя
- генерация и валидация JWT токенов
- защита приватных эндпоинтов
- middleware-авторизация
- работа с JSON и HTTP статусами

---

# 🧠 Навыки

## Backend
- Go (net/http, goroutines, channels)
- REST API / gRPC
- PostgreSQL, MongoDB, Redis
- RabbitMQ

## DevOps
- Docker / docker-compose
- Kubernetes (basic deployments)
- Linux

## Архитектура
- Clean Architecture
- Microservices
- Event-driven systems

## Инструменты
- Git / GitHub
- OpenTelemetry (basic tracing)
- Postman

---

# 📬 Контакты

- Telegram: @nikita_solokha
- Email: nikitasoloha166@gmail.com

---

# 🎯 Цель

Ищу стажировку или позицию **Junior Golang Backend Developer**, где смогу развиваться в backend, распределённых системах и high-load архитектуре.