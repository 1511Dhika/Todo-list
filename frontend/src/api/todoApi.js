// src/api/todoApi.js
import axiosClient from "./axiosClient";

export const getTodos = async () => {
  const res = await axiosClient.get("/todos");
  return res.data.data;
};

export const createTodo = async (data) => {
  const res = await axiosClient.post("/todos", data);
  return res.data;
};

export const deleteTodoApi = async (id) => {
  await axiosClient.delete(`/todos/${id}`);
};

export const toggleTodoApi = async (id) => {
  const res = await axiosClient.patch(`/todos/${id}/complete`);
  return res.data;
};
