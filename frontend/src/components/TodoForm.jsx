import React, { useState } from "react";
import { Input, Button } from "antd";
import { useTodo } from "../context/TodoContext";

const TodoForm = () => {
  const [title, setTitle] = useState("");
  const { addTodo } = useTodo();

  const handleSubmit = () => {
    if (title.trim() === "") return;
    addTodo(title);
    setTitle("");
  };

  return (
    <div>
      <Input
        placeholder="Tambahkan todo..."
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />

      <Button type="primary" onClick={handleSubmit} style={{ marginTop: 10 }}>
        Tambah
      </Button>
    </div>
  );
};

export default TodoForm;
