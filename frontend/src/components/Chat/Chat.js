import React, { useEffect, useState } from 'react';
import './Chat.css';

const Chat = () => {
    const [socket, setSocket] = useState(null)
    const [messages, setMessages] = useState([])
    const [newMessage, setNewMessage] = useState("")

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
            <p>This page is for testing the chatting module.</p>

            <div className="chat-container">
                <div className="messages">
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
        </>
    )
}

export default Chat;