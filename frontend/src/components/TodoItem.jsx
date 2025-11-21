// src/components/TodoItem.jsx
import React from "react";
import { useTodo } from "../context/TodoContext";

const TodoItem = ({ todo }) => {
  const { deleteTodo, toggleTodo } = useTodo();

  return (
    <div className="card mb-2 p-3 shadow-sm d-flex justify-content-between align-items-center">
      <div>
        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() => toggleTodo(todo.id)}
          style={{ marginRight: "10px" }}
        />
        <span style={{ textDecoration: todo.completed ? "line-through" : "none" }}>
          {todo.title}
        </span>
      </div>

      <button
        className="btn btn-danger btn-sm"
        onClick={() => deleteTodo(todo.id)}
      >
        Hapus
      </button>
    </div>
  );
};

export default TodoItem;
