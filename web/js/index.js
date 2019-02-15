document.addEventListener("keyup", function(event) {
    console.log(event.keyCode);
    if (event.keyCode == 13) {
        window.location.href='/web/wait.html';
    }
});
