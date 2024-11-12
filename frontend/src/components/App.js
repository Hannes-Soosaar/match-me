import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Header from './Header';
import Landing from './Landing';

function App() {

  return (
    <div>
      <BrowserRouter>
        <div>
          <Header />
          <Routes>
            <Route exact path='/' element={<Landing />} />
          </Routes>
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
