<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.1.5/css/uikit.min.css" />
    <script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
    <script>function doDelete(postId) {$.ajax({ url: '/post/' + postId, type: 'DELETE', success: function () { location.reload() }})}</script>
</head>