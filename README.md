# Task Runner API

## 1. Запуск


```bash
go run cmd/main.go
```

### 2. Отправка запросов
## Создать задачу
```
curl -X POST http://localhost:8080/tasks
# {"id":"0b9894f0-f9b8-4d1d-8bae-489f04035497","created_at":"2025-06-13T01:21:51.676127+05:00","started_at":"2025-06-13T01:21:51.677265+05:00","status":"running"}
```

## Получить статус задачи
```
curl http://localhost:8080/tasks/<task_id>

# {"id":"0b9894f0-f9b8-4d1d-8bae-489f04035497","created_at":"2025-06-13T01:21:51.676127+05:00","started_at":"2025-06-13T01:21:51.677265+05:00","finished_at":"2025-06-13T01:21:54.678327+05:00","status":"completed","result":"task completed successfully"}
```

## Удалить задачу
```
curl -X DELETE http://localhost:8080/tasks/<task_id>

#{
  "message": "task deleted successfully"
}

```
