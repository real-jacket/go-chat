import React, { useEffect, useState } from 'react'
import { connect, sendMsg } from './api'
import Header from './components/Header/Header'
import ChatHistory from './components/ChatHistory/ChatHistory'
import ChatInput from './components/ChatInput/ChatInput'

function App() {
    const [chatHistory, setChatHistory] = useState([])
    useEffect(() => {
        connect(msg => {
            setChatHistory([...chatHistory, msg])
        })
    }, [chatHistory])

    function send(event) {
        if (event.keyCode === 13) {
            sendMsg(event.target.value)
            event.target.value = ''
        }
    }

    return (
        <div className="App">
            <Header />
            <ChatHistory chatHistory={chatHistory} />
            <ChatInput send={send} />
        </div>
    )
}

export default App
