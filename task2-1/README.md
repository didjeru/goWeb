# Задание №2

## Условие задачи

1.Используя функцию для поиска из прошлого практического задания, постройте сервер, который будет принимать JSON с поисковым запросом в POST-запросе и возвращать ответ в виде массива строк в JSON.

```JSON
{
  "search":"фраза для поиска",
  "sites": [
      "первый сайт",
      "второй сайт"
  ]
}
```

2.Напишите два роута: один будет записывать информацию в Cookie (например, имя), а второй — получать ее и выводить в ответе на запрос.

## Запуск

```go
go run main.go
```

## Тестирование

```shell
$ go run main.go
2020/02/25 01:05:53 start listen on port 8080
```

```shell
$ curl -v -H "Content-Type: application/json" -d'{"search": "Россия", "sites": ["https://yandex.ru", "https://golang.org", "https://ria.ru"]}' http://localhost:8080/search
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /search HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 98
>
* upload completely sent off: 98 out of 98 bytes
< HTTP/1.1 200 OK
< Date: Sun, 01 Mar 2020 05:32:04 GMT
< Content-Length: 38
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
["https://yandex.ru","https://ria.ru"]* Closing connection 0
```
