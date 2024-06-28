import React from 'react'
import "./reset.scss"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/header/header'
import './Pages/ProductPage/ProductPage.scss';
import './App.css'
import ProductPage from "./Pages/ProductPage/ProductPage.jsx";

function App() {
    return (
        <>
            <Router>
                <Header />
                <Routes>
                    <Route path="/product/:id" element={<ProductPage />} />

                </Routes>
            </Router>
        </>
    )
}

export default App
