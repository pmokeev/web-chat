import React, {SyntheticEvent} from 'react';
import {IJWTInterface} from "../Interfaces/IJWTInterface";
import './styles/login-register.css'
import ErrorMessage from "../Components/ErrorMessage";

class Register extends React.Component<IJWTInterface> {
  that = this;
  private _name: string;
  private _email: string;
  private _password: string;
  private _statusCode: number;

  private get name() {
    return this._name;
  }
  private set name(name: string) {
    this._name = name;
  }
  private get email(): string {
    return this._email;
  }
  private set email(value: string) {
    this._email = value;
  }
  private get statusCode(): number {
    return this._statusCode;
  }
  private set statusCode(value: number) {
    this._statusCode = value;
  }
  private get password(): string {
    return this._password;
  }
  private set password(value: string) {
    this._password = value;
  }

  constructor(props: IJWTInterface) {
    super(props);
    this._name = ``;
    this._email = ``;
    this._password = ``;
    this._statusCode = 0;
  }

  async submitForm(e: SyntheticEvent, name: string, email: string, password: string) {
    e.preventDefault();

    const response = await fetch('http://localhost:8000/api/auth/sign-up', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        name,
        email,
        password,
      })
    });

    this.statusCode = response.status;
  }

  render() {
    return (
      <div className="form-signin">
        <form onSubmit={e => {
          this.submitForm(e, this.name, this.email, this.password);
          console.log(this._statusCode);
        }}>
          <ErrorMessage statusCode={this.statusCode} />
          <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
          <div className="form-floating">
            <input type="name" className="form-control" id="floatingInput" placeholder="name"
                   onChange={e => this.name = e.target.value}
            />
            <label htmlFor="floatingInput">Your name</label>
          </div>
          <div className="form-floating">
            <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com"
                   onChange={e => this.email = e.target.value}
            />
            <label htmlFor="floatingInput">Email address</label>
          </div>
          <div className="form-floating">
            <input type="password" className="form-control" id="floatingPassword" placeholder="Password"
                   onChange={e => this.password = e.target.value}
            />
            <label htmlFor="floatingPassword">Password</label>
          </div>
          <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
      </div>
    );
  }
}

export default Register;