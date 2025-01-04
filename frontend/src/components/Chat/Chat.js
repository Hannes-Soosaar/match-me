import React, { useEffect, useState } from 'react';
import './Chat.css';
import axios from 'axios';

const Chat = () => {
    const [socket, setSocket] = useState(null)
    const [messages, setMessages] = useState([])
    const [newMessage, setNewMessage] = useState("")
    const [connectionsID, setConnectionsID] = useState([])
    const [connections, setConnections] = useState([])

    const authToken = localStorage.getItem('token');

    useEffect(() => {
        const fetchConnectionsID = async () => {
            try {
                const response = await axios.get('/connections', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                setConnectionsID(response.data)
                console.log(response.data)
            } catch (error) {
                console.error('Error fetching connections:', error);
            }
        }

        fetchConnectionsID()
    }, [authToken])

    useEffect(() => {
        const fetchConnections = async () => {

        }
    })

    useEffect(() => {
        const ws = new WebSocket("ws://localhost:4000/ws")

        ws.onopen = () => {
            console.log("WebSocket connected")
        }

        ws.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data])
        }

        ws.onclose = () => {
            console.log("WebSocket disconnected")
            //TODO: Save chat history to the backend when WebSocket closes
        }

        setSocket(ws)

        return () => {
            ws.close()
        }

    }, [])

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
                    {connectionsID.map((connection, index) => (
                        <ul key={index}>{connection}</ul>
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