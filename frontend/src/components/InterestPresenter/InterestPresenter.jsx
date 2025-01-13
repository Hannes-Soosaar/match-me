import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './InterestPresenter.css';

const InterestPresenter = () => {
  const [bioValues, setBioValues] = useState([]);
  const authToken = localStorage.getItem('token');

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('/me/bio', {
          headers: {
            Authorization: `Bearer ${authToken}`,
          },
        });
        console.log('Fetched data:', response.data);

        // Extract and store only the values from the response
        const values = Object.values(response.data);
        setBioValues(values);
      } catch (error) {
        console.error('Error fetching bio data:', error);
      }
    };

    fetchData();
  }, [authToken]);

  return (
    <div className="interest-section">
      <p className="heading">Your interests:</p>
      {bioValues.length > 0 ? (
        <div className="interests-container">
          {bioValues.map((value, index) => (
            <button key={index} className="selected">
              {value}
            </button>
          ))}
        </div>
      ) : (
        <p>Choose your interests</p>
      )}
    </div>
  );
  
};

export default InterestPresenter;
