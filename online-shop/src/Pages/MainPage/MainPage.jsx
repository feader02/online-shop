import React, { useState, useEffect } from 'react';
import axios from 'axios';

const MainPage = () => {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                const response = await axios.get('http://localhost:8000/api/products?price_radius=250-250');
                setProducts(response.data);
            } catch (error) {
                console.error('Error fetching products:', error);
            }
        };

        fetchProducts();
    }, []);

    return (
        <div>
            <h1>Список продуктів</h1>
            <ul>
                {products.map(product => (
                    <li key={product.id}>
                        <h2>{product.name}</h2>
                        <p>Ціна: {product.price}</p>
                        <p>Опис: {product.description}</p>
                        <p>Розміри: {product.height}x{product.width}x{product.depth}</p>
                        <img src={product.photo} alt={product.name} />
                        <p>Тип: {product.type}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MainPage;
