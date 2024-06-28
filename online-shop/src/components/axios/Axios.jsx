import axios from 'axios';

async function GetProducts(apiUrl) {
    try {
        const response = await axios.get("http://localhost:8000" + apiUrl);
        return response.data;
    } catch (error) {
        console.error('Error fetching products:', error);
        return [];
    }
}

export default GetProducts