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
            setUserInterests(response2.data);
        }
        catch (error) {
            console.error('Error fetching UserInterest data:', error);
        }
    };

    fetchData();

}, [authToken]);

return (
    <div>
        <h1>Interest Section</h1>
        <div>
            <h2>Categories</h2>
            <ul>
                {categories.map((category) => (
                    // Sub this with the Interest Button component 
                    <li key={category.id}>{category.name} : {category.interest}</li>
                ))}
            </ul>  
            <h2>User Interests</h2>
            <ul>
                {userInterests.map((userInterest) => (
                    <li key={userInterest.id}>{userInterest.interestId}</li>
                ))} 
            </ul>
    </div>
    </div>
);

}
export default InterestSection;