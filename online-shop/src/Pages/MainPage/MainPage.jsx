// components/MainPage.js
import React from 'react';
import { Link } from 'react-router-dom';

const MainPage = () => {
  return (
    <div>
      <h1>Main Page</h1>
      <Link to="/chair">Chair</Link>
    </div>
  );
};

export default MainPage;


