import React from 'react';
import {Link} from "react-router-dom";

const Navigation = (props: { isJWTCorrect: boolean, setIsJWTCorrect: (isJWTCorrect: boolean) => void }) => {
  let menu;

  const logout = async () => {
    await fetch('http://localhost:8000/api/auth/logout', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
    });

    props.setIsJWTCorrect(false);
  }

  if (props.isJWTCorrect) {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        <li className="nav-item active">
          <Link to="#" className="nav-link">Chat</Link>
        </li>
        <li className="nav-item active">
          <Link to="/profile" className="nav-link">Profile</Link>
        </li>
        <li className="nav-item active">
          <Link to="/login" className="nav-link" onClick={logout}>Logout</Link>
        </li>
      </ul>
    )
  } else {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        <li className="nav-item">
          <Link to="/login" className="nav-link">Login</Link>
        </li>
        <li className="nav-item">
          <Link to="/register" className="nav-link">Register</Link>
        </li>
      </ul>
    )
  }


  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <Link to="/" className="navbar-brand" href="#">Home</Link>

        <div>
          {menu}
        </div>
      </div>
    </nav>
  );
};

export default Navigation;