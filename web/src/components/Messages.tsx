import React from 'react';
import Message from "./Message";

const Messages = (props: { messages: Message[] }) => {
  return (
    <div>{props.messages.length}</div>
  )
}

export default Messages;