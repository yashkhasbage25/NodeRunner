ws = new WebSocket("ws://192.168.105.49:8080/wait");

ws.onopen = function() {
    console.log("wait.js websocket opened");
}

ws.onmessage = function(event) {
    data = JSON.parse(event.data)
    console.log("event received: " + data.redirect)
    if (data.redirect) {
        window.location = "/web/game.html"
    }
}

ws.onclose = function() {
    console.log("wait.js websocket closed");
}
