import React, { createContext, useContext, useState, useEffect } from "react";
import categoryApi from "../api/categoryApi";

const CategoryContext = createContext(null);

export const CategoryProvider = ({ children }) => {
  const [categories, setCategories] = useState([]);

  const fetchCategories = async () => {
    try {
      const res = await categoryApi.getCategories();
      // in your backend handler we returned { "message": "success", "categories": [...] }
      // or maybe returned array. adjust accordingly:
      // if res.data.categories exists use it, otherwise res.data
      setCategories(res.data.categories ?? res.data);
    } catch (err) {
      console.error(err);
    }
  };

  const createCategory = async (data) => {
    await categoryApi.createCategory(data);
    await fetchCategories();
  };

  const updateCategory = async (id, data) => {
    await categoryApi.updateCategory(id, data);
    await fetchCategories();
  };

  const deleteCategory = async (id) => {
    await categoryApi.deleteCategory(id);
    await fetchCategories();
  };

  useEffect(() => {
    fetchCategories();
  }, []);

  return (
    <CategoryContext.Provider
      value={{
        categories,
        fetchCategories,
        createCategory,
        updateCategory,
        deleteCategory,
      }}
    >
      {children}
    </CategoryContext.Provider>
  );
};

export const useCategory = () => {
  const ctx = useContext(CategoryContext);
  if (!ctx) throw new Error("useCategory must be used within CategoryProvider");
  return ctx;
};
