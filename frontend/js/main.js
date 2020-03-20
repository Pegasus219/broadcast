$(function() {

    var conn;
    var wsUrl = "ws://localhost:8181/ws";
    var group = "test";
    var showDiv = $("#show");

    function appendLog(msg) {
        var d = showDiv[0];
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        msg.appendTo(showDiv);
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }

    function createWebsocket() {
        if (window["WebSocket"]) {
            conn = new WebSocket(wsUrl + "?group=" + group);
            conn.onclose = function(evt) {
                appendLog($("<div><b>Connection Closed.</b></div>"))
            };
            conn.onmessage = function(evt) {
                appendLog($("<div/>").text(evt.data));
            };
            return true
        } else {
            appendLog($("<div><b>WebSockets Not Support.</b></div>"));
            return false
        }
    }

    createWebsocket();
});