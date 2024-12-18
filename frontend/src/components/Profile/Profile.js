import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import './Profile.css';

const Profile = () => {

    const [username, setUsername] = useState('');
    const [aboutMe, setAboutMe] = useState('');
    const [usernameText, setUsernameText] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        alert(`This feature is WIP\nUsername: ${username}\nAbout Me: ${aboutMe}`);
    }



    useEffect(() => {
        const profileNotExist = localStorage.getItem('profileExists') === 'doesNotExist';
        if (!profileNotExist) {
            setUsernameText("Change your username");
        } else {
            setUsernameText("Choose your username")
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
                <div className='inputs'>
                    <div className='profile-text'>{usernameText}</div>
                    <div className='input'>
                        <input
                            type='username'
                            placeholder='Username'
                            maxLength="20"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required />
                    </div>
                    <div className='profile-text'>Write something about yourself</div>
                    <div className='input-textarea'>
                        <textarea
                            placeholder='About me'
                            maxLength="500"
                            value={aboutMe}
                            onChange={(e) => setAboutMe(e.target.value)}
                            required />
                    </div>
                    <div className='profile-text'>Upload a profile picture*</div>
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

export default Profile;