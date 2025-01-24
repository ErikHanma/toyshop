import React from 'react';
import { Link } from 'react-router-dom';

const ResetPasswordPage = () => {
  return (
    <div className="auth-container">
      <form className="auth-form">
        <h2>Reset Password</h2>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" placeholder="Enter your email to reset password" required />
        </div>
        <button type="submit" className="btn-primary">Reset Password</button>
        <p className="auth-helper">
          Remembered your password? <Link to="/login">Login</Link>
        </p>
      </form>
    </div>
  );
};

export default ResetPasswordPage;
