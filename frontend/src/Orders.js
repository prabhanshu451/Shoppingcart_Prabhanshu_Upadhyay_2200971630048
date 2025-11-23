import React, { useEffect, useState } from "react";
import api from "./api";

export default function Orders() {
  const [orders, setOrders] = useState([]);

  // Load all orders from backend
  const loadOrders = async () => {
    const res = await api("/orders", "GET");
    if (res.ok) {
      setOrders(res.data);
    } else {
      alert("Failed to load orders");
    }
  };

  useEffect(() => {
    loadOrders();
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h2>Order History</h2>

      {orders.length === 0 ? (
        <p>No orders found.</p>
      ) : (
        <ul>
          {orders.map((o) => (
            <li key={o.id}>Order #{o.id}</li>
          ))}
        </ul>
      )}
    </div>
  );
}
