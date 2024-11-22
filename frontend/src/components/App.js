import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';

import Header from './Header/Header';
import Landing from './Landing/Landing';
import LoginSignup from './LoginSignup/LoginSignup';

function App() {

  const isAuthenticated = !!localStorage.getItem('token');

  return (
    <BrowserRouter>
      <div>
        <Header />
        <Routes>
          <Route exact path='/' element={<Landing />} />
          <Route
            path='/login'
            element={!isAuthenticated ? <LoginSignup /> : <Navigate to="/" />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
