
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './BuddiesSection.css'
import '../MatchCard/MatchCard.jsx'
import MatchCard from '../MatchCard/MatchCard.jsx';

const authToken = localStorage.getItem('token');





const Matches = () => {

    const [data, setData] = useState([])
    const [matches, setMatches] = useState([])
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('/buddies', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });

                if (response.data === null) {
                    response.data = [];
                }
                console.log(response.data);
                setMatches(response.data);
            }
            catch (error) {
                console.error('Error fetching data: ', error)
            }
        }
        fetchData();
    }, [])

    return (
        <>
            <h1>Buddies</h1>
            <div className="body-div">
                <div className="body-sides"></div>
                <div className='body-content'>
                    {matches.length > 0 ? (
                    matches.map((item, index) =>
                    (<p key={index}>
                        <MatchCard userProfile={item} key={index}></MatchCard>
                    </p>
                    ))):(<p>No matches found. Try updating your preferences or check back later!</p>)}
                </div>
                <div className="body-sides"></div>
            </div>

        </>
    );
};

export default Matches;