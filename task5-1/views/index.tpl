<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Блог</title>
    {{template "head.tpl"}}
</head>
<body>
{{template "header.tpl"}}
<div class="uk-card uk-card-default uk-card-body">
    <h1>Блог</h1>
    <h2>Приветствую!</h2>
    <p>Заголовок</p>
    <div class="uk-card uk-card-body">
        <ul class="uk-list">
            {{range .Posts}}
                <li>
                    <div class="uk-card uk-card-default uk-card-body">
                        <h3>{{.Title}}</h3>
                        <p>{{.Content}}</p>
                        <a class="uk-button uk-button-default" href="/post/{{.ID}}">ПОДРОБНЕЕ</a>
                        <a class="uk-button uk-button-default" href="/prepare/{{.ID}}">РЕДАКТИРОВАТЬ ПОСТ</a>
                    </div>
                </li>
            {{end}}
        </ul>
    </div>
</div>
{{template "footer.tpl"}}
</body>
</html>