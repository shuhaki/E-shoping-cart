import React, { useEffect, useState } from "react";
import API from "../api";

export default function ItemList() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    API.get("/items").then((res) => setItems(res.data));
  }, []);

  const addToCart = async (itemId) => {
    await API.post("/carts", {
      user_id: localStorage.getItem("user_id"),
      item_id: itemId,
    });
    alert("Item added to cart!");
  };

  return (
    <div>
      <h2>Items</h2>
      {items.map((item) => (
        <div key={item._id}>
          {item.name} - {item.status}
          <button onClick={() => addToCart(item._id)}>Add to Cart</button>
        </div>
      ))}
      <hr />
      <button onClick={() => window.location.href = "/cart"}>🛒 View Cart</button>
      <button onClick={() => window.location.href = "/orders"}>📜 Order History</button>
    </div>
  );
}
