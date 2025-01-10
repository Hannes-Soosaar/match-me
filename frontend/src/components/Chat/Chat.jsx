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
    const [username, setUsername] = useState("")
    const [receiverUsername, setReceiverUsername] = useState("")
    const [receiverProfilePicture, setReceiverProfilePicture] = useState("")
    const [matchID, setMatchID] = useState(null)
    const [senderID, setSenderID] = useState("")
    const [receiverID, setReceiverID] = useState("")
    const basePictureURL = "http://localhost:4000/uploads/";
    const authToken = localStorage.getItem('token');

    useEffect(() => {
        if (connections === null) {
            return
        }
        if (connections.length > 0) {
            handleConnectionClick(connections[0])
        }
    }, [connections])

    useEffect(() => {
        const fetchConnections = async () => {
            try {
                const response = await axios.get('/buddies', {
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
                setUsername(response.data.username)
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
            setMessages((prevMessages) => [...(prevMessages || []), event.data])
        }

    }, [socket])

    useEffect(() => {
        if (senderID && matchID) {
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
        }
    }, [matchID, senderID, authToken])

    useEffect(() => {
        const fetchReceiverProfile = async () => {
            try {
                const response = await axios.get(`/users/${receiverID}/profile`, {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    }
                })
                console.log('/users/receiverID:', response.data)
                setReceiverUsername(response.data.username)
                setReceiverProfilePicture(response.data.profile_picture)
            } catch (error) {
                console.log('Error getting receiver profile', error)
            }
        }
        if (receiverID) {
            fetchReceiverProfile()
        }
    }, [authToken, receiverID])

    const sendMessage = async () => {
        if (socket && newMessage && matchID && senderID && receiverID) {
            var message = chatUsername + newMessage

            const msgToSend = {
                senderID: senderID,
                receiverID: receiverID,
                message: message
            }

            const jsonMessage = JSON.stringify(msgToSend)

            socket.send(jsonMessage)

            try {
                const response = await axios.post('/saveMessage', {
                    matchID: parseInt(matchID, 10),
                    senderID: senderID,
                    receiverID: receiverID,
                    message: message,
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

        //TODO: Make API for pulling all messages with match id equal to connection.match_id a.k.a matchID, then setMessages(array of all pulled messages)
        const fetchChatHistory = async () => {
            try {
                const response = await axios.get('/chatHistory', {
                    params: { matchID: parseInt(connection.match_id, 10) },
                })
                console.log("History received")
                setMessages(response.data)
            } catch (error) {
                console.log('Error getting chat history FE', error)
            }
        }
        fetchChatHistory()
    }

    return (
        <>
            <div className="chat-container">
                <div className="chat-sidebar">
                    { connections && connections.length > 0 ? (
                    connections.map((connection, index) => (
                        <div
                            key={index}
                            className={`connection-item ${selectedConnection === connection.matched_user_name ? 'selected' : ''}`}
                            onClick={() => handleConnectionClick(connection)}
                        >
                            <img src={basePictureURL + connection.matched_user_picture} alt={connection.matched_user_name} />
                            <h4>{connection.matched_user_name}</h4>
                        </div>
                    ))):(
                        <p>No matches found. Try updating your preferences or check back later!</p>
                    )}
                </div>
                <div className="chat-right-container">
                    <div className="chat-messages">
                        {messages != null ?
                            messages.map((msg, index) => (
                                <p key={index}>{msg}</p>
                            )) : null
                        }
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