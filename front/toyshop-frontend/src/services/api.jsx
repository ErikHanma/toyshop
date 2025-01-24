import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:5000/api', // Здесь укажите URL вашего backend
});

export default api;
