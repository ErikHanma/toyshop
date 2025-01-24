import React, { useState } from 'react';

const CheckoutPage = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    address: '',
    paymentMethod: 'card',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Order submitted:', formData);
    alert('Спасибо за заказ!');
  };

  return (
    <div className="checkout-page">
      <h1>Оформление заказа</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="name"
          placeholder="Ваше имя"
          value={formData.name}
          onChange={handleInputChange}
        />
        <input
          type="email"
          name="email"
          placeholder="Ваш email"
          value={formData.email}
          onChange={handleInputChange}
        />
        <input
          type="text"
          name="address"
          placeholder="Адрес доставки"
          value={formData.address}
          onChange={handleInputChange}
        />
        <select name="paymentMethod" onChange={handleInputChange}>
          <option value="card">Банковская карта</option>
          <option value="cash">Наличные</option>
        </select>
        <button type="submit" className="btn-primary">Подтвердить заказ</button>
      </form>
    </div>
  );
};

export default CheckoutPage;
