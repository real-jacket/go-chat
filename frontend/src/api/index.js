const socket = new WebSocket('ws://localhost:8080/ws')

const connect = cb => {
    console.log('connecting....')

    socket.onopen = () => {
        console.log('Successfully Connected')
    }

    socket.onmessage = msg => {
        console.log(msg)
        cb(msg)
    }

    socket.onclose = event => {
        console.log('Socket Closed Connection: ', event)
    }

    socket.onerror = error => {
        console.log('Socket msg: ', error)
    }
}

const sendMsg = msg => {
    console.log('sending msg: ', msg)
    socket.send(msg)
}

export { connect, sendMsg }
