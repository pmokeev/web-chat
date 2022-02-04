import React from 'react';

interface IErrorMessage {
  statusCode: number;
}

class ErrorMessage extends React.Component<IErrorMessage> {
  statusCode: number;

  constructor(props: ErrorMessage) {
    super(props);
    this.statusCode = props.statusCode;
  }

  render() {
    return (
      <div>
        {this.statusCode === 409 ? <h5 className="ErrorMsg">Error, this email already exist</h5> : ""}
      </div>
    );
  }
}

export default ErrorMessage;