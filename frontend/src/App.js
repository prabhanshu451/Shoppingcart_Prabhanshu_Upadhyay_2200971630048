import React, { useState } from 'react';
import Login from './Login';
import Items from './Items';

function App() {
  const [token, setToken] = useState(localStorage.getItem('token') || null);

  if (!token) {
    return <Login onLogin={(t)=>setToken(t)} />
  }

  return <Items token={token} />;
}

export default App;
