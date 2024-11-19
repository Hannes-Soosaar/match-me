import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Header from './Header/Header';
import Landing from './Landing/Landing';
import LoginSignup from './LoginSignup/LoginSignup';

function App() {

  return (
    <div>
      <BrowserRouter>
        <div>
          <Header />
          <Routes>
            <Route exact path='/' element={<Landing />} />
            <Route path='/login' element={<LoginSignup />} />
          </Routes>
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
