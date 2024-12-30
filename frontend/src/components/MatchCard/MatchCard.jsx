import React from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';

const MatchCard = ({ userProfile }) => {
    // update this match the data we are sending
    const { imagePath, firstName, lastName, location, matchScore, description } = userProfile;

    return (
        <>
            <div className="match-card">
                <img className="match-card-image" src={defaultProfilePic}></img>
                <h2>{`${firstName} ${lastName}`}</h2>
                <p>{location}</p>
                <p>{matchScore}</p>
                <p>{description}</p>
            </div>
        </>
    )
}

export default MatchCard