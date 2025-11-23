import React, { useState } from "react";
import  api  from "./api";

export default function Login({ onLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const submit = async (e) => {
    e.preventDefault();

    const user = username.trim();
    const pass = password;

    if (!user || !pass) {
      window.alert("Please enter both username and password");
      return;
    }

    setLoading(true);
    try {
      const res = await api("/users/login", "POST", { username: user, password: pass });
      if (!res.ok) {
        // If backend sent a message, show it; otherwise generic message
        const err = (res.data && (res.data.error || res.data.message)) || "Invalid username/password";
        window.alert(err);
        return;
      }

      const token = res.data && res.data.token;
      if (!token) {
        window.alert("Login succeeded but no token was returned by the server.");
        return;
      }

      localStorage.setItem("token", token);
      onLogin(token);
    } catch (err) {
      // network or unexpected error
      console.error("Login error:", err);
      window.alert("Failed to connect to server. Make sure the backend is running on http://localhost:8080");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: 360, margin: "40px auto", textAlign: "center" }}>
      <h2>Login</h2>
      <form onSubmit={submit}>
        <div style={{ marginBottom: 8 }}>
          <input
            placeholder="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            disabled={loading}
          />
        </div>
        <div style={{ marginBottom: 8 }}>
          <input
            placeholder="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            disabled={loading}
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? "Logging in…" : "Login"}
        </button>
      </form>
    </div>
  );
}
