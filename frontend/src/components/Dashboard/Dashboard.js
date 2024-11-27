import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Landing = () => {
    const [userData, setUserData] = useState(null); // Store user data
    const [loading, setLoading] = useState(true); // Track loading state
    const [error, setError] = useState(null); // Track errors

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    setError('No token found'); // Handle missing token
                    setLoading(false);
                    return;
                }

                // Wait for the API response with async/await
                const response = await axios.get('/me', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                setUserData(response.data); // Store the user data
            } catch (err) {
                setError(err.response ? err.response.data : 'An error occurred');
            } finally {
                setLoading(false); // Stop loading
            }
        };

        fetchUserData();
    }, []);

    if (loading) {
        // Show a loading screen while waiting for the response
        return <div>Loading...</div>;
    }

    if (error) {
        // Show an error message if the request fails
        return <div>Error: {error}</div>;
    }

    // Render the dashboard if the response is successful
    return (
        <div style={{ textAlign: 'center' }}>
            <h1>Welcome, {userData.name || 'User'}</h1>
            <p>This is your dashboard.</p>
        </div>
    );
};

export default Landing;