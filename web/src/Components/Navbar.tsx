import React from 'react';
import {Link} from "react-router-dom";
import {IJWTInterface} from "../Interfaces/IJWTInterface";

class Navbar extends React.Component<IJWTInterface> {
  constructor(props: IJWTInterface) {
    super(props);
  }

  logout = async () => {
    await fetch('http://localhost:8000/api/auth/logout', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
    });

    this.props.setIsJWTConnect(false);
  }

  getMenu() : JSX.Element {
    if (this.props.isJWTConnect) {
      return (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
          <li className="nav-item active">
            <Link to="/chat" className="nav-link">Chat</Link>
          </li>
          <li className="nav-item active">
            <Link to="/profile" className="nav-link">Profile</Link>
          </li>
          <li className="nav-item active">
            <Link to="/login" className="nav-link" onClick={this.logout}>Logout</Link>
          </li>
        </ul>
      )
    }
    return (
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

  render() {
    return (
      <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand" href="#">Home</Link>
          <div>
            {this.getMenu()}
          </div>
        </div>
      </nav>
    );
  }
}

export default Navbar;