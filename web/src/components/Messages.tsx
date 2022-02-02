import React from 'react';
import Message from "./Message";

const Messages = (props: { messages: Array<Message> }) => {
  return (
    <p key={props.messages[0].id}><strong>{props.messages[0].sender}:</strong> {props.messages[0].body}</p>
  )

  /*{props.messages.map((message: Message) => {
    return (
      <p key={message.id}><strong>{message.sender}:</strong> {message.body}</p>
    );
  })}*/
}

export default Messages;