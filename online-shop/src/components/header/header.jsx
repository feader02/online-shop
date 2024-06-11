// components/Header.js
import React from 'react';
import './header.scss';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header className="header">
      <div className="header-icons">
        <div className="icon black-square"></div>
        <div className="icon black-square"></div>
        <div className="icon black-square"></div>
      </div>
      <h1 className="header-title">Avion</h1>
      
      <nav className="header-nav">
        <ul>
          <li><Link to="/">Plant pots</Link></li>
          <li><Link to="/">Ceramics</Link></li>
          <li><Link to="/">Tables</Link></li>
          <li><Link to="/">Chairs</Link></li>
          <li><Link to="/">Crockery</Link></li>
          <li><Link to="/">Tableware</Link></li>
          <li><Link to="/">Cutlery</Link></li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
