import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header className="header">
      <div className="logo">
        <Link to="/">Магазин игрушек</Link>
      </div>
      <div className="search-bar">
        <input type="text" placeholder="Искать игрушки..." />
        <button>🔍</button>
      </div>
      <nav>
        <ul>
          <li><Link to="/">Главная</Link></li>
          <li><Link to="/cart">Корзина</Link></li>
          <li><Link to="/recommendations">Рекомендации</Link></li>
          <li><Link to="/login">Войти</Link></li>
          <li><Link to="/register">Регистрация</Link></li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
