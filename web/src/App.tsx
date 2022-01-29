import React from 'react';
import './App.css';
import Login from "./pages/Login";
import Navigation from "./components/Navigation";
import {BrowserRouter, Route} from "react-router-dom";
import Register from "./pages/Register";
import Home from "./pages/Home";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Navigation/>

        <main className="form-signin">
          <Route exact path="/" component={Home}/>
          <Route path="/login" component={Login}/>
          <Route path="/register" component={Register}/>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
