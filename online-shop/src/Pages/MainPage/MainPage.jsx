// components/MainPage.js
import React from 'react';
import "./MainPage.scss"
import { Link } from 'react-router-dom';
import delivery from './images/delivery.png'
import checkmark from './images/checkmark.png'
import purchase from './images/purchase.png'
import sprout from './images/sprout.png'


const MainPage = () => {
  return (
    <div className='container'>
      <div className='mainpage'>
        <a className='collection' Link to = "/collection">Viev collection</a>
      </div>
      <div className='div-panel'>
        <h1 className='div-panel__title'>
        What makes our brand different
        </h1>
        <div className='panel'>
        <div className='features'>
          <img src={delivery} alt="" />
          <h1>
          Next day as standard
          </h1>
          <p>
          Order before 3pm and get your order
          the next day as standard
          </p>
        </div>
        <div  className='features'>
          <img src={checkmark} alt="" />
          <h1>
          Made by true artisans
          </h1>
          <p>
          Handmade crafted goods made with
real passion and craftmanship
          </p>
        </div>
        <div  className='features'>
          <img src={purchase} alt="" />
          <h1>
          Unbeatable prices
          </h1>
          <p>
          For our materials and quality you won’t find better prices anywhere
          </p>
        </div>
        <div  className='features'>
          <img src={sprout} alt="" />
          <h1>
          Recycled packaging
          </h1>
          <p>
          We use 100% recycled packaging to ensure our footprint is manageable
          </p>
        </div>
      </div>
      </div>
      <div className='new-ceramics'>
        <h1>
        New ceramics
        </h1>
      </div>
    </div>
   
  );
};

export default MainPage;


