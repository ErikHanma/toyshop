import React, { useEffect, useState } from 'react';
import axios from 'axios';
import ToyCard from '../components/ToyCard';

const Recommendations = () => {
  const [recommendedToys, setRecommendedToys] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:5000/api/recommendations')
      .then(response => setRecommendedToys(response.data))
      .catch(error => console.error('Error fetching recommendations:', error));
  }, []);

  return (
    <div>
      <h2>Рекомендации для вас</h2>
      <div className="toy-list">
        {recommendedToys.map(toy => (
          <ToyCard key={toy.id} toy={toy} />
        ))}
      </div>
    </div>
  );
};

export default Recommendations;
