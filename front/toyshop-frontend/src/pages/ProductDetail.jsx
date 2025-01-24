import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';

const ProductDetail = () => {
  const { id } = useParams(); // Извлекаем ID из параметров маршрута
  const [toy, setToy] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:5000/api/toys/${id}`)
      .then(response => setToy(response.data))
      .catch(error => console.error('Error fetching toy details:', error));
  }, [id]);

  if (!toy) return <p>Загружается...</p>;

  return (
    <div className="product-detail">
      <img src={toy.image} alt={toy.name} />
      <h2>{toy.name}</h2>
      <p>{toy.description}</p>
      <p>Цена: {toy.price}₽</p>
      <button>Добавить в корзину</button>
    </div>
  );
};

export default ProductDetail;
