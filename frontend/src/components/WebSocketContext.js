import React, { createContext, useState, useEffect } from 'react'

export const WebSocketContext = createContext()

export const WebSocketProvider = ({ children }) => {
    const [socket, setSocket] = useState(null)
    const authToken = localStorage.getItem('token')

    useEffect(() => {
        if (authToken) {
            const ws = new WebSocket("ws://localhost:4000/ws")

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

    }, [])

    return (
        <WebSocketContext.Provider value={socket}>
            {children}
        </WebSocketContext.Provider>
    )
}