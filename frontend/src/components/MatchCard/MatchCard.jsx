import React from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';

const MatchCard = ({ userProfile }) => {


    // const [requestMatch, setRequestMatch] = useState([])
    // const [confirmMatch, setConfirmMatch] = useState([])


    // update this match the data we are sending
    const { picture,userName, matchId, location, userScore  } = userProfile;

    return (
        <>
            <div className="match-card">
                <img className="match-card-image" src={defaultProfilePic}></img>
                <h2></h2>
                <p>{location}</p>
                <p></p>
                <p></p>
            </div>
        </>
    )
}

export default MatchCard