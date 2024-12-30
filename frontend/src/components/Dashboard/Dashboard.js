import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import './Dashboard.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';

const Dashboard = () => {
    const [userData, setUserData] = useState(null); // Store user data
    const [loading, setLoading] = useState(true); // Track loading state for user data
    const [locationLoading, setLocationLoading] = useState(true); // Track loading state for location
    const [error, setError] = useState(null); // Track errors
    const [profilePic, setProfilePic] = useState(null);
    const [location, setLocation] = useState(null); // Store the user's location
    const navigate = useNavigate();

    const requestLocation = () => {
        return new Promise((resolve, reject) => {
            if (navigator.geolocation) {
                navigator.geolocation.getCurrentPosition(
                    async (position) => {
                        const { latitude, longitude } = position.coords;
                        setLocation({ latitude, longitude });

                        try {
                            // Send location to the /browserlocation endpoint with Bearer token
                            const token = localStorage.getItem('token');
                            if (!token) {
                                reject(new Error('No token found'));
                                return;
                            }

                            const payload = { latitude, longitude };
                            await axios.post('/browserlocation', payload, {
                                headers: {
                                    Authorization: `Bearer ${token}`,
                                },
                            });
                            resolve({ latitude, longitude });
                        } catch (error) {
                            console.error('Error sending location to backend:', error);
                            reject(error);
                        }
                    },
                    (error) => {
                        console.error('Error getting location:', error.message);
                        reject(error);
                    }
                );
            } else {
                console.error('Geolocation is not supported by this browser.');
                reject(new Error('Geolocation not supported'));
            }
        });
    };

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                await requestLocation(); // Request location before user data

                const token = localStorage.getItem('token');
                if (!token) {
                    setError('No token found'); // Handle missing token
                    setLoading(false);
                    return;
                }

                const response = await axios.get('/me', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                const data = response.data; // Store the user data

                setUserData(data);

                // Handle profile picture (default or from backend)
                if (data.profile_picture) {
                    setProfilePic(`/uploads/${data.profile_picture}`); // Assuming it's a URL
                } else {
                    setProfilePic(defaultProfilePic); // Use default profile picture
                }

            } catch (err) {
                setError(err.response ? err.response.data : 'An error occurred');
            } finally {
                setLoading(false); // Stop loading for user data
                setLocationLoading(false); // Stop loading for location
            }
        };

        fetchUserData();
    }, []);

    useEffect(() => {
        if (!loading && userData) {
            if (
                userData.id === 0 ||
                userData.username === "" ||
                userData.profile_picture === ""
            ) {
                localStorage.setItem('profileExists', 'doesNotExist');
                navigate("/profile");
            }
        }
    }, [loading, userData, navigate]);

    if (locationLoading) {
        // Show loading for location
        return <div>Fetching your location...</div>;
    }

    if (loading) {
        // Show loading for user data
        return <div>Loading user data...</div>;
    }

    if (error) {
        // Show an error message if the request fails
        return <div>Error: {error}</div>;
    }

    return (
        <>
            <div className="dashboard-container">
                <div className="dashboard-card">
                    <div className="dashboard-profile-pic">
                        {profilePic ? (
                            <img src={profilePic} alt="Profile" />
                        ) : (
                            <img src={defaultProfilePic} alt="Default Profile" />
                        )}
                    </div>
                    <h2>{userData?.username}</h2>
                    <p>{userData?.email}</p>
                    <p>{userData?.age}</p>
                    <p>{`${userData?.user_nation}, ${userData?.user_region}, ${userData?.user_city}`}</p>
                    <p>{userData?.about_me}</p>

                    {/* Check if location is available before rendering */}
                    {location && (
                        <p>{`Location: ${location.latitude}, ${location.longitude}`}</p>
                    )}
                </div>
            </div>
        </>
    );
};

export default Dashboard;
