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

  toJSON() {
    return {
      id: this.id,
      sender: this.sender,
      body: this.body,
    }
  }

  public toString = () : string => {
    return "Id";
  }
}

export default Message;