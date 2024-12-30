import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import './Dashboard.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';

const Dashboard = () => {
    const [userData, setUserData] = useState(null); // Store user data
    const [loading, setLoading] = useState(true); // Track loading state for user data
    const [error, setError] = useState(null); // Track errors
    const [profilePic, setProfilePic] = useState(null);
    const navigate = useNavigate();


    useEffect(() => {
        const fetchUserData = async () => {
            try {
                // Request location before user data

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


                </div>
            </div>
        </>
    );
};

export default Dashboard;
