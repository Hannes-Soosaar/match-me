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
    const [typingStatus, setTypingStatus] = useState("")
    const [offset, setOffset] = useState(0)
    const [hasMore, setHasMore] = useState(false)
    const [buttonsDisabled, setButtonsDisabled] = useState(false)
    const basePictureURL = "http://localhost:4000/uploads/";
    const onlineURL = "/images/OnlineIconPNG.png"
    const offlineURL = "/images/OfflineIconPNG.png"
    const authToken = localStorage.getItem('token');

    /*
    type BuddiesResponse struct {
        MatchID                int    `json:"match_id"`
        MatchScore             int    `json:"match_score"`
        Status                 string `json:"status"`
        MatchedUserName        string `json:"matched_user_name"`
        MatchedUserPicture     string `json:"matched_user_picture"`
        MatchedUserDescription string `json:"matched_user_description"`
        MatchedUserLocation    string `json:"matched_user_location"`
        IsOnline			   bool   `json:"is_online"`
        UserInterests		   []string `json:"user_interests"`
        add notifications field
        ChatNotifications bool `json:"has_notification"`
    }
        */

    useEffect(() => {
        if (!selectedConnection) {
            if (connections === null) {
                return
            }
            if (connections.length > 0) {
                handleConnectionClick(connections[0])
                setOffset(0)
            }
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

    let typingTimeouts = [];

    useEffect(() => {
        if (!socket) return;
        socket.onmessage = (event) => {
            try {
                const data = JSON.parse(event.data);
                console.log('Data type:', data.type);

                if (data.type === "login") {
                    console.log('FE logged in:', data)
                    Object.entries(connections).forEach(([_, connection]) => {
                        if (connection.matched_user_name === data.username) {
                            connection.is_online = true
                        }
                    })
                } else if (data.type === "logout") {
                    console.log('FE logged out', data)
                    Object.entries(connections).forEach(([_, connection]) => {
                        if (connection.matched_user_name === data.username) {
                            connection.is_online = false
                        }
                    })
                }


                if (receiverID === data.senderID && senderID === data.receiverID) {
                    if (data.type === "typing") {
                        setTypingStatus(`${selectedConnection} is typing.`)

                        typingTimeouts.forEach(timeout => clearTimeout(timeout));
                        typingTimeouts = [];

                        const timeout1 = setTimeout(() => {
                            setTypingStatus(`${selectedConnection} is typing..`)
                        }, 1000)

                        const timeout2 = setTimeout(() => {
                            setTypingStatus(`${selectedConnection} is typing...`);
                        }, 2000)

                        const timeout3 = setTimeout(() => {
                            setTypingStatus("")
                        }, 3000)

                        typingTimeouts.push(timeout1, timeout2, timeout3);
                    } else if (data.type === "top-connection") {
                        //rearrange connections so that data.Username is set as top connection in the list
                    }
                } else {
                    setOffset((prevOffset) => prevOffset + 1)
                }

                console.log(
                    "receiverID:", receiverID, "\n", "senderID:", senderID, "\n", "data.receiverID:", data.receiverID, "\n", "data.senderID:", data.senderID, "\n",
                )
                if ((receiverID === data.senderID && senderID === data.receiverID) || senderID === data.senderID) {
                    setMessages((prevMessages) => [...(prevMessages || []), data.message])
                    setConnections(prevConnections => {
                        // Create a copy of the connections array
                        const updatedConnections = [...prevConnections];

                        // Iterate through the connections and move the matched connection to the start
                        updatedConnections.forEach((connection, index, array) => {
                            if (connection.matched_user_name === data.username) {
                                // Remove the matched connection from its current position
                                const [matchedConnection] = array.splice(index, 1);

                                // Insert it at the start of the array
                                array.unshift(matchedConnection);
                            }
                        });

                        // Return the updated array
                        return updatedConnections;
                    });

                }



            } catch (error) {
                console.error("Error parsing message data:", error)
            }
        }
        return () => {
            typingTimeouts.forEach(timeout => clearTimeout(timeout));
        };
    }, [socket, senderID, receiverID, selectedConnection])

    const handleTyping = () => {
        if (socket) {
            socket.send(JSON.stringify({ type: "typing", senderID, receiverID }))
        }
    }

    /*const handleStopTyping = () => {
        if (socket) {
            socket.send(JSON.stringify({ type: "stopTyping", senderID, receiverID }));
        }
    };*/

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

    const getCurrentDateTime = () => {
        const now = new Date()

        const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]

        const day = now.getDate()
        const month = monthNames[now.getMonth()]
        const year = String(now.getFullYear()).slice(-2)
        let hours = now.getHours()
        const minutes = String(now.getMinutes()).padStart(2, '0')
        const ampm = hours >= 12 ? 'PM' : 'AM'

        hours = hours % 12 || 12

        return `(${day}-${month}-${year} ${hours}:${minutes} ${ampm}) `
    }

    const sendMessage = async () => {
        //send notification to receiver
        if (socket && newMessage && matchID && senderID && receiverID) {
            const currentDateTime = getCurrentDateTime()
            var message = currentDateTime + chatUsername + newMessage

            const msgToSend = {
                senderID: senderID,
                receiverID: receiverID,
                message: message,
                username: username
            }

            const jsonMessage = JSON.stringify(msgToSend)
            setOffset((prevOffset) => prevOffset + 1)
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

    const handleKeyDown = (event) => {
        if (event.key === "Enter") {
            event.preventDefault()
            sendMessage()
        }
    }

    const handleConnectionClick = (connection) => {
        setTypingStatus("")
        console.log('Connection clicked:', connection)
        if (selectedConnection === connection.matched_user_name) {
            console.log(selectedConnection, connection.matched_user_name)
            return
        }
        setHasMore(false)
        setOffset(0)
        setSelectedConnection(connection.matched_user_name)
        setMatchID(connection.match_id)

        //TODO: Make API for pulling all messages with match id equal to connection.match_id a.k.a matchID, then setMessages(array of all pulled messages)
        const fetchChatHistory = async () => {
            try {
                const response = await axios.get('/chatHistory', {
                    params: {
                        matchID: parseInt(connection.match_id, 10),
                    },
                })
                if (response.data == null) {
                    setMessages([])
                    setOffset(0)
                } else if (response.data.length > 14) {
                    setHasMore(true)
                    setOffset(15)
                    setMessages(response.data.slice(0, 15))
                } else {
                    setHasMore(false)
                    setMessages(response.data)
                }
                console.log("History received")

                console.log("response", response.data)
            } catch (error) {
                console.log('Error getting chat history FE', error)
            }
        }
        fetchChatHistory()
    }

    const fetchMoreHistory = async () => {
        try {
            const response = await axios.get('/chatHistory', {
                params: {
                    matchID: parseInt(matchID, 10),
                    offset: offset,
                },
            })
            if (response.data.length > 14) {
                setHasMore(true)
                setOffset(offset + 15)
            } else {
                setHasMore(false)
            }
            console.log("History received")
            setMessages((prevMessages) => [
                ...response.data.slice(0, 15),
                ...(prevMessages || [])
            ]);
        } catch (error) {
            console.log('Error getting chat history FE', error)
        }
    }

    const handleClick = (connection) => {
        if (!buttonsDisabled) {
            setButtonsDisabled(true)
            handleConnectionClick(connection)

            setTimeout(() => {
                setButtonsDisabled(false)
            }, 1000)
        }
    }

    return (
        <>
            <div className="chat-container">
                <div className="chat-sidebar">
                    {connections && connections.length > 0 ? (
                        connections.map((connection, index) => (
                            <div
                                key={index}
                                className={`connection-item ${selectedConnection === connection.matched_user_name ? 'selected' : ''}`}
                                onClick={() => !buttonsDisabled && handleClick(connection)}
                            >
                                <img src={basePictureURL + connection.matched_user_picture} alt={connection.matched_user_name} />
                                <h4>{connection.matched_user_name}</h4>
                                {connection.is_online ?
                                    <img src={onlineURL} alt="User online" className="status-icon"></img>
                                    :
                                    <img src={offlineURL} alt="User offline" className="status-icon"></img>
                                }
                            </div>
                        ))) : (
                        <p>No matches found. Try updating your preferences or check back later!</p>
                    )}
                </div>
                <div className="chat-right-container">
                    <div className="chat-messages">
                        {/* if there are more than 25 messages in history, display clickable show more button*/}
                        {hasMore ?
                            <button onClick={fetchMoreHistory}>Show more</button>
                            : null}
                        {messages != null ?
                            messages.map((msg, index) => (
                                <p key={index}>{msg}</p>
                            )) : null
                        }
                    </div>
                    <div className="chat-input-container">
                        <div className="typing-status-container">
                            <p1>{typingStatus}</p1>
                        </div>
                        <div className="input-container">
                            <input
                                type="text"
                                value={newMessage}
                                maxLength={300}
                                onChange={(e) => setNewMessage(e.target.value)}
                                placeholder="Type your message..."
                                onKeyDown={handleKeyDown}
                                onKeyPress={handleTyping}
                            />
                            <button onClick={sendMessage}>Send</button>
                        </div>
                        <div className="typing-status-container"></div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Chat;