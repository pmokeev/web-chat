import React from 'react';
import './components-styles/status.css';

const Status = (props: { status: string }) => {
  return (
    <span className="status">
      <span className={`status-icon ${props.status}`}></span>
      <span className="status-text">{props.status}</span>
    </span>
  )
}

export default Status;