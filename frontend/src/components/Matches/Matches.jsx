import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Matches.css'
import '../MatchCard/MatchCard.jsx'
import MatchCard from '../MatchCard/MatchCard.jsx';

const authToken = localStorage.getItem('token');

const Matches = () => {

    const [data, setData] = useState([])
    const [matches, setMatches] = useState([])
    useEffect(() => {
        const fetchData = async () => {

        //GET
            try {
                const response = await axios.get('/matches', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
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
            <div style={{ textAlign: 'center' }}>
                <h1>Matches</h1>
                {matches.map((item, index) =>
                (<p key={index}>
            <MatchCard userProfile={item} key={index}></MatchCard>
                </p>
                ))}
            </div>
            
        </>
    );
};

export default Matches;