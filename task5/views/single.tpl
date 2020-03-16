<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>{{.Post.Title}}</title>
    {{template "head.tpl"}}
</head>
<body>
{{template "header.tpl"}}
<div class="uk-card uk-card-default uk-card-body">
    <h3>{{.Post.Title}}</h3>
    <div class="uk-card uk-card-default uk-card-body">
        <p>{{.Post.Content}}</p>
    </div>
    <br>
    <a class="uk-button uk-button-default" href="/">ВЕРНУТЬСЯ НА ГЛАВНУЮ >></a>
</div>
{{template "footer.tpl"}}
</body>
</html>