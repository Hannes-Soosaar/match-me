import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import { WebSocketProvider } from './WebSocketContext';
import { useState, useEffect } from 'react';

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
  const [profileExists, setProfileExists] = useState(localStorage.getItem('profileExists'));

  useEffect(() => {
    // Initialize on mount
    setProfileExists(localStorage.getItem('profileExists') === 'true');
  }, []);

  return (
    <WebSocketProvider>
      <BrowserRouter>
        <div>
          <Header />
          <Routes>
            <Route exact
              path='/'
              element={isAuthenticated && !profileExists ? <Navigate to="/profile" /> : <Landing />} />
            <Route
              path='/login'
              element={!isAuthenticated ? <LoginSignup /> : <Navigate to="/" />} />
            <Route
              path='/dashboard'
              element={isAuthenticated && profileExists ? <Dashboard /> : <Navigate to="/profile" />} />
            <Route
              path='/profile'
              element={isAuthenticated ? <Profile /> : <Navigate to="/" />} />
            <Route
              path='/matches'
              element={isAuthenticated && profileExists ? <Matches /> : <Navigate to="/" />} />
            <Route
              path='/connections'
              element={isAuthenticated && profileExists ? <BuddiesSection /> : <Navigate to="/" />} />
            <Route
              path='/chat'
              element={isAuthenticated && profileExists ? <Chat /> : <Navigate to="/" />} />
          </Routes>
        </div>
      </BrowserRouter >
    </WebSocketProvider>
  );
}

export default App;
