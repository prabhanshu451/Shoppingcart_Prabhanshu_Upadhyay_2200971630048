// src/App.js
import React, { useEffect, useState } from "react";
import Login from "./Login";
import Items from "./Items";
import Cart from "./Cart"; // if you have this
import Orders from "./Orders"; // if you have this

export default function App() {
  const [token, setToken] = useState(null);

  useEffect(() => {
    // load token from localStorage on start
    const t = localStorage.getItem("token");
    if (t) setToken(t);
  }, []);

  const handleLogin = (newToken) => {
    localStorage.setItem("token", newToken);
    setToken(newToken);
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setToken(null);
  };

  // If no token -> show Login page only
  if (!token) {
    return <Login onLogin={handleLogin} />;
  }

  // After login: show your main UI (replace with your layout)
  return (
    <div>
      <header style={{ padding: "16px" }}>
        <button onClick={handleLogout}>Logout</button>
      </header>

      <main style={{ padding: "16px" }}>
        {/* show Items etc. */}
        <Items />
        {/* <Cart /> */}
        {/* <Orders /> */}
      </main>
    </div>
  );
}
