import React from 'react';

class Message {
  id: number;
  sender: string;
  body: string;

  constructor(id: number, sender: string, body: string) {
    this.id = id;
    this.sender = sender;
    this.body = body;
  }
}

export default Message;