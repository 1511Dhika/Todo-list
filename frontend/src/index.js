import React from "react";
import { createRoot } from "react-dom/client"; // Vite
import App from "./App";
import { TodoProvider } from "./context/TodoContext";
import { CategoryProvider } from "./context/CategoryContext";

createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <CategoryProvider>
      <TodoProvider>
        <App />
      </TodoProvider>
    </CategoryProvider>
  </React.StrictMode>
);
