import React, { useEffect, useState } from 'react';
import axios from 'axios';
import ToyCard from './ToyCard';  // Импортируем компонент карточки игрушки

const ToyList = () => {
  const [toys, setToys] = useState([]);

  useEffect(() => {
    // Здесь идет запрос к вашему API для получения списка игрушек
    axios.get('http://localhost:5000/api/toys')  // Поменяйте URL на тот, который используется на вашем бэкенде
      .then(response => setToys(response.data))
      .catch(error => console.error('Error fetching toys:', error));
  }, []);

  return (
    <div className="toy-list">
      {toys.length === 0 ? (
        <p>Нет доступных игрушек.</p>
      ) : (
        toys.map(toy => (
          <ToyCard key={toy.id} toy={toy} />
        ))
      )}
    </div>
  );
};

export default ToyList;
