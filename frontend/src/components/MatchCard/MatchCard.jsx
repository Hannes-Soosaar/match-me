import React from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import axios from 'axios';


const authToken = localStorage.getItem('token');

const MatchCard = ({ userProfile }) => {

    const { match_id, match_score, matched_user_name, matched_user_picture, status } = userProfile;  

    // Set the default profile picture if no picture is provided
    const userProfilePic = matched_user_picture ? matched_user_picture : defaultProfilePic;



    //TODO build the logic on which buttons are shown based on the status of the match

    const handleRemoveMatch = async () => {
        try {
            const response = await axios.put('/matches/remove', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match Remove:', response.data);

        } catch (error) {
            console.error('Error removing the match:', error);
        }
    };


    const handleConnectMatch = async () => {
        try {
            const response = await axios.put('/matches/connect', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match accepted:', response.data);
        } catch (error) {
            console.error('Error connecting to user:', error);
        }
    };


    const handleRequestMatch = async () => {
        try {
            const response = await axios.put('/matches/request', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Requested to match:', response.data);
            // You can implement additional logic like updating the UI or showing a success message
        } catch (error) {
            console.error('Error requesting to connect:', error);
        }
    };

    const handleBlockMatch = async () => {
        try {
            const response = await axios.put('/matches/block', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match Blocked:', response.data);
            // You can implement additional logic like updating the UI or showing a success message
        } catch (error) {
            console.error('Error blocking user:', error);
        }
    };



    return (
        <>
            <div className="match-card">
                <img className="match-card-image" src={userProfilePic}></img>
                <h2></h2>
                <h3>MatchId:</h3>
                <p>{match_id}</p>
                <h3>MatchScore:</h3>
                <p>{match_score}</p>
                <h3>Name:</h3>
                <p>{matched_user_name}</p>
                <h3>Status:</h3>
                <p>{status}</p>
                <p>----</p>
                <div className="match-card-buttons">

                <button onClick={handleConnectMatch} className="match-card-button">
                    Connect
                </button>
                <button onClick={handleRemoveMatch} className="match-card-button">
                    Delete
                </button>
                <button onClick={handleRequestMatch} className="match-card-button">
                    Request
                </button>
                <button onClick={handleBlockMatch} className="match-card-button">
                    Block User
                </button>
            </div>

            </div>
        </>
    )
}

export default MatchCard