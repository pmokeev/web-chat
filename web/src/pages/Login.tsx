import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from "react-router-dom";
import Home from "./Home";

const Login = (props: { isJWTCorrect: boolean, setIsJWTCorrect: (isJWTCorrect: boolean) => void }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [statusCode, setStatusCode] = useState(0);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    const response = await fetch('http://localhost:8000/api/auth/sign-in', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
      body: JSON.stringify({
        email,
        password,
      })
    });

    setStatusCode(response.status);
    props.setIsJWTCorrect(true);
  }

  if (props.isJWTCorrect || statusCode === 200) {
    return <Redirect to="/"/>
  }

  return (
    <form onSubmit={submit}>
      {statusCode === 409 ? <h5 className="ErrorMsg">Incorrect email/password</h5> : ""}
      <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
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
      <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
    </form>
  );
};

export default Login;