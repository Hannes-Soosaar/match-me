import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Matches.css'

const Matches = () => {

    const [data, setData] = useState([])

    useEffect(() => {
        const fetchData = async () => {
            try {
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
            <div style={{ textAlign: 'center' }}>
                <h1>This is the page with the Matches</h1>
                {data.map((item, index) =>
                (<p key={index}>{index+1} : {Object.values(item)[1]}</p>
                ))}
            </div>
        </>

    );
};

export default Matches;