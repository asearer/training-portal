import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/auth/Login";
import Dashboard from "./pages/dashboard/Dashboard";
import CourseList from "./pages/courses/CourseList";
import CourseDetail from "./pages/courses/CourseDetail";

// This file is a placeholder for future extensibility.
// Currently, routing is handled in App.tsx, but you can move it here for larger apps.

const AppRouter: React.FC = () => {
  const isAuthenticated = !!localStorage.getItem("token");

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="/dashboard"
          element={isAuthenticated ? <Dashboard /> : <Navigate to="/login" />}
        />
        <Route
          path="/courses"
          element={isAuthenticated ? <CourseList /> : <Navigate to="/login" />}
        />
        <Route
          path="/course/:id"
          element={isAuthenticated ? <CourseDetail /> : <Navigate to="/login" />}
        />
        <Route path="*" element={<Navigate to={isAuthenticated ? "/dashboard" : "/login"} />} />
      </Routes>
    </Router>
  );
};

export default AppRouter;
