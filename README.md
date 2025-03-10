## ТЗ

Сделать REST API на Go для создания/удаления/редактирования юзеров. Любой framework (или без него).

Запушить код на github. В идеале с unit тестами. БД - PostgreSQL.

### Запросы:
- `POST /users` - create user
- `GET /user/<id>` - get user
- `PATCH /user/<id>` - edit user

```go
type User struct {
    ID uuid
    Firstname string
    Lastname string
    Email string
    Age uint
    CreatedAt time.Time
}
```

`ID` / `CreatedAt` генерим сами. Остальные — обязательны, и валидируем на входе. Результат завернуть в docker-compose