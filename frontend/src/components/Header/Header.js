import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css'

function Header() {
    const isAuthenticated = !!localStorage.getItem('token');

    
    
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
                        <Link to='/' className="logo">
                            Match me
                        </Link>
                        {isAuthenticated && (
                            <>
                                <Link to="/dashboard" className="nav-link">
                                    Dashboard
                                </Link>
                                <Link to="/profile" className="nav-link">
                                    Profile
                                </Link>
                                <Link to="/matches" className="nav-link">
                                    Matches
                                </Link>
                                <Link to="/connections" className="nav-link">
                                    Buddies
                                </Link>
                                <Link to="/chat" className="nav-link">
                                    Chat
                                </Link>
                            </>
                        )}
                    </div>
                        <div className='nav-container'></div>
                    <div className="nav-right">
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