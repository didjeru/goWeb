# Задание №1-1

## Условие задачи

Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок.

## Запуск

```go
go run main.go
```

## Тестирование

```shell
$ go run main.go
No matches Посольство in http://yandex.ru
No matches Посольство in http://rambler.ru
No matches Посольство in http://ria.ru
Россия - http://yandex.ru
Россия - http://yandex.ru
Россия - http://rambler.ru
Россия - http://rambler.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
Россия - http://ria.ru
No matches Литва окончательно in http://yandex.ru
No matches Литва окончательно in http://rambler.ru
Литва окончательно - http://ria.ru
Литва окончательно - http://ria.ru
```
