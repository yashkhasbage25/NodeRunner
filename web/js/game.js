var player1 = document.getElementById("player1");
var player2 = document.getElementById("player2");

ws = new WebSocket("ws://192.168.105.49:8080/ws");

ws.onmessage = function(event) {
    console.log(event)
    var rect = player2_sprite.getBoundingClientRect();
    data = JSON.parse(event.data)
    document.getElementById("player2").style.left = data.px + "px";
    document.getElementById("player2").style.top = data.py + "px";
}
ws.onclose = function() {
    console.log("   closed  ");
}

document.onkeyup = function(event) {
    var rect = player2.getBoundingClientRect();
    ws.send(JSON.stringify({etype: "up", px: rect.left, py: rect.top}));
}
document.onkeydown = function(event) {
    var rect = player2.getBoundingClientRect();
    ws.send(JSON.stringify({etype: "down", px: rect.left, py: rect.top}));
}
