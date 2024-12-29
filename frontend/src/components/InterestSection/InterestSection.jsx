import React, { useState, useEffect } from 'react';
import axios from 'axios';

const InterestSection = () => {

const authToken = localStorage.getItem('token');

const [categories, setCategories] = useState([]);
const [userInterests, setUserInterests] = useState([]);

useEffect(() => {
// If the user is not logged in send only the interests for FE population if the user in logged in send the interests and the user interests
    const fetchData = async () => {
        try {
            const response1 = await axios.get('/interests', {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log(response1.data)
            setCategories(response1.data);
        } catch (error) {
            console.error('Error fetching Interest data:', error);
        }
       

// Get the user Interest data
        try{
            const response2 = await axios.get('/userInterests', {
                headers: {
                    // Get the user id from the authToken
                    Authorization: `Bearer ${authToken}`,
                },
            });
            if (response2.data.length !== 0) {
                setUserInterests(response2.data);
            }
           
        }
        catch (error) {
            console.error('Error fetching UserInterest data:', error);
        }
    };

    fetchData();

}, [authToken]);


const handleInterestClick = (interestId) => {
    console.log(`Interest ID clicked: ${interestId}`);
  };

return (
    <div>
      <h1>Interest Section</h1>
      <div>
        <h2>Categories</h2>
        <ul>
          {categories.map((category) => (
            <li key={category.category_id}>
              <h3>{category.category}</h3>
              <div>
                {category.interests.map((interest) => (
                  <button
                    key={interest.id}
                    onClick={() => handleInterestClick(interest.id)}
                  >
                    {interest.interestName}
                  </button>
                ))}
              </div>
            </li>
          ))}
        </ul>
        <h2>User Interests</h2>
        <ul>
          {userInterests.map((userInterest) => (
            <li key={userInterest.interestId}>{userInterest.interestId}</li>
          ))}
        </ul>
      </div>
    </div>
);

}
export default InterestSection;