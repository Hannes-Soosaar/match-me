import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './InterestSection.css';

const InterestSection = () => {

const authToken = localStorage.getItem('token');

const [categories, setCategories] = useState([]);
const [userInterests, setUserInterests] = useState([]);



// TODO: do not reload the page after adding an interest, swap the color but update on save and close. 

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
  return (
        <div className='interest-section' > 
        <p className='heading'>Select your interest and matching parameters</p>
            {categories.map((category) => (
              <>
                <div className ='category-section' key={category.category_id}>
                <p className='title-section' key={category.category_id}>{category.category}</p>
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
                <br/>
                </>
              ))}
        </div>
  );

}
export default InterestSection;