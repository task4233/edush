
url = 'ws://'+document.location.host + "/ws";
conn = new WebSocket(url);

conn.onmessage = function(event){
    console.log(event.data);
    obj = JSON.parse(event.data);
    if(obj.Owner) {
        var si = document.getElementById("standard-input-left");
        var messages = atob(obj.StdOut).split('\n');
        for (var i=messages.length-1; i> -1; i--) {
            var item = document.createElement("span");
            item.setAttribute("class", "result");
            item.innerText = messages[i];
            si.after(item);
        }
        changeTargetLeft();//4
        changeReadOnlyLeft();
        newLineLeft();//6
    }
    else {
        var si = document.getElementById("standard-input-right");
        var messages = atob(obj.StdOut).split('\n');
        for (var i=messages.length-1; i> -1; i--) {
            var item = document.createElement("span");
            item.setAttribute("class", "result");
            item.innerText = messages[i];
            si.after(item);
        }
        changeTargetRight();//4
        changeReadOnlyRight();
        newLineRight()
    }
}

window.addEventListener("load", function(){
    si = document.getElementById("standard-input-left");
    si.focus();
});

window.addEventListener("keyup", function(e){
    let key = e.code;
    if(key=="Enter") {
        cmd = document.getElementById("standard-input-left").value; //1
        if(cmd === "clear") {
            clear();
            return;
        }
        if(!conn) {
            return false;
        }
        if(!cmd) {
            return false;
        }
        conn.send(cmd);
		console.log("[DEBUG] " + cmd);
		return;
    }
})

function newLineLeft() {
    var list = document.getElementsByClassName("terminal")[0];
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.setAttribute("id", "console-left");
    var console = document.createTextNode("$");
    span.appendChild(console);
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input-left");
    input.setAttribute("spellcheck", "false");
    li.appendChild(span);
    li.appendChild(input);
    list.appendChild(li);

    input.focus();
}

function newLineRight() {
    var list = document.getElementsByClassName("terminal")[1];
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.setAttribute("id", "console-right");
    var console = document.createTextNode("$");
    span.appendChild(console);
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input-right");
    input.setAttribute("spellcheck", "false");
    input.readOnly = true;
    li.appendChild(span);
    li.appendChild(input);
    list.appendChild(li);

    input.focus();
}


function changeTargetLeft() {
    var console = document.getElementById("console-left");
    console.removeAttribute("id");
    console.setAttribute("class", "console-left");

    var si = document.getElementById("standard-input-left");
    si.removeAttribute("id");
    si.setAttribute("class","standard-input-left");
}

function changeTargetRight() {
    var console = document.getElementById("console-right");
    console.removeAttribute("id");
    console.setAttribute("class", "console-right");

    var si = document.getElementById("standard-input-right");
    si.removeAttribute("id");
    si.setAttribute("class","standard-input-right");
}


function changeReadOnlyLeft() {
    var standardInputs = document.getElementsByClassName("standard-input-left");
    for (var i = 0; i < standardInputs.length; i++) {
        var si = standardInputs[i];
        si.readOnly = true;
    }
}

function changeReadOnlyRight() {
    var standardInputs = document.getElementsByClassName("standard-input-right");
    for (var i = 0; i < standardInputs.length; i++) {
        var si = standardInputs[i];
        si.readOnly = true;
    }
}

function clear() {
    var terminal = document.getElementsByClassName("terminal")[0];
    while( terminal.firstChild ){
        terminal.removeChild( terminal.firstChild );
      }
    newLine();
}