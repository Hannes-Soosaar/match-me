import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css'

function Header() {
    const isAuthenticated = !!localStorage.getItem('token');

    const handleLogout = () => {
        localStorage.removeItem('token');
        window.location.href = '/login';
    }

    return (
        <header className="header">
            <Link to='/' className="logo">
                Match me
            </Link>
            {!isAuthenticated ? (
                <Link to='/login' className='signup'>
                    Sign up/Login
                </Link>
            ) : (
                <Link
                    to='#'
                    className='signup'
                    onClick={(e) => {
                        e.preventDefault();
                        handleLogout();
                    }}
                >
                    Logout
                </Link>
            )}
        </header>
    )
}

export default Header;