import React from 'react';
import Message from "./Message";

const Messages = (props: { messages: Array<Message> }) => {
  return (
    props.messages && props.messages.map(message => <p key={message.id}><strong>{message.sender}:</strong> {message.body}</p>)
  )
}

export default Messages;