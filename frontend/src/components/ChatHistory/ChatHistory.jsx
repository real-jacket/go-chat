import React from 'react'
import './ChatHistory.scss'
import Message from '../Message/Message'

function ChatHistory(props) {
    const message = props.chatHistory.map((msg, index) => (
        <Message key={index} message={msg.data} />
    ))
    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {message}
        </div>
    )
}

export default ChatHistory
