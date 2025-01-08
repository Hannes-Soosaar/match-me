import React, { createContext, useState, useEffect } from 'react'
import axios from 'axios';

export const WebSocketContext = createContext()

export const WebSocketProvider = ({ children }) => {
    const [socket, setSocket] = useState(null)
    const [senderID, setSenderID] = useState("")
    const authToken = localStorage.getItem('token')

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
        if (authToken) {
            console.log("ws made for senderID:", senderID)
            const ws = new WebSocket(`ws://localhost:4000/ws?userID=${senderID}`)

            ws.onopen = () => {
                console.log("WebSocket connected")
            }

            ws.onclose = () => {
                console.log("WebSocket disconnected")
            }

            setSocket(ws)

            return () => {
                ws.close()
            }
        }

    }, [authToken, senderID])

    return (
        <WebSocketContext.Provider value={socket}>
            {children}
        </WebSocketContext.Provider>
    )
}