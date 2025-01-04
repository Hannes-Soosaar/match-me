import React, { useEffect, useState, useContext } from 'react';
import './Chat.css';
import axios from 'axios';
import { WebSocketContext } from '../WebSocketContext';

const Chat = () => {
    const socket = useContext(WebSocketContext)
    const [messages, setMessages] = useState([])
    const [newMessage, setNewMessage] = useState("")
    const [connections, setConnections] = useState([])
    const basePictureURL = "http://localhost:4000/uploads/";
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        const fetchConnections = async () => {
            try {
                const response = await axios.get('/matches', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                })

                setConnections(response.data)
            } catch (error) {
                console.error('Error getting connections profiles:', error)
            }
        }
        fetchConnections()
    }, [authToken])

    useEffect(() => {
        if (!socket) return;

        socket.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data])
        }

    }, [socket])

    const sendMessage = () => {
        if (socket && newMessage) {
            socket.send(newMessage)
            setNewMessage("")
        }
    }

    return (
        <>
            <div className="chat-container">
                <div className="chat-sidebar">
                    {connections.map((connection, index) => (
                        <div key={index} className="connection-item">
                            <img src={basePictureURL + connection.matched_user_picture} alt={connection.matched_user_name} />
                            <h4>{connection.matched_user_name}</h4>
                        </div>
                    ))}
                </div>
                <div className="chat-right-container">
                    <div className="chat-messages">
                        {messages.map((msg, index) => (
                            <p key={index}>{msg}</p>
                        ))}
                    </div>
                    <div className="chat-input-container">
                        <input
                            type="text"
                            value={newMessage}
                            maxLength={300}
                            onChange={(e) => setNewMessage(e.target.value)}
                            placeholder="Type your message..."
                        />
                        <button onClick={sendMessage}>Send</button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Chat;