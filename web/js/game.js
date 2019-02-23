// var player1 = document.getElementById("player1");
// var player2 = document.getElementById("player2");

ws = new WebSocket("ws://192.168.105.49:8080/game");
var clientID

ws.onopen = function() {
    console.log("game.html websocket opened");
}

ws.onclose = function() {
    console.log("game.html websocket closed");
}

ws.onmessage = function(event) {
    console.log(event.data);
    // var rect = player2_sprite.getBoundingClientRect();
    data = JSON.parse(event.data)
    if(data.etype == "SetClientID") {
        clientID = parseInt(data.object);
    }
    console.log("This client has ID:", clientID);
    // document.getElementById("player2").style.left = data.px + "px";
    // document.getElementById("player2").style.top = data.py + "px";
}
//
// document.onkeyup = function(event) {
//     var rect = player2.getBoundingClientRect();
//     ws.send(JSON.stringify({etype: "up", px: rect.left, py: rect.top}));
// }
// document.onkeydown = function(event) {
//     var rect = player2.getBoundingClientRect();
//     ws.send(JSON.stringify({etype: "down", px: rect.left, py: rect.top}));
// }
