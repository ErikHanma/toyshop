import React from 'react';
import { Link } from 'react-router-dom';

const LoginPage = () => {
  return (
    <div className="auth-container">
      <form className="auth-form">
        <h2>Login</h2>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" placeholder="Enter your email" required />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input type="password" id="password" placeholder="Enter your password" required />
        </div>
        <button type="submit" className="btn-primary">Login</button>
        <p className="auth-helper">
          Don't have an account? <Link to="/register">Sign Up</Link>
        </p>
        <p className="auth-helper">
          <Link to="/reset-password">Forgot Password?</Link>
        </p>
      </form>
    </div>
  );
};

export default LoginPage;
