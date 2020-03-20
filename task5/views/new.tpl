<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{.Post.Title}}</title>
    {{template "head.tpl"}}
</head>
<body>
{{template "header.tpl"}}
<div class="container">
    <h3>Post</h3>
    <form method="POST" action="/post">
        <div class="form-group">
            <label>Title</label>
            <input type="title" name="title" class="form-control">
        </div>
        <div class="form-group">
            <label>Content</label>
            <textarea name="content" class="form-control"></textarea>
        </div>
        <input class="btn btn-primary" type="submit" value="Save">
        <a class="btn btn-outline-primary" href="/">Back</a>
    </form>
</div>
{{template "footer.tpl"}}
</body>
</html>