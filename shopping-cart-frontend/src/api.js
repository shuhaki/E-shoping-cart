/*import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080/api",
});

API.interceptors.request.use((req) => {
  const token = localStorage.getItem("token");
  if (token) req.headers.Authorization = `Bearer ${token}`;
  return req;
});

export default API;*/
const API = {
  post: async (url, data) => {
    if (url === "/users/login") {
      if (data.username === "demo" && data.password === "demo") {
        return {
          data: {
            token: "mocked-token-123",
            user_id: "mocked-user-id"
          }
        };
      } else {
        throw new Error("Invalid username/password");
      }
    }

    if (url === "/carts" || url === "/orders") {
      return { data: { success: true } };
    }
  },

  get: async (url) => {
    if (url === "/items") {
      return {
        data: [
          { _id: "1", name: "Notebook", status: "available" },
          { _id: "2", name: "Pen", status: "available" },
        ],
      };
    }

    if (url === "/carts") {
      return {
        data: [
          {
            user_id: "mocked-user-id",
            items: [{ item_id: "1" }, { item_id: "2" }]
          },
        ],
      };
    }

    if (url === "/orders") {
      return {
        data: [
          {
            _id: "order123",
            user_id: "mocked-user-id",
            items: [{ item_id: "1" }, { item_id: "2" }],
          },
        ],
      };
    }
  },
};

export default API;

