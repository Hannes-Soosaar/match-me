import React, { useEffect, useState, useContext } from 'react';
import './BuddyCard.css';
import axios from 'axios';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import Modal from '../Modal/Modal.jsx';
import Chat from '../Chat/Chat.jsx';


const authToken = localStorage.getItem('token');

// The buddy card take in the data from the the /matches API and get the interest from the /interests API
const BuddyCard = ({ buddyProfile, onUpdate }) => {




    const { match_id,
        match_score,
        status,
        is_online,
        requester,
        matched_user_id,
        matched_user_name, 
        matched_user_picture,
        matched_user_description, 
        matched_user_location } = buddyProfile;  

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
        window.location.reload()
    };

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

    
    const renderButtons = () => {
        switch (status) {
            case 'connected':
                return (
                    <>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
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
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Un-Buddy
                        </button>
                    </>
                );
        }
    };


    
    return (
        <>
            <div className="match-card">
                <div className="match-card-status">
                <div className="user-name" >{matched_user_name}</div>  
                {is_online ? <img src={onlineURL} alt="User online" className="status-icon"></img>
                    :
                <img src={offlineURL} alt="User offline" className="status-icon"></img>
                }
                </div>

                <div className="match-card-info">
                    <img className="match-card-image" src={userProfilePic} alt ="User"></img>

                    <h2>{matched_user_location }</h2>
                    <h3>MatchScore:</h3>
                    <p>{match_score}</p>
                    
                </div>

                <div className="match-card-buttons">
                    {renderButtons()}   
                    <button onClick={handleViewMatchedProfile} className="match-card-button">
                        Chat
                    </button>
                </div>
            </div>
            <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                <Chat></Chat>
            </Modal>
        </>
);

}

export default BuddyCard;  