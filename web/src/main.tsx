import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";
import axios from "axios";

// Set Axios base URL from environment variable or default
axios.defaults.baseURL =
  import.meta.env.VITE_API_URL || "http://localhost:3000";

const root = document.getElementById("root");

if (root) {
  ReactDOM.createRoot(root).render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  );
}
