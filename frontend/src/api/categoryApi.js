import axiosClient from "./axiosClient";

const categoryApi = {
  getCategories: () => axiosClient.get("/categories"),
  createCategory: (data) => axiosClient.post("/categories", data),
  updateCategory: (id, data) => axiosClient.put(`/categories/${id}`, data),
  deleteCategory: (id) => axiosClient.delete(`/categories/${id}`)
};

export default categoryApi;
