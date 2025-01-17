import React from "react";
import {useState, useEffect} from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import axios from 'axios';
import Modal from '../Modal/Modal.jsx';
import BuddyCard from "../BuddyCard/BuddyCard.jsx";


const authToken = localStorage.getItem('token');

const MatchCard = ({ userProfile, onUpdate }) => {

    const { match_id,
        match_score,
        status,
        is_online,
        matched_user_id,
        matched_user_name, 
        matched_user_picture,
        matched_user_description, 
        matched_user_location } = userProfile;  

    const basePictureURL = "http://localhost:4000/uploads/";
    const onlineURL = "/images/OnlineIconPNG.png"
    const offlineURL = "/images/OfflineIconPNG.png"
    // Set the default profile picture if no picture is provided
    let userProfilePic = matched_user_picture ? matched_user_picture : defaultProfilePic;

    if (userProfilePic !== defaultProfilePic) {
        userProfilePic = basePictureURL + userProfilePic;
    }

    const [isModalOpen, setModalOpen] = useState(false);

    const handleViewMatchedProfile = () => {
        setModalOpen(true);
    };

    const handleCloseModal = () => {
        setModalOpen(false);
    };

    const handleRemoveMatch = async () => {
        try {
            const response = await axios.put('/matches/remove', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match Remove:', response.data);
            onUpdate(match_id);
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
            onUpdate(match_id);
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
            onUpdate(match_id);
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
            onUpdate(match_id);
            // You can implement additional logic like updating the UI or showing a success message
        } catch (error) {
            console.error('Error blocking user:', error);
        }
    };


    const renderButtons = () => {
        switch (status) {
            case 'new':
                return (
                    <>
                        <button onClick={handleRequestMatch} className="match-card-button">
                            Request
                        </button>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
                        </button>
                    </>
                );
            case 'requested':
                return (
                    <>
                        <button onClick={handleConnectMatch} className="match-card-button">
                            Connect
                        </button>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Delete
                        </button>
                    </>
                );
            case 'blocked':
                return (
                    <>
                        <p>
                            You are not authorized to contact this user.
                        </p>
                    </>
                );
            default:
                return (
                    <>
                        <button onClick={handleRequestMatch} className="match-card-button">
                            Request
                        </button>
                        
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
                        </button>
                    </>
                );
        }
    };

    return (
        <>
            <div className="match-card">
                <div className="match-card-info">
                    <img className="match-card-image" src={userProfilePic} alt ="User"></img>
                    <h2>{matched_user_location }</h2>
                    <h3>MatchScore:</h3>
                    <p>{match_score}</p>
                    <h3>Name:</h3>
                    <p>{matched_user_name}</p>
                </div>

                <div className="match-card-buttons">
                    {renderButtons()}   
                    <button onClick={handleViewMatchedProfile} className="match-card-button">
                        View Profile
                    </button>
                    {is_online ? <p>Online</p> : <p>Offline</p>}
                </div>
            </div>

            <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                <BuddyCard buddyProfile={userProfile} />
            </Modal>
        </>
    )
}

export default MatchCard