10 тестов на понимание golang. Нужно прочитав код ответить на один и тот же вопрос - "что выполняет программа",
далее можно проверить свой ответ запустив конкрентный файл

```bash
go get -d -u github.com/dmitryrpm/maxima-tests
cd ${GOPATH}/src/github.com/dmitryrpm/maxima-tests
go run test1.go
# etc
```

Так же популярные вопросы задаваемые на интервью

Общие вопросы:

- Как хранятся переменные в Golang? Что такое "стек" и "куча"? Почему аллокация в "куче" дорогая? Во сколько раз?
- Как в golang освобождаетс память и можно ли отключить это поведение и зачем это делать?
- Что такое интерфейс и как используется? Как устроен пустой интерфейс? 
- Как устроен слайс и чем он отличается от массива? Как создать многомерный массив в Golang? Нужно ли передавать slice по ссылке в фукнцию?
- Что происходит в runtime Golang?
- В чем различия goroutine от потока системы?
- Как огранить число потоков на системы при запуске Golang программы и возможно ли огранить их до 1 потока?
- Что такое каналы и каких видов они бывают? Что будет если писать в закрытый канал? Что будет если писать в неинициализированный канал?
- Расскажите о ООП в Golang

Вопросы по database/sql:

- Когда создается и закрывается подключение к БД?
- Как ограничить кол-во подключений к БД, сколько их создается по умолчанию?
- Что если в пуле нет соединеней?
- Что произойдет если БД закроет соединение?

PS: Если Вы посещали собеседования и нашли интересную задачку/тест/вопрос по языку Golang - смело делайте PR в мастер, давайте вместе расширим базу знаний, которая поможет множеству программистов разных уровней разобраться в особенности языка. Любые вопросы/пожеления/pr (особенно в целях обучения) и тд - приветствуются!