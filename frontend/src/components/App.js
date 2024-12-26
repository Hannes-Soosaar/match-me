import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';

import Header from './Header/Header';
import Landing from './Landing/Landing';
import LoginSignup from './LoginSignup/LoginSignup';
import Dashboard from './Dashboard/Dashboard';
import Profile from './Profile/Profile';
import Matches from './Matches/Matches';


function App() {

  const isAuthenticated = !!localStorage.getItem('token');
  const profileNotExist = localStorage.getItem('profileExists') !== 'doesNotExist';

  return (
    <BrowserRouter>
      <div>
        <Header/>
        <Routes>
          <Route exact
            path='/'
            element={isAuthenticated && profileNotExist ? <Navigate to="/profile" /> : <Landing />} />
          <Route
            path='/login'
            element={!isAuthenticated ? <LoginSignup /> : <Navigate to="/" />} />
          <Route
            path='/dashboard'
            element={isAuthenticated ? <Dashboard /> : <Navigate to="/" />} />
          <Route
            path='/profile'
            element={isAuthenticated ? <Profile /> : <Navigate to="/" />} />
          <Route
            path='/matches'
            element={isAuthenticated ? <Matches/> : <Navigate to="/" />} />
        </Routes>
      </div>
    </BrowserRouter >
  );
}

export default App;
