import React, { useEffect, useState } from "react";
import API from "../api";

export default function Cart() {
  const [cart, setCart] = useState(null);

  useEffect(() => {
    API.get("/carts").then((res) => {
      const userCart = res.data.find((c) => c.user_id === localStorage.getItem("user_id"));
      setCart(userCart);
    });
  }, []);

  const checkout = async () => {
    await API.post("/orders", { user_id: localStorage.getItem("user_id") });
    alert("Order Successful!");
    window.location.href = "/orders";
  };

  return (
    <div>
      <h2>Your Cart</h2>
      {cart?.items?.map((item, idx) => (
        <div key={idx}>Item ID: {item.item_id}</div>
      ))}
      <button onClick={checkout}>Checkout</button>
    </div>
  );
}
