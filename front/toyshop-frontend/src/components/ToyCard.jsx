import React from 'react';
import { Link } from 'react-router-dom';

const ToyCard = ({ toy, addToCart }) => {
  return (
    <div className="toy-card">
      <img src={toy.image} alt={toy.name} />
      <h3>{toy.name}</h3>
      <p>Цена: {toy.price}₽</p>
      <button className="btn-primary" onClick={() => addToCart(toy)}>
        Добавить в корзину
      </button>
      <Link to={`/product/${toy.id}`} className="btn-secondary">Подробнее</Link>
    </div>
  );
};

export default ToyCard;
