import { createContext, useContext, useState, useEffect } from "react";
import {
  getTodos,
  createTodo,
  deleteTodoApi,
  toggleTodoApi,
} from "../api/todoApi";
import { List } from "antd";

const TodoContext = createContext();

export const TodoProvider = ({ children }) => {
  const [todos, setTodos] = useState([]);

  // load data
  useEffect(() => {
    loadTodos();
  }, []);

  const loadTodos = async () => {
    const list = await getTodos();
    setTodos(Array.isArray(list) ? list : []);
  };


  const addTodo = async (title) => {
    const res = await createTodo({ title });
    const newTodo = res.todo;
    setTodos((prev) => [...prev, newTodo]);
  };

  const deleteTodo = async (id) => {
    await deleteTodoApi(id);
    setTodos(todos.filter((t) => t.id !== id));
  };

  const toggleTodo = async (id) => {
    const res = await toggleTodoApi(id);
    const updated = res.todo;
    setTodos((prev) => prev.map((t) => (t.id === id ? updated : t)));
  };

  return (
    <TodoContext.Provider value={{ todos, addTodo, deleteTodo, toggleTodo }}>
      {children}
    </TodoContext.Provider>
  );
};

export const useTodo = () => useContext(TodoContext);
