import React from 'react'
import './ChatInput.scss'

function ChatInput(props) {
    return (
        <div className="ChatInput">
            <input onKeyDown={props.send} />
        </div>
    )
}

export default ChatInput
