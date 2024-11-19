import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css'

function Header() {
    return (
        <header className="header">
            <Link to='/' className="logo">
                Match me
            </Link>
            <Link to='/login' className="signup">
                Sign up/login
            </Link>
        </header>
    )
}

export default Header;