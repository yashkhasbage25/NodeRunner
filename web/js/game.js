// var player1 = document.getElementById("player1");
// var player2 = document.getElementById("player2");

function getElement(id) {
    item = document.getElementById(id);
    if(item == null) {
        throw "No element found with id " + id;
    }
    itemHeight = item.style.height;
    itemWidth = item.style.width;
    return {
        elem: item,
        height: itemHeight,
        width: itemWidth
    }
}

ws = new WebSocket("ws://192.168.105.49:8080/game");
var clientNumber = -1
var clientID = ""
var pressedKeys = []
var playerOne = getElement("player1");
var playerTwo = getElement("player2");

var gemOne = getElement("gem1");
var gemTwo = getElement("gem2");
var gemThree = getElement("gem3");
var gemFounr = getElement("gem4");

var botOne = getElement("bot1");
var botTwo = getElement("bot2");
var botThree = getElement("bot3");

var playerHeight = playerOne.height;
var playerWidth = playerOne.width;

ws.onopen = function() {
    console.log("game.html websocket opened");
}

ws.onclose = function() {
    console.log("game.html websocket closed");
    ws.send(JSON.stringify({etype: "SocketClosedUnexpectedly", object: clientID}))
}

ws.onmessage = function(event) {
    console.log(event.data);
    // var rect = player2_sprite.getBoundingClientRect();
    data = JSON.parse(event.data)
    if(data.etype == "SetClientID") {
        clientNumber = parseInt(data.object);
        clientID = "p" + (clientNumber + 1).toString();
    } else if (data.etype == "SendUpdate") {
        ws.send(JSON.stringify(getAllCurrentPositions()))
    }
    console.log("This client has ID:", clientID, clientNumber);
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

function getPositionOfElement(element) {

    let centerX = element.offsetLeft + element.offsetWidth / 2;
    let centerY = element.offsetTop + element.offsetHeight / 2;
    return {
        x: centerX,
        y: centerY
    };
}

function getCurrentPositions() {
    return {
        etype: "None",
        object: "None",
        p1_pos: getPositionOfElement(playerOne),
        p2_pos: getPositionOfElement(playerTwo),
        b1_pos: getPositionOfElement(botOne),
        b2_pos: getPositionOfElement(botTwo),
        b3_pos: getPositionOfElement(botThree),
        g1_pos: getPositionOfElement(gemOne),
        g2_pos: getPositionOfElement(gemTwo),
        g3_pos: getPositionOfElement(gemThree),
        g4_pos: getPositionOfElement(gemFour)
    };
}

function removeFromArray(array, value) {
   return array.filter(function(item) {
       return item != value;
   });
}

var result = arrayRemove(array, 6);

document.onkeydown = function(event) {
    pressedKeys.push(event.keyCode);
}

document.onkeyup = function(event) {
    pressedKeys = arrayRemove(pressedKeys, event.keyCode);
}

document.onkeypress = function(event) {
    currentPositions = getCurrentPositions();
    currentPositions.object = clientID;
    if(event.keyCode == 37) {
        currentPositions.etype = "Left";
    } else if (event.keyCode == 38) {
        currentPositions.etype = "Up";
    } else if (event.keyCode == 39) {
        currentPositions.etype = "Right";
    } else if (event.keyCode == 40) {
        currentPositions.etype = "Down";
    }
    ws.send(JSON.stringify(currentPositions));
}
//
// var keys = [];
// window.addEventListener("keydown",
//     function(e){
//         keys[e.keyCode] = true;
//         checkCombinations(e);
//     },
// );
//
// window.addEventListener('keyup',
//     function(e){
//         keys[e.keyCode] = false;
//     },
// false);
//
// function checkCombinations(e){
//     if(keys["a".charCodeAt(0)] && e.ctrlKey){
//         alert("You're not allowed to mark all content!");
//         e.preventDefault();
//     }
// }
