import React from "react";
import {useState, useEffect} from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import axios from 'axios';
import Modal from '../Modal/Modal.jsx';
import BuddyCard from "../BuddyCard/BuddyCard.jsx";


const authToken = localStorage.getItem('token');

const MatchCard = ({ userProfile }) => {

    const { match_id, match_score,status,matched_user_id, matched_user_name, matched_user_picture,matched_user_description, matched_user_location } = userProfile;  

    // Set the default profile picture if no picture is provided
    // const userProfilePic = matched_user_picture ? matched_user_picture : defaultProfilePic;
    const userProfilePic = defaultProfilePic;
    const [isModalOpen, setModalOpen] = useState(false);

    const handleViewMatchedProfile = () => {
        setModalOpen(true);
    };

    const handleCloseModal = () => {
        setModalOpen(false);
    };



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
                <div className="match-card-info">
                    <img className="match-card-image" src={userProfilePic}></img>
                    <h2>{matched_user_location }</h2>
                    <h3>MatchScore:</h3>
                    <p>{match_score}</p>
                    <h3>Name:</h3>
                    <p>{matched_user_name}</p>
                </div>

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
                    <button onClick={handleViewMatchedProfile} className="match-card-button">
                        view Profile
                    </button>
                </div>
            </div>

            <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                <BuddyCard buddyProfile={userProfile} />
            </Modal>
        </>
    )
}

export default MatchCard