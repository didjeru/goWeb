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
    <form method="POST" action="/edit/{{.Post.ID}}">
        <div class="form-group">
            <label>ID</label>
            <input type="id" name="id" class="form-control" value="{{.Post.ID}}">
        </div>
        <div class="form-group">
            <label>Title</label>
            <input type="title" name="title" class="form-control" value="{{.Post.Title}}">
        </div>
        <div class="form-group">
            <label>Content</label>
            <textarea name="content" class="form-control">{{.Post.Content}}</textarea>
        </div>
        <input class="btn btn-primary" type="submit" value="Save">
        <a class="btn btn-outline-primary" href="/">Back</a>
    </form>
</div>
{{template "footer.tpl"}}
</body>
</html>