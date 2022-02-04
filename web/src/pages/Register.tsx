import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from "react-router-dom";
import './pages-styles/login-register.css';

const Register = (props: { isJWTCorrect: boolean }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [statusCode, setStatusCode] = useState(0);

  const submit = async (e: SyntheticEvent) => {
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

    setStatusCode(response.status);
  }

  if (props.isJWTCorrect) {
    return <Redirect to="/"/>
  }

  if (statusCode === 200) {
    return <Redirect to="/login"/>
  }

  return (
    <form onSubmit={submit} className="form-signin">
      {statusCode === 409 ? <h5 className="ErrorMsg">Error, this email already exist</h5> : ""}
      <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
      <div className="form-floating">
        <input type="name" className="form-control" id="floatingInput" placeholder="name"
          onChange={e => setName(e.target.value)}
        />
        <label htmlFor="floatingInput">Your name</label>
      </div>
      <div className="form-floating">
        <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com"
          onChange={e => setEmail(e.target.value)}
        />
        <label htmlFor="floatingInput">Email address</label>
      </div>
      <div className="form-floating">
        <input type="password" className="form-control" id="floatingPassword" placeholder="Password"
          onChange={e => setPassword(e.target.value)}
        />
        <label htmlFor="floatingPassword">Password</label>
      </div>
      <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
  );
};

export default Register;