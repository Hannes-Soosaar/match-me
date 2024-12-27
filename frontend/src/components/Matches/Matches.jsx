import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Matches.css'
import '../MatchCard/MatchCard.jsx'
import MatchCard from '../MatchCard/MatchCard.jsx';



const Matches = () => {

    const [data, setData] = useState([])

// Substitue data with API call in the future.
    const users = [
        {
            imagePath: '/images/user1.jpg',
            firstName: 'John',
            lastName: 'Doe',
            location: 'New York, USA',
            matchScore: '85%',
            description: 'A software developer with a passion for open-source projects.',
        },
        {
            imagePath: '/images/user2.jpg',
            firstName: 'Jane',
            lastName: 'Smith',
            location: 'London, UK',
            matchScore: '92%',
            description: 'A creative designer who loves bringing ideas to life.',
        },
        {
            imagePath: '/images/user3.jpg',
            firstName: 'Carlos',
            lastName: 'Gonzalez',
            location: 'Madrid, Spain',
            matchScore: '78%',
            description: 'An experienced project manager with a knack for problem-solving.',
        },
    ];


    useEffect(() => {
        const fetchData = async () => {
            try {
                //Change to the correct API
                const response = await axios.get('/test')
                setData(response.data);
            }
            catch (error) {
                console.error('Error fetching data: ', error)
            }
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
            {users.map((user, index) => (
            <MatchCard userProfile={user} key={index}></MatchCard>
            ))}
            
        </>

    );
};

export default Matches;