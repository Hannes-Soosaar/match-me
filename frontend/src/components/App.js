import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import { WebSocketProvider } from './WebSocketContext';

import Header from './Header/Header';
import Landing from './Landing/Landing';
import LoginSignup from './LoginSignup/LoginSignup';
import Dashboard from './Dashboard/Dashboard';
import Profile from './Profile/Profile';
import Matches from './Matches/Matches';
import BuddiesSection from './BuddiesSection/BuddiesSection.jsx';
import Chat from './Chat/Chat';


function App() {

  const isAuthenticated = !!localStorage.getItem('token');
  const profileNotExist = localStorage.getItem('profileExists') !== 'doesNotExist';

  return (
    <WebSocketProvider>
      <BrowserRouter>
        <div>
          <Header />
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
              element={isAuthenticated ? <Matches /> : <Navigate to="/" />} />
            <Route
              path='/connections'
              element={isAuthenticated ? <BuddiesSection /> : <Navigate to="/" />} />
            <Route
              path='/chat'
              element={isAuthenticated ? <Chat /> : <Navigate to="/" />} />
          </Routes>
        </div>
      </BrowserRouter >
    </WebSocketProvider>
  );
}

export default App;
