import React, { useState } from 'react';
import './LoginSignup.css';
import email_icon from '../Assets/email.png';
import password_icon from '../Assets/password.png';
import axios from 'axios';

const LoginSignup = () => {

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (e, isLoginState) => {
        e.preventDefault();

        const url = isLoginState ? '/login' : '/register';

        const payload = {
            email,
            password
        }

        try {
            const response = await axios.post(url, payload);
            console.log('Response:', response.data);
            alert(response.data.message)
        } catch (error) {
            if (error.response) {
                const errorMessage = error.response.data.error

                if (error.response.status === 400) {
                    alert(errorMessage)
                } else if (error.response.status === 409) {
                    alert(errorMessage);
                } else {
                    alert(errorMessage)
                }
            }
        }
    };

    return (
        <div>
            <div className='login-container'>
                <div className='login-header'>
                    <div className='login-text'>Sign Up or Log In</div>
                </div>
                <div className='inputs'>
                    <div className='input'>
                        <img src={email_icon} alt="" />
                        <input
                            type='email'
                            placeholder='Email'
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required />
                    </div>
                    <div className='input'>
                        <img src={password_icon} alt="" />
                        <input
                            type='password'
                            placeholder='Password'
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required />
                    </div>
                </div>
                <div className='submit-container'>
                    <div
                        className='submit'
                        onClick={(e) => {
                            handleSubmit(e, false);
                        }}>
                        Sign Up
                    </div>
                    <div
                        className='submit'
                        onClick={(e) => {
                            handleSubmit(e, true);
                        }}>
                        Log In
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LoginSignup