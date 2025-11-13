const server = 'ws://localhost:1234'
const sessionEndpoint = '/api/v1/session'

const socket = new WebSocket(`${server}${sessionEndpoint}`);

socket.addEventListener('open', (event) => {
    console.log('WebSocket connection opened:', event);
});

socket.addEventListener('message', (event) => {
  console.log('Message received from server:', event.data);
  try {
    document.getElementById('commandResponse').innerHTML = event.data;
    document.getElementById('commandBuffer').innerHTML = '';
  } catch {
    console.error("error ????");
  }
});

socket.addEventListener('error', (event) => {
  console.error('WebSocket error:', event);
});

socket.addEventListener('close', (event) => {
    console.log('WebSocket connection closed:', event);
    if (event.wasClean) {
        console.log(`Connection closed cleanly, code=${event.code}, reason=${event.reason}`);
    } else {
        console.error('Connection died unexpectedly');
    }
});

function sendCommand() {
  let command = document.getElementById('commandBuffer').value;
  if (socket.readyState === WebSocket.OPEN) {
      socket.send(command);
      console.log('Sent message:', command);
  } else {
      console.warn('WebSocket is not open. Cannot send message.');
  }
}