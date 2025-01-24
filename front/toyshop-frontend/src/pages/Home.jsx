import React from 'react';
import ToyList from '../components/ToyList';

const Home = () => {
  return (
    <div>
      {/* Главный баннер */}
      <section className="hero-banner">
        <h1>Добро пожаловать в магазин игрушек!</h1>
        <p>Находите любимые игрушки для своих детей</p>
        <button className="btn-primary">Смотреть коллекцию</button>
      </section>

      {/* Популярные товары */}
      <section className="popular-toys">
        <h2>Популярные товары</h2>
        <ToyList />
      </section>

      {/* Категории */}
      <section className="categories">
        <h2>Категории игрушек</h2>
        <div className="category-list">
          <div className="category-card">Для малышей</div>
          <div className="category-card">Конструкторы</div>
          <div className="category-card">Мягкие игрушки</div>
          <div className="category-card">Образовательные</div>
        </div>
      </section>

      {/* Отзывы */}
      <section className="reviews">
        <h2>Отзывы наших клиентов</h2>
        <div className="reviews-container">
          <div className="review-card">
            <p>"Игрушки отличного качества, дети в восторге!"</p>
            <span>- Екатерина</span>
          </div>
          <div className="review-card">
            <p>"Большой выбор и быстрая доставка. Рекомендую!"</p>
            <span>- Иван</span>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Home;
