import React from 'react';
import Message from "./Message";

const Messages = (props: { messages: Message[] }) => {
    return (
      <div>
        {props.messages && props.messages.map(message => <p key={message.id}><strong>{message.sender}:</strong> {message.body}</p>)}
      </div>
    )
}

export default Messages;