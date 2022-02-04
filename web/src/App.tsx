import React, {useState} from 'react';
import Navbar from "./Components/Navbar";
import {BrowserRouter} from "react-router-dom";

const App = () => {
  const [isJWTConnect, setIsJWTConnect] = useState(false);



  return (
    <div className="App">
      <BrowserRouter>
        <Navbar isJWTConnect={isJWTConnect} setIsJWTConnect={setIsJWTConnect} />
      </BrowserRouter>
    </div>
  );
};

export default App;