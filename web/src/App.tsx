import React, {useEffect, useState} from 'react';
import Navbar from "./Components/Navbar";
import {BrowserRouter, Route} from "react-router-dom";
import Register from "./Pages/Register";
import Home from "./Pages/Home";

class App extends React.Component {
  private isJWTConnect = false;

  print(): void {
    console.log(":(");
  }

  render() {
    return (
      <BrowserRouter>
        <Route path="/register" component={() => <Register isJWTConnect={this.isJWTConnect} setIsJWTConnect={this.print} />}/>
      </BrowserRouter>
    )
  }
}


/*const App = () => {
  const [isJWTConnect, setIsJWTConnect] = useState(false);

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://localhost:8000/api/auth/jwtverify', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });

        response.status === 200 ? setIsJWTConnect(true) : setIsJWTConnect(false);
      }
    )();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Navbar isJWTConnect={isJWTConnect} setIsJWTConnect={setIsJWTConnect} />

        <Route exact path="/" component={() => <Home />}/>
        <Route path="/register" component={() => <Register isJWTConnect={isJWTConnect} setIsJWTConnect={setIsJWTConnect} />}/>
      </BrowserRouter>
    </div>
  );
};*/

export default App;