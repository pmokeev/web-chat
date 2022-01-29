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
        <Navigation isJWTCorrect={isJWTCorrect} setIsJWTCorrect={setIsJWTCorrect}/>

        <main className="form-signin">
          <Route exact path="/" component={() => <Home isJWTCorrect={isJWTCorrect}/>}/>
          <Route path="/login" component={() => <Login isJWTCorrect={isJWTCorrect}/>}/>
          <Route path="/register" component={() => <Register isJWTCorrect={isJWTCorrect}/>}/>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
