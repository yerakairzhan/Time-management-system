# ⏱ Time Management System

API для измерения времени, потраченного на выполнение задач.

---

## 🧩 Возможности:
- ✅ Создание задачи с названием, описанием, приоритетом и сроком выполнения (дедлайном)
- ▶️ Запуск таймера на задачу при начале выполнения
- 📊 Получение мини-отчёта по каждой выполненной задаче

📄 Подробнее: [Документация в Notion](https://www.notion.so/Time-Management-System-1d5bef225bcb80dc9811f15f88e763e8?pvs=4)

---

## 🚀 Как развернуть проект

### 🔁 Клонирование репозитория


git clone https://github.com/yerakairzhan/Time-management-system
cd Time-management-system

🛠 Развёртывание базы данных

Создайте таблицы с помощью следующих SQL-запросов:

```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(255) NOT NULL,
    priority VARCHAR(50) NOT NULL CHECK (priority IN ('low', 'medium', 'high')),
    deadline TIMESTAMP
);

CREATE TABLE task_time_logs (
    id SERIAL PRIMARY KEY,
    task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255),
    message TEXT,
    read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

▶️ Запуск проекта

`go run main.go`
