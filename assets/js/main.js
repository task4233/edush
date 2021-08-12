
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
        changeTarget("left");//4
        changeRead("left");
        newLine("left");//6
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
        changeTarget("right");//4
        changeRead("right");
        newLine("right")
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

function newLine(mode) {
    switch(mode) {
        case 'left':
            var list = document.getElementsByClassName("terminal")[0];
            break
        case 'right':
            var list = document.getElementsByClassName("terminal")[1];
            break
    }
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.setAttribute("id", "console-" + mode);
    var console = document.createTextNode("$");
    span.appendChild(console);
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input-" + mode);
    input.setAttribute("spellcheck", "false");
    if (mode=='right') input.readOnly = true;
    li.appendChild(span);
    li.appendChild(input);
    list.appendChild(li);

    input.focus();
}

function changeTarget(mode) {
    var console = document.getElementById("console-" + mode);
    console.removeAttribute("id");
    console.setAttribute("class", "console-" + mode);

    var si = document.getElementById("standard-input-" + mode);
    si.removeAttribute("id");
    si.setAttribute("class","standard-input-" +  mode);
}

function changeRead(mode) {
    var standardInputs = document.getElementsByClassName("standard-input-" + mode);
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