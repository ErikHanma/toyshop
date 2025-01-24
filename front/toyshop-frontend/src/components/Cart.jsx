import React from 'react';

const Cart = ({ cartItems = [], removeFromCart }) => {
  return (
    <div className="cart">
      <h2>Корзина</h2>
      {cartItems.length === 0 ? (
        <p>Ваша корзина пуста.</p>
      ) : (
        cartItems.map(item => (
          <div key={item.id} className="cart-item">
            <img src={item.image} alt={item.name} />
            <div>
              <h3>{item.name}</h3>
              <p>{item.price}₽</p>
              <button onClick={() => removeFromCart(item.id)}>Удалить</button>
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default Cart;
