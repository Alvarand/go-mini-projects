# Todo App - консольное приложение для создания todo list'а

Данные хранятся в csv формате

**Для начала необходимо создать исполняемый файл командой**: `go build -o todoapp [путь к cmd/main.go]`

## Возможности

- добавить задачу
- удалить задачу
- вывести список задач
- завершить задачу

## Возможные команды

### Добавить задачу

```text
./todoapp add <desciption>
```

Пример:

```text
./todoapp add "Создать todo app"
```

---

### Удалить задачу

```text
./todoapp delete <taskID>
```

Пример:

```text
./todoapp delete 2
```

---

### Вывести список задач

```text
./todoapp list
```

---

### Завершить задачу

```text
./todoapp complete <taskID>
```

Пример:

```text
./todoapp complete 3
```
