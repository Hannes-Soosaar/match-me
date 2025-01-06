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
    const [matchID, setMatchID] = useState(null)
    const [senderID, setSenderID] = useState("")
    const [receiverID, setReceiverID] = useState("")
    const basePictureURL = "http://localhost:4000/uploads/";
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        if (connections.length > 0) {
            setSelectedConnection(connections[0].matched_user_name)
            setMatchID(connections[0].match_id)
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
                console.log('Profile/me:', response.data)
            } catch (error) {
                console.error('Error getting username:', error)
            }
        }
        fetchUsername()
    }, [authToken])

    useEffect(() => {
        const fetchUUID = async () => {
            try {
                const response = await axios.get('/me/uuid', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                })
                setSenderID(response.data)
                console.log('Profile/me/uuid:', response.data)
            } catch (error) {
                console.error('Error getting sender UUID:', error)
            }
        }
        fetchUUID()
    }, [authToken])

    useEffect(() => {
        if (!socket) return;

        socket.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data])
        }

    }, [socket])

    useEffect(() => {
        const fetchReceiverUUID = async () => {
            try {
                const response = await axios.get('/receiver', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                    params: { senderID, matchID }
                })
                setReceiverID(response.data)
                console.log('/receiver:', response.data)
            } catch (error) {
                console.log('Error getting receiver UUID', error)
            }
        }
        if (matchID) {
            fetchReceiverUUID()
        }
    }, [matchID, senderID, authToken])

    const sendMessage = async () => {
        if (socket && newMessage && matchID && senderID && receiverID) {
            var msgToSend = chatUsername + newMessage
            socket.send(msgToSend)

            try {
                const response = await axios.post('/saveMessage', {
                    matchID: parseInt(matchID, 10),
                    senderID: senderID,
                    receiverID: receiverID,
                    message: newMessage,
                })

                console.log("Message saved:", response.data)
            } catch (error) {
                console.error('Error saving message:', error)
            }
            setNewMessage("")
        }
    }

    const handleConnectionClick = (connection) => {
        console.log('Connection clicked:', connection)
        setSelectedConnection(connection.matched_user_name)
        setMatchID(connection.match_id)
        console.log(connection.match_id)
        //TODO: Make API for pulling all messages with match id equal to connection.match_id a.k.a matchID, then setMessages(array of all pulled messages)
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