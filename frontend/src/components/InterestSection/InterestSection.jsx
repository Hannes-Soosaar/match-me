import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './InterestSection.css';

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

    try {
        axios.post('/userInterest', {
            interestId,
        },
        {
            headers: {
                Authorization: `Bearer ${authToken}`,
            },
        })  .then((response) => {
            window.location.reload(); 
        }).catch((error) => {
            console.error('Error adding interest:', error);
        });
    } catch (error) {
        console.error('Error adding interest:', error);
    }

  };


//TODO extract the button element to a separate component
//TODO have different styling for the buttons that have the same interestId as the user already selected.
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
                    className={userInterests.includes(interest.id) ? 'selected' : 'unselected'}
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
            <li key={userInterest}>{userInterest}</li>
          ))}
        </ul>
      </div>
    </div>
);

}
export default InterestSection;