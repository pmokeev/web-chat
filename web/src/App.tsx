import React, {useEffect, useState} from 'react';
import './App.css';
import Login from "./pages/Login";
import Navigation from "./components/Navigation";
import {BrowserRouter, Route} from "react-router-dom";
import Register from "./pages/Register";
import Home from "./pages/Home";

function App() {
  const [isJWTCorrect, setIsJWTCorrect] = useState(false);

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://localhost:8000/api/auth/jwtverify', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });

        const statusCode = response.status;

        if (statusCode === 200) {
          setIsJWTCorrect(true);
        } else {
          setIsJWTCorrect(false);
        }
      }
    )();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Navigation/>

        <main className="form-signin">
          <Route exact path="/" component={() => <Home isJWTCorrect={isJWTCorrect}/>}/>
          <Route path="/login" component={Login}/>
          <Route path="/register" component={Register}/>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
