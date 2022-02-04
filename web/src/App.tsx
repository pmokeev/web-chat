import React, {useEffect, useState} from 'react';
import './App.css';
import Login from "./pages/Login";
import Navigation from "./Components/Navigation";
import {BrowserRouter, Route} from "react-router-dom";
import Register from "./pages/Register";
import Home from "./pages/Home";
import Profile from "./pages/Profile";
import Chat from "./pages/Chat";

function App() {
  const [isJWTCorrect, setIsJWTCorrect] = useState(false);

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://localhost:8000/api/auth/jwtverify', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });

        response.status === 200 ? setIsJWTCorrect(true) : setIsJWTCorrect(false);
      }
    )();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Navigation isJWTCorrect={isJWTCorrect} setIsJWTCorrect={setIsJWTCorrect}/>

        <main className="form-signin">
          <Route exact path="/" component={() => <Home isJWTCorrect={isJWTCorrect}/>}/>
          <Route path="/login" component={() => <Login isJWTCorrect={isJWTCorrect} setIsJWTCorrect={setIsJWTCorrect}/>}/>
          <Route path="/register" component={() => <Register isJWTCorrect={isJWTCorrect}/>}/>
          <Route path="/profile" component={() => <Profile isJWTCorrect={isJWTCorrect}/>}/>
        </main>
        <Route path="/chat" component={Chat}/>
      </BrowserRouter>
    </div>
  );
}

export default App;
