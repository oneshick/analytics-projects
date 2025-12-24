# Мини-система сбора и анализа данных

## Цель
Создать end-to-end систему, которая генерирует данные, сохраняет их в базу данных и позволяет анализировать через Redash и Jupyter Notebook.

## Структура проекта
- `generator/` — генератор данных на Go  
- `db/init.sql` — SQL-схема базы данных  
- `notebooks/analysis.ipynb` — Jupyter Notebook для анализа  
- `docker-compose.yml` — развёртывание всех сервисов  

## Шаги установки и запуска

### 1. Redash

1. Скачать и установить WSL, если ещё не установлен.  
2. Перейти в директорию Redash:
   ```bash
   cd /mnt/c/Users/bro.system/Desktop/course/data-analysis/redash
3. Сделать скрипт setup.sh исполняемым:
    chmod +x setup.sh
4. Запустить установку:
    sudo ./setup.sh

### 2. Запуск своего проекта через Docker Compose

1. Сборка и запуск всех сервисов:
    docker-compose up --build
2. Остановка и удаление контейнеров:
    docker-compose down

### 3. Работа с Jupyter Notebook

1. Проверить список запущенных контейнеров:
    docker ps
2. Зайти в контейнер Jupyter:
    docker exec -it <имя_контейнера_jupyter> bash
3. Настроить пароль:
    jupyter notebook password
4. Перезапустить контейнер:
    docker restart <имя_контейнера_jupyter>