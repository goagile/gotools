var input = document.getElementById("input")
var output = document.getElementById("output")

let host = "localhost"
let port = 8081
let addr = `ws://${host}:${port}/ws`
var socket = new WebSocket(addr)

socket.onopen = () => {
    output.innerHTML += "Status: Connected\n"
}

socket.onmessage = (e) => {
    output.innerHTML += "Server: " + e.data + "\n"
}

let send = () => {
    socket.send(input.value)
    input.value = ""
}
