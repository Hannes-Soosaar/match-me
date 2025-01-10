import React, { createContext, useState, useEffect, useContext } from 'react';

// Create AuthContext
const AuthContext = createContext();

// Define AuthProvider to wrap your app
export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null); // User info
  const [session, setSession] = useState(null); // Session UUID
  const [loading, setLoading] = useState(true); // Loading state

  // Login function
  const login = async (credentials) => {
    const response = await fetch('/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credentials),
    });
    const data = await response.json();
    if (response.ok) {
      const { authToken, sessionUUID, user } = data;
      localStorage.setItem('authToken', authToken);
      localStorage.setItem('sessionUUID', sessionUUID);
      setUser(user);
      setSession(sessionUUID);
    }
  };

  // Logout function
  const logout = () => {
    localStorage.removeItem('authToken');
    localStorage.removeItem('sessionUUID');
    setUser(null);
    setSession(null);
  };

  // Auto-login on refresh
  useEffect(() => {
    const authToken = localStorage.getItem('authToken');
    const sessionUUID = localStorage.getItem('sessionUUID');
    if (authToken && sessionUUID) {
      fetch('/session', {
        headers: { Authorization: `Bearer ${authToken}` },
      })
        .then((response) => {
          if (response.ok) return response.json();
          throw new Error('Session invalid');
        })
        .then((data) => {
          setUser(data.user);
          setSession(sessionUUID);
        })
        .catch(() => logout());
    }
    setLoading(false);
  }, []);

  return (
    <AuthContext.Provider value={{ user, session, login, logout }}>
      {!loading && children}
    </AuthContext.Provider>
  );
};

// Export hook for easy consumption
export const useAuth = () => useContext(AuthContext);