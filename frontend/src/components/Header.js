import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css'

function Header() {
    return (
        <header className="header">
            <Link to='/' className="logo">
                Match me
            </Link>
        </header>
    )
}

export default Header;