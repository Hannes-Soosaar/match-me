import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './Header.css'

function Header() {
    const [isOnline, setIsOnline] = useState(null);
    const isAuthenticated = !!localStorage.getItem('token');
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        const fetchOnlineStatus = async () => {
            if (isAuthenticated) {
                try {
                    const response = await axios.get('/online', {
                        headers: {
                            Authorization: `Bearer ${authToken}`,
                        },
                    });
                    setIsOnline(response.data); // Assuming the backend returns true/false
                } catch (error) {
                    console.error('Error fetching online status:', error);
                }
            }
        };
        fetchOnlineStatus();
    }, [isAuthenticated, authToken]);


    const handleLogout = () => {

        const logout = async () => {
            try {
                const response = await axios.post('/logout', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
            }
            catch (error) {
                console.error('logging out: ', error)
                return
            }
        }
        logout();

        localStorage.removeItem('token');
        localStorage.removeItem('profileExists');
        window.location.href = '/login';
    }

    return (
        <>
            <div className='body-side'></div>
            <header className="header">
                    <div className="nav-left">

                        {isAuthenticated ?  (
                            <>
                                <Link to="/dashboard" className="nav-link">
                                    Dashboard
                                </Link>
                                {/* <Link to="/profile" className="nav-link">
                                    Profile
                                </Link> */}
                                <Link to="/matches" className="nav-link">
                                    Matches
                                </Link>
                                <Link to="/Requests" className="nav-link">
                                    Requests
                                </Link>
                                <Link to="/connections" className="nav-link">
                                    Buddies
                                </Link>
                                <Link to="/chat" className="nav-link">
                                    Chat
                                </Link>
                            </>
                        ) : ( 
                            <Link to='/' className="logo">
                                Gamers Pot
                            </Link> 

                        )}
                    </div>
                        <div className='nav-container'></div>
                    <div className="nav-right">
                    {isAuthenticated && (
                        <div className="online-status">
                            <span
                                className={`status-light ${
                                    isOnline === true
                                        ? 'online' // Green if online
                                        : isOnline === false
                                        ? 'offline' // Red if offline
                                        : '' // No color if status is unknown
                                }`}
                            ></span>
                            <span className="indicator-text">{isOnline === true ? 'Online' : isOnline === false ? 'Offline' : 'Loading...'}</span>
                        </div>
                    )}
                        {!isAuthenticated ? (
                            <Link to="/login" className="signup">
                                Sign up/Login
                            </Link>
                        ) : (
                            <Link
                                to="#"
                                className="signup"
                                onClick={(e) => {
                                    e.preventDefault();
                                    handleLogout();
                                }}
                            >
                                Logout
                            </Link>
                        )}
                    </div>
                </header>
            <div className='body-side'></div>
        </>
    );
}

export default Header;