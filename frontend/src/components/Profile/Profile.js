import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import './Profile.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';

const Profile = () => {

    const [username, setUsername] = useState('');
    const [aboutMe, setAboutMe] = useState('');
    const [usernameText, setUsernameText] = useState('');
    const [profilePic, setProfilePic] = useState(null);
    const [previewPic, setPreviewPic] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();

        //TODO send form data to backend, if profilePic == null, save the default profile pic path to database for profile.
        //if profilePic exists, upload it to /frontend/src/components/Assets/ProfilePictures and save the path to database for profile.
        //save username and about me in the database for profile

        alert(`This feature is WIP\nUsername: ${username}\nAbout Me: ${aboutMe}\nProfile pic: ${profilePic}`);
    }

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            setProfilePic(file)
        }

        const reader = new FileReader();
        reader.onloadend = () => {
            setPreviewPic(reader.result);
        };
        reader.readAsDataURL(file);
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
                    <div className='input-profile-pic'>
                        <label htmlFor="file-input" className="profile-pic-label">
                            {previewPic ?
                                (<img src={previewPic} alt="Preview" />

                                )
                                :
                                (
                                    <img src={defaultProfilePic} alt="Default Profile" />
                                )}
                        </label>
                        <input
                            id="file-input"
                            type="file"
                            accept="image/*"
                            onChange={handleImageChange}
                            style={{ display: 'none' }}
                        />
                    </div>
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