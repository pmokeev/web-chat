import React, {useState} from 'react';
import Message from "../Components/Message";
import Status from "../Components/Status";
import './pages-styles/chat.css';
import Messages from "../Components/Messages";
import InputText from "../Components/InputText";

const webSocketURL = 'ws://localhost:8001/api/chat';
let webSocket: WebSocket = new WebSocket('ws://placeholder'); // TODO: fix in future

const Chat = () => {
  const [connect, setConnect] = useState(false);
  const [messages, setMessages] = useState(Array<Message>());
  const [message, setMessage] = useState('');

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
      tempMessages.push(new Message(msgParsed.id, msgParsed.sender, msgParsed.body));
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