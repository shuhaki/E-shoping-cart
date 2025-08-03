import React, { useEffect, useState } from "react";
import API from "../api";

export default function Orders() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    API.get("/orders").then((res) => {
      const myOrders = res.data.filter((o) => o.user_id === localStorage.getItem("user_id"));
      setOrders(myOrders);
    });
  }, []);

  return (
    <div>
      <h2>Order History</h2>
      {orders.map((o, idx) => (
        <div key={idx}>Order ID: {o._id}</div>
      ))}
    </div>
  );
}
