import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import GetProducts from "../../components/axios/Axios.jsx";
import Delivery from "./images/Delivery.svg";
import Checkmark from "./images/Checkmark--outline.svg";
import Purchase from "./images/Purchase.svg";
import Sprout from "./images/Sprout.svg";


function ProductPage() {
    const { id } = useParams();
    const [product, setProduct] = useState(null);
    const [otherProducts, setOtherProducts] = useState([]);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const productData = await GetProducts(`/api/products/${id}`);
                setProduct(productData);

                const allProductsData = await GetProducts(`/api/products?page_size=5`);

                const filteredProducts = allProductsData.filter(p => p.id.toString() !== id);

                if (filteredProducts.length > 4) {
                    filteredProducts.pop();
                }

                setOtherProducts(filteredProducts);
            } catch (error) {
                console.error('Error fetching product:', error);
            }
        };

        fetchData();
    }, [id]);

    if (!product) {
        return <div>Loading...</div>;
    }

    return (
        <>
            <section className="product">
                <img src={product.photo} alt={product.name} className="product-img"/>
                <div className="product-info">
                    <h1 className="product-title">{product.name}</h1>
                    <p className="product-price">£{product.price}</p>
                    <article className="product-description-block">
                        <p className="product-description-static">Description</p>
                        <p className="product-description">{product.description}</p>
                    </article>
                    <article className="product-dimensions-block">
                        <p className="product-dimensions">Dimensions</p>
                        <dl>
                            <div className="product-height-block">
                                <dt className="product-height-static">Height</dt>
                                <dd className="product-height">{product.height}cm</dd>
                            </div>
                            <div className="product-width-block">
                                <dt className="product-width-static">Width</dt>
                                <dd className="product-width">{product.width}cm</dd>
                            </div>
                            <div className="product-depth-block">
                                <dt className="product-depth-static">Depth</dt>
                                <dd className="product-depth">{product.depth}cm</dd>
                            </div>
                        </dl>
                    </article>
                </div>
            </section>
            <section className="products-list-section">
                <h1 className="product-list-main-h1">You might also like</h1>
                <div className="products-list">
                    {otherProducts.map((otherProduct) => (
                        <a href={`/product/${otherProduct.id}`} className="products-list-block">
                            <img src={otherProduct.photo} alt={otherProduct.name} className="products-list-img"/>
                            <h1 className="products-list-title">{otherProduct.name}</h1>
                            <p className="products-list-price">£{otherProduct.price}</p>
                        </a>
                    ))}
                </div>
                <a href="/all-products" className="all-products-button">View collection</a>
            </section>
        </>
    )
}

export default ProductPage