import React, { useEffect, useState, useContext } from 'react';
import './BuddyCard.css';
import axios from 'axios';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';


const authToken = localStorage.getItem('token');

// The buddy card take in the data from the the /matches API and get the interest from the /interests API
const BuddyCard = ({ buddyProfile }) => {

    const { match_id, match_score,status,matched_user_id, matched_user_name, matched_user_picture,matched_user_description, matched_user_location } = buddyProfile;  


    // Temporary, sest the pictrue to be always the defaultProfilePic
       const userProfilePic = defaultProfilePic;



    useEffect(() => {
        // Fetch the interests of the matched user
        const fetchInterests = async () => {
            try {
                const response = await axios.get(`/interests/${matched_user_id}`, {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                console.log('Interests:', response.data);
                // setUserInterests(response.data);
            } catch (error) {
                console.error('Error fetching interests:', error);
            }
        };
        fetchInterests();
    }, [matched_user_id]);

    

/*
Need to have the picture 
Need to have the Username // by default the first half of the email?
Need to have the location
Need to show interests
*/




    return (

<>
        <div className='content-container'>

            <img src={userProfilePic} alt={`${matched_user_name}'s profile`} style={{ width: "100px", borderRadius: "50%" }} />
           
            <div className='user-info'>
                
                <h2>{matched_user_name}</h2>
                <p><strong>Location:</strong> {matched_user_location}</p>
                <p><strong>Description:</strong> {matched_user_description}</p>
            </div>

            <div>
            Interest
            </div>
        </div>

</>

);

}

export default BuddyCard;  