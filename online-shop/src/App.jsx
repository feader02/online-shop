import React from 'react'
import "./reset.scss"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/header/header'
import MainPage from './Pages/MainPage/MainPage';
import ChairPage from './Pages/cheir/chair';
import './App.css'

function App() {


  return (
    <>
     <Router>
      <Header />
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/chair" element={<ChairPage />} />
      </Routes>
    </Router>
    </>
  )
}

export default App
