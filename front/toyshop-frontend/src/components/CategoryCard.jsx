import React from 'react';

const CategoryCard = ({ category, onClick }) => {
  return (
    <div className="category-card" onClick={onClick}>
      <h3>{category.name}</h3>
      <img src={category.image} alt={category.name} />
    </div>
  );
};

export default CategoryCard;
