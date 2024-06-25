import React from 'react'
import "./reset.scss"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/header/header'
import MainPage from './Pages/MainPage/MainPage';
import './App.css'

function App() {


  return (
    <>
     <Router>
      <Routes>
        <Route path="/" element={<MainPage />} />
        
      </Routes>
    </Router>
    </>
  )
}

export default App
