// components/MainPage.js
import React from 'react';
import "./MainPage.scss"
import { Link } from 'react-router-dom';


const MainPage = () => {
  return (
    <div className='mainpage'>
      <a className='collection' Link to = "/collection">Viev collection</a>
    </div>
  );
};

export default MainPage;


