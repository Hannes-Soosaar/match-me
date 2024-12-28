import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Profile.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import { CitySelect, CountrySelect, StateSelect } from 'react-country-state-city';
import 'react-country-state-city/dist/react-country-state-city.css';

const Profile = () => {
    const [countryId, setCountryId] = useState(0);
    const [stateId, setStateId] = useState(0);
    const [username, setUsername] = useState('');
    const [about, setAboutMe] = useState('');
    const [usernameText, setUsernameText] = useState('');
    const [profilePic, setProfilePic] = useState(null);
    const [previewPic, setPreviewPic] = useState(null);
    const [formData, setFormData] = useState({
        country: '',
        state: '',
        city: '',
    });

    // Retrieve authentication token
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('/me', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                const data = response.data;

                // Populate fields with data, or leave them empty if not provided
                setUsername(data.username || '');
                setAboutMe(data.about_me || '');
                setCountryId(data.user_nation || 0);
                setStateId(data.user_state || 0);
                setFormData({
                    country: data.user_nation || '',
                    state: data.user_region || '',
                    city: data.user_city || '',
                });

                // Handle profile picture (default or from backend)
                if (data.profilePic) {
                    setPreviewPic(data.profilePic); // Assuming it's a URL
                } else {
                    setPreviewPic(defaultProfilePic);
                }
            } catch (error) {
                console.error('Error fetching profile data:', error);
            }
        };

        fetchData();
    }, [authToken]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validate the location data
        if (!countryId || !stateId || !formData.city) {
            alert('Please select a valid country, state, and city.');
            return;
        }

        const payload = {
            ...formData,
            countryId,
            stateId,
        };

        try {
            // Send the username to the /username endpoint
            await axios.post(
                '/username',
                { username },
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            console.log('Username updated successfully!');

            // Send the "About Me" text to the /about endpoint
            await axios.post(
                '/about',
                { about },
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            console.log('About Me updated successfully!');

            // Send the selected city, state, and region details to /city endpoint
            await axios.post(
                '/city',
                payload,
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            console.log('City information updated successfully!');

            // Handle profile picture upload
            if (profilePic) {
                const picData = new FormData();
                picData.append('profilePic', profilePic);

                await axios.post('/me', picData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                console.log('Profile picture updated successfully!');
            }

            alert('Profile updated successfully!');
        } catch (error) {
            console.error('Error updating profile:', error);
            alert('There was an error updating your profile. Please try again.');
        }
    };

    const onCountryChange = (country) => {
        console.log('Selected country:', country);
        if (country?.id && country?.name) {
            setCountryId(country.id); // Update country ID
            setFormData((prevData) => ({
                ...prevData,
                country: country.name, // Store country name
                state: '', // Reset state
                city: '', // Reset city
            }));
        } else {
            console.error('Invalid country data received:', country);
        }
    };

    const onStateChange = (state) => {
        console.log('Selected state:', state);
        if (state?.id && state?.name) {
            setStateId(state.id); // Update state ID
            setFormData((prevData) => ({
                ...prevData,
                state: state.name, // Store state name
                city: '', // Reset city
            }));
        } else {
            console.error('Invalid state data received:', state);
        }
    };

    const handleCitySelect = (city) => {
        console.log('Selected city:', city);
        if (city?.name) {
            setFormData((prevData) => ({
                ...prevData,
                city: city.name, // Store city name
            }));
        } else {
            console.error('Invalid city data received:', city);
        }
    };

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            setProfilePic(file);
        }

        const reader = new FileReader();
        reader.onloadend = () => {
            setPreviewPic(reader.result);
        };
        reader.readAsDataURL(file);
    };

    useEffect(() => {
        const profileNotExist = localStorage.getItem('profileExists') === 'doesNotExist';
        setUsernameText(profileNotExist ? "Choose your username" : "Change your username");
    }, []);

    return (
        <div style={{ textAlign: 'center' }}>
            <div className='profile-container'>
                <div className='inputs'>
                    <div className='profile-text'>{usernameText}</div>
                    <div className='input'>
                        <input
                            type='text'
                            placeholder='Username'
                            maxLength="20"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required
                        />
                    </div>
                    <div className='profile-text'>Write something about yourself</div>
                    <div className='input-textarea'>
                        <textarea
                            placeholder='About me'
                            maxLength="500"
                            value={about}
                            onChange={(e) => setAboutMe(e.target.value)}
                            required
                        />
                    </div>
                    <div className='profile-text'>Upload a profile picture*</div>
                    <div className='input-profile-pic'>
                        <label htmlFor="file-input" className="profile-pic-label">
                            {previewPic ? (
                                <img src={previewPic} alt="Preview" />
                            ) : (
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
                <div className="inputGroup">
                    <h6>Country</h6>
                    <CountrySelect
                        onChange={onCountryChange}
                        placeHolder="Select Country"
                        defaultValue={{ id: countryId, name: formData.country }}
                    />
                    <h6>State</h6>
                    <StateSelect
                        countryid={countryId}
                        onChange={onStateChange}
                        placeHolder="Select State"
                        defaultValue={{ id: stateId, name: formData.state }}
                    />
                    <h6>City</h6>
                    <CitySelect
                        countryid={countryId}
                        stateid={stateId}
                        onChange={handleCitySelect}
                        placeHolder="Select City"
                        defaultValue={{ name: formData.city }}
                    />
                </div>
                <div className='submit-container'>
                    <button
                        className='submit'
                        onClick={(e) => handleSubmit(e)}
                    >
                        Create profile
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Profile;
