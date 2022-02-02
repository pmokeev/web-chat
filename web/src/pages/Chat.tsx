import React, {useState} from 'react';
import Message from "../components/Message";
import Status from "../components/Status";
import './pages-styles/chat.css';
import Messages from "../components/Messages";
import InputText from "../components/InputText";

const webSocketURL = 'ws://localhost:8001/api/chat';

const Chat = () => {
  const [connect, setConnect] = useState(false);
  let webSocket: WebSocket = new WebSocket('ws://placeholder'); // TODO: fix in future
  let messages: Message[] = [];
  let message: string = '';
  const username = "USER";

  const setMessage = (value: string) => {
    message = value;
  }

  const setMessages = (value: Message) => {
    messages.push(value);
  }

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
      setMessages(message.data);
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

  return (
    <div className="chat">
      <h1>WebChat</h1>
      <Status status={connect ? 'connected' : 'disconnected'} />
      {
        connect && messages && <Messages messages={messages} />
      }
      <div className="chat-inputs">
        <InputText type={'text'}
                   placeholder={connect ? 'Write message' : 'Enter with your username'}
                   onChange={value => connect ? setMessage(value) : username}
                   defaultValue={connect ? message.toString() : username} />
      </div>
      <button type="button" onClick={() => connect ? sendMessage() : enterChat()}>
        {connect ? 'Send' : 'Enter'}
      </button>
    </div>
  )
}

export default Chat;