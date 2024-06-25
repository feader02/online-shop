// components/Header.js
import React from 'react';
import './header.scss';
import { Link } from 'react-router-dom';
import avion from "./images/avion.png"
import search from "./images/search.png"
import line from "./images/line.png"
import user from "./images/User.png"
import cart from "./images/cart.png"




const Header = () => {
  return (
    <header className="header">
      <div className='div_img_header'>
      <img className='title_img' src={search} alt="search" />
      <img src={avion} alt='Avion' className="header-title" />
      <img className='title_img' src={cart} alt="cart" />
      <img className='title_img' src={user} alt="user" />
      </div>
      <img className='line_img' src={line} alt="" />
      
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
