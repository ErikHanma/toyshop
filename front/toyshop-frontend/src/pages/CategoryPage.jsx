import React, { useEffect, useState } from 'react';
import CategoryCard from '../components/CategoryCard';
import axios from 'axios';

const CategoryPage = () => {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:5000/api/categories')
      .then(response => setCategories(response.data))
      .catch(error => console.error('Error fetching categories:', error));
  }, []);

  return (
    <div className="category-page">
      <h1>Категории игрушек</h1>
      <div className="category-list">
        {categories.map(category => (
          <CategoryCard key={category.id} category={category} />
        ))}
      </div>
    </div>
  );
};

export default CategoryPage;
