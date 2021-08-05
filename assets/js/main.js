
url = 'ws://'+document.location.host + "/ws";
conn = new WebSocket(url);

conn.onmessage = function(event){
    console.log(event.data)
    var si = document.getElementById("standard-input");
    var messages = event.data.split('\n');
    for (var i=messages.length-1; i> -1; i--) {
        var item = document.createElement("span");
        item.setAttribute("class", "result");
        item.innerText = messages[i];
        si.after(item);
    }

    changeTarget();//4
    changeReadOnly();
    newLine();//6
}

window.addEventListener("load", function(){
    si = document.getElementById("standard-input");
    si.focus();
});

window.addEventListener("keyup", function(e){
    let key = e.code;
    if(key=="Enter") {
        cmd = document.getElementById("standard-input").value; //1
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

function newLine() {
    var list = document.getElementsByClassName("terminal")[0];
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.setAttribute("id", "console");
    var console = document.createTextNode("$");
    span.appendChild(console);
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input");
    input.setAttribute("spellcheck", "false");
    li.appendChild(span);
    li.appendChild(input);
    list.appendChild(li);

    input.focus();
}


function changeTarget() {
    var console = document.getElementById("console");
    console.removeAttribute("id");
    console.setAttribute("class", "console");

    var si = document.getElementById("standard-input");
    si.removeAttribute("id");
    si.setAttribute("class","standard-input");
}


function changeReadOnly() {
    var standardInputs = document.getElementsByClassName("standard-input");
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