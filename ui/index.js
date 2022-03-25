function Action(type, payload, sessionId) {
    this.type = type;
    this.payload = payload;
    this.source = "ui";
    this.timestamp = Date.now();
    this.sessionId = sessionId;
}

(function () {
    const btnSearch = document.getElementById("btnSearch");
    const txtSearch = document.getElementById("search");
    const resSearch = document.getElementById("searchResult");

    var searchList = [];
    
    var connectionId;
    const wsConn = new WebSocket("ws://127.0.0.1:8484/ws");

    function initApplication() {
        btnSearch.addEventListener("click", function() {
            console.log("click", txtSearch.value);

            searchList = [txtSearch.value, ...searchList];
            wsConn.send(JSON.stringify(new Action("search", txtSearch.value, connectionId)));
        })
    }

    function handleEvent(event) {
        const action = JSON.parse(event.data);
        console.log(action);

        if (action.type === "ui:welcome") {
            connectionId = action.payload.connectionId;
            localStorage.setItem("gobullet:sessionId", connectionId);
        }

        if (action.type === "UI:SEARCH") {
            resSearch.setAttribute("href", action.payload);
            resSearch.innerText = searchList[0];
        }
    }

    wsConn.onopen = initApplication

    wsConn.onclose = function(evt) {
        console.log("CLOSE");
        ws = null;
    }

    wsConn.onerror = function(err) {
        console.log(err)
        console.log("ERROR: " + JSON.stringify(err, ["message", "arguments", "type", "name"]));
    }

    wsConn.onmessage = handleEvent
})()