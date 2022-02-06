import React, {useEffect, useState} from 'react';
import Message from "../components/Message";
import Status from "../components/Status";
import './pages-styles/chat.css';
import Messages from "../components/Messages";
import InputText from "../components/InputText";
import {Redirect} from "react-router-dom";

const webSocketURL = 'ws://localhost:8001/api/chat';
let webSocket: WebSocket = new WebSocket('ws://placeholder'); // TODO: fix in future

const Chat = (props: { isJWTCorrect: boolean }) => {
  const [connect, setConnect] = useState(false);
  const [messages, setMessages] = useState(Array<Message>());
  const [message, setMessage] = useState('');
  const [username, setUsername] = useState('');

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://localhost:8000/api/auth/profile', {
          method: 'GET',
          credentials: 'include',
        });

        const content = response.json();
        content.then(data => {
          setUsername(data["name"]);
        });
      }
    )();
  });

  const enterChat = () => {
    let ws = new WebSocket(webSocketURL);

    ws.onopen = (event) => {
      console.log('Websocket opened!', {event});
      setConnect(true);
    }

    ws.onclose = (event) => {
      console.log('Websocket closed!', {event})
      setConnect(false);
      webSocket = new WebSocket('ws://placeholder');
    }

    ws.onmessage = (message) => {
      console.log('Websocket message: ', {message})

      let tempMessages = messages;
      let msgParsed = JSON.parse(message.data);
      tempMessages.push(new Message(msgParsed.id, username, msgParsed.body));
      setMessages([...messages]);
    }

    ws.onerror = (error) => {
      console.log('Websocket error:', {error})
    }

    webSocket = ws;
  }

  const sendMessage = () => {
    webSocket.send(message);
    setMessage('');
  }

  if (!props.isJWTCorrect) {
    return <Redirect to="/" />;
  }

  return (
    <div className="chat">
      <h1>WebChat</h1>
      <Status status={connect ? 'connected' : 'disconnected'} />
      {
        connect && <Messages messages={messages} />
      }
      {
        connect ?
          <div className="chat-inputs">
            <InputText
              type={"text"}
              placeholder={"Write message"}
              onChange={value => setMessage(value)}
              defaultValue={message} />
          </div> :
          <div />
      }
      <button type="button" onClick={() => connect ? sendMessage() : enterChat()}>
        {connect ? 'Send' : 'Enter'}
      </button>
    </div>
  )
}

export default Chat;