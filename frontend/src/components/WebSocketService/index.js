// Desc: WebSocket service to connect to the backend
const WebSocketService = {
  socket: null,
  init(dispatch, token) {
    if (token) {
      this.socket = new WebSocket("ws://localhost:8080/ws", ["token", token]);
      this.socket.onopen = function () {
        console.log("WebSocket connected");
      };
      this.socket.onmessage = function (event) {
        const message = JSON.parse(event.data);
        console.log(message);
        setTimeout(() => {
          alert("You have a new notification!");
        }, 1000); // Delay the alert by 1 second
        dispatch({ type: "UPDATE_MESSAGES", payload: message });
      };
      this.socket.onclose = function (event) {
        dispatch({ type: "CLEAR_MESSAGES" });
        console.log("WebSocket connection closed: ", event);
      };
      this.socket.onerror = function (error) {
        console.log("WebSocket error: ", error);
      };
    } else {
      console.log("No token found in sessionStorage. WebSocket connection skipped.");
    }
  },
// Desc: Close the websocket connection
  close() {
    if (this.socket) {
      this.socket.close();
    }
  },
};

export default WebSocketService;
