<!DOCTYPE html>
<head>
    <title>Electronic Realtime Poll written in Go</title>
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.2.13/semantic.min.css">
    <script src="https://code.jquery.com/jquery-3.1.1.min.js" integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8=" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.2.13/semantic.min.js"></script>
    <style type="text/css"> h2 { margin: 2em 0em; }  .ui.container { padding-top: 5em; padding-bottom: 5em; } </style>
    <script src="https://js.pusher.com/4.1/pusher.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/canvasjs/1.7.0/canvasjs.js"></script>
</head>
<body>
    <div class="ui container">
        <h2 class="ui header">TV Show - Poll Results</h2><br/>
        <div class="ui one column grid link cards">
            <div class="card">
                <div id="canvas" style="height: 500px; width: 100%;"></div>
            </div>
        </div>
    </div>
    <div style="color:#fff" id="key">{{.Key}}</div>
    <div style="color:#fff" id="cluster">{{.Cluster}}</div>
    <script src="/chart.js"></script>
</body>
</html>