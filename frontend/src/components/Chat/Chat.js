import React, { useEffect, useState, useContext } from 'react';
import './Chat.css';
import axios from 'axios';
import { WebSocketContext } from '../WebSocketContext';

const Chat = () => {
    const socket = useContext(WebSocketContext)
    const [messages, setMessages] = useState([])
    const [newMessage, setNewMessage] = useState("")
    const [connections, setConnections] = useState([])
    const [selectedConnection, setSelectedConnection] = useState(null)
    const [chatUsername, setChatUsername] = useState("")
    const basePictureURL = "http://localhost:4000/uploads/";
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        if (connections.length > 0) {
            setSelectedConnection(connections[0].matched_user_name)
        }
    }, [connections])

    useEffect(() => {
        const fetchConnections = async () => {
            try {
                const response = await axios.get('/matches', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                })

                setConnections(response.data)
                console.log(response.data)
            } catch (error) {
                console.error('Error getting connections profiles:', error)
            }
        }
        fetchConnections()
    }, [authToken])

    useEffect(() => {
        const fetchUsername = async () => {
            try {
                const response = await axios.get('/me', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                })
                setChatUsername(response.data.username + ": ")
            } catch (error) {
                console.error('Error getting username:', error)
            }
        }
        fetchUsername()
    }, [authToken])

    useEffect(() => {
        if (!socket) return;

        socket.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data])
        }

    }, [socket])

    const sendMessage = () => {
        if (socket && newMessage) {
            const msgToSend = chatUsername + newMessage
            socket.send(msgToSend)
            setNewMessage("")
        }
    }

    const handleConnectionClick = (connection) => {
        console.log('Connection clicked:', connection)
        setSelectedConnection(connection.matched_user_name)
        //retrieve chat history and display it.
    }

    return (
        <>
            <div className="chat-container">
                <div className="chat-sidebar">
                    {connections.map((connection, index) => (
                        <div
                            key={index}
                            className={`connection-item ${selectedConnection === connection.matched_user_name ? 'selected' : ''}`}
                            onClick={() => handleConnectionClick(connection)}
                        >
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