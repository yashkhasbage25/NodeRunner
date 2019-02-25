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
    };
}

// ws = new WebSocket("ws://192.168.105.49:8080/game");
var clientNumber = -1;
var clientID = "";
var pressedKeys = [];
var playerOne = getElement("player1");
var playerTwo = getElement("player2");

var gemOne = getElement("gem1");
var gemTwo = getElement("gem2");
var gemThree = getElement("gem3");
var gemFounr = getElement("gem4");

var botOne = getElement("bot1");
var botTwo = getElement("bot2");
var botThree = getElement("bot3");

var myHealth = getElement("myHealth");
var otherHealth = getElement("otherHealth");

var playerHeight = playerOne.height;
var playerWidth = playerOne.width;

ws.onopen = function() {
    console.log("game.html websocket opened");
}

ws.onclose = function() {
    console.log("game.html websocket closed");
    ws.send(JSON.stringify({etype: "SocketClosedUnexpectedly", object: clientID}));
}

ws.onmessage = function(event) {
    console.log(event.data);
    // var rect = player2_sprite.getBoundingClientRect();
    data = JSON.parse(event.data)
    if(data.etype == "SetClientID") {
        clientNumber = parseInt(data.object);
        clientID = "p" + (clientNumber + 1).toString();
    } else if (data.etype == "SendUpdate") {
        ws.send(JSON.stringify(getAllCurrentPositions()));
    } else if (data.etype == "Update") {
        setAllPositions(data);
    } else if (data.etype == "Win") {
        alert("You Win");
    } else if (data.etype == "Lose") {
        alret("You Lose");
    }
    console.log("This client has ID:", clientID, clientNumber);
}

function setPosition(elem, position) {
    elem.left = position.x.toString() + "px";
    elem.top = position.y.toString() + "px";
}

function setHealth(elem, health) {
    health /= 10;
    elem.style.width = health.toString() + "%";
    elem.innerHTML = health + '%';
}

function setAllPositions(data) {
    setPosition(playerOne, data.p1_pos);
    setPosition(playerTwo, data.p2_pos);
    setPosition(botOne, data.b1_pos);
    setPosition(botTwo, data.b2_pos);
    setPosition(botThree, data.b3_pos);
    setPosition(gemOne, data.g1_pos);
    setPosition(gemTwo, data.g2_pos);
    setPosition(gemThree, data.g3_pos);
    setPosition(gemFour, data.g4_pos);
    if (clientID == "p1") {
        setHealth(myHealth, data.h1);
        setHealth(otherHealth, data.h2);
    } else {
        setHealth(myHealth, data.h2);
        setHealth(otherHealth, data.h1);
    }
}

function getPositionOfElement(element) {

    let centerX = element.offsetLeft + element.offsetWidth / 2;
    let centerY = element.offsetTop + element.offsetHeight / 2;
    return {
        x: centerX,
        y: centerY
    };
}

function getHealth(elem) {
    healthStr = elem.style.width;
    health = parseInt(healthStr.substring(0, healthStr.length-1));
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
        g4_pos: getPositionOfElement(gemFour),
        h1: getHealth(myHealth) * 10,
        h2: getHealth(otherHealth) * 10
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
    if (event.keyCode == 37) {
        currentPositions.etype = "Left";
    } else if (event.keyCode == 38) {
        currentPositions.etype = "Up";
    } else if (event.keyCode == 39) {
        currentPositions.etype = "Right";
    } else if (event.keyCode == 40) {
        currentPositions.etype = "Down";
    } else if (event.keyCode == 32) {
        currentPositions.etype = "Teleport";
    } else {
        console.log("Unknown keyCode detected", event.keyCode);
    }
    ws.send(JSON.stringify(currentPositions));
}
