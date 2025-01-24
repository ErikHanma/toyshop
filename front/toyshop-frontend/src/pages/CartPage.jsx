import React, { useState } from 'react';
import Cart from '../components/Cart';

const CartPage = () => {
  const [cartItems, setCartItems] = useState([]); // Пример данных корзины

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter(item => item.id !== id));
  };

  return (
    <div>
      <h1>Ваша корзина</h1>
      <Cart cartItems={cartItems} removeFromCart={removeFromCart} />
    </div>
  );
};

export default CartPage;
