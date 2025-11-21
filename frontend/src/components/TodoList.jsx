import React from "react";
import { List } from "antd";
import { useTodo } from "../context/TodoContext";
import TodoItem from "./TodoItem";

const TodoList = () => {
  const { todos } = useTodo();

  return (
    <List
      dataSource={todos}
      renderItem={(todo) => <TodoItem todo={todo} />}
    />
  );
};

export default TodoList;
