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
 
        //POST

        //PUT

        }
        fetchData();
    }, [])

    return (
        <>
    {/* TODO: We need to figure out the payloads that need to be send and compiled at the BE*/}
            <div style={{ textAlign: 'center' }}>
                <h1>This is the page Where I am testing out stuff!</h1>
                {data.map((item, index) =>
                (<p key={index}>{index + 1} : {Object.values(item)[1]}</p>
                ))}
            </div>
            {/* {users.map((user, index) => (
            <MatchCard userProfile={user} key={index}></MatchCard>
            ))} */}
            
        </>

    );
};

export default Matches;