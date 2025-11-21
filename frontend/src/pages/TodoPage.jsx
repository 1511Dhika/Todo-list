import React, { useState } from "react";
import { Button, Typography } from "antd";
import TodoList from "../components/TodoList";
import TodoForm from "../components/TodoForm";

const { Title } = Typography;

const TodoPage = () => {
  const [openForm, setOpenForm] = useState(false);

  return (
    <div>
      <Title level={2}>Todo List</Title>

      <Button
        type="primary"
        onClick={() => setOpenForm(true)}
        style={{ marginBottom: 20 }}
      >
        + Tambah Todo
      </Button>

      {/* Modal Form Tambah Todo */}
      <TodoForm open={openForm} onClose={() => setOpenForm(false)} />

      {/* List Todo */}
      <TodoList />
    </div>
  );
};

export default TodoPage;
