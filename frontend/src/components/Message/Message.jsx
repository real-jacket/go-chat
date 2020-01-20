import React, { useState } from 'react'
import './Message.scss'

function Message(props) {
    const temp = JSON.parse(props.message)
    const [message] = useState(temp)
    return <div className="Message">{message.body}</div>
}

export default Message
