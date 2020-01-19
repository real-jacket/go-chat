import React, { useEffect, useState } from 'react'
import { connect, sendMsg } from './api'
import Header from './components/header/Header'
import ChatHistory from './components/ChatHistory/ChatHistory'

function App() {
    const [chatHistory, setChatHistory] = useState([])
    useEffect(() => {
        connect(msg => {
            setChatHistory([...chatHistory, msg])
        })
    }, [chatHistory])

    function send() {
        console.log('hello')
        sendMsg('hello')
    }

    return (
        <div className="App">
            <Header />
            <button onClick={send}>Hit</button>
            <ChatHistory chatHistory={chatHistory} />
        </div>
    )
}

export default App
