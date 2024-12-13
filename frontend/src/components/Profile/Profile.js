import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import './Profile.css';

const Profile = () => {

    const [username, setUsername] = useState('');
    const [headerText, setHeaderText] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        alert('This feature is WIP')
    }



    useEffect(() => {
        const profileNotExist = localStorage.getItem('profileExists') === 'doesNotExist';
        if (profileNotExist) {
            console.log(profileNotExist)
            setHeaderText("Please complete your profile");
        } else {
            console.log(profileNotExist)
            setHeaderText("Change your profile information if needed");
        }
    }, [])

    return (
        <div style={{ textAlign: 'center' }}>
            <div className='profile-container'>
                {/* 
                        TODO force user to fill out their profile before they have access to anything else.
                        Creating a username and answering biographical information questions is mandatory,
                        setting a profile picture should be possible. If nothing is set, a placeholder image is used instead.
                        "About me" section is optional.
                        */}
                <div className="profile-header">
                    <div className="profile-text">
                        {headerText}
                    </div>
                </div>
                <div className='inputs'>
                    <div className='profile-text'>Choose your username</div>
                    <div className='input'>
                        <input
                            type='username'
                            placeholder='Username'
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required />
                    </div>
                    <div className='profile-text'>What is your...</div>
                    <div className='profile-text'>Favorite game genres?</div>
                </div>
                <div className='submit-container'>
                    <div
                        className='submit'
                        onClick={(e) => {
                            handleSubmit(e);
                        }}>
                        Create profile
                    </div>
                </div>
            </div>

        </div>
    )

};

// Render the dashboard if the response is successful



// if (userData.id === 0 || userData.username === '' || userData.profile_picture === '') {
//     return (
//         <div style={{ textAlign: 'center' }}>
//             <div className='profile-container'>
//                 {/* 
//                 TODO force user to fill out their profile before they have access to anything else.
//                 Creating a username and answering biographical information questions is mandatory,
//                 setting a profile picture should be possible. If nothing is set, a placeholder image is used instead.
//                 "About me" section is optional.
//                 */}
//                 <div className='inputs'>
//                     <div className='profile-text'>Choose your username</div>
//                     <div className='input'>
//                         <input
//                             type='username'
//                             placeholder='Username'
//                             value={username}
//                             onChange={(e) => setUsername(e.target.value)}
//                             required />
//                     </div>
//                     <div className='profile-text'>What is your...</div>
//                     <div className='profile-text'>Favorite game genres?</div>
//                 </div>
//                 <div className='submit-container'>
//                     <div
//                         className='submit'
//                         onClick={(e) => {
//                             handleSubmit(e);
//                         }}>
//                         Create profile
//                     </div>
//                 </div>
//             </div>

//         </div>
//     )
// }

export default Profile;