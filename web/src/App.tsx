import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import Dashboard from "./pages/dashboard/Dashboard";
import CourseList from "./pages/courses/CourseList";
import CourseDetail from "./pages/courses/CourseDetail";
import AdminPanel from "./pages/admin/AdminPanel";
import Profile from "./pages/profile/Profile";
import PasswordReset from "./pages/auth/PasswordReset";

import jwt_decode from "jwt-decode";

function getUserRole(): string | null {
  const token = localStorage.getItem("token");
  if (!token) return null;
  try {
    const decoded: any = jwt_decode(token);
    return decoded.role || null;
  } catch {
    return null;
  }
}

function App() {
  const isAuthenticated = !!localStorage.getItem("token");
  const userRole = getUserRole();

  // Route guards
  const RequireAuth = ({ children }: { children: JSX.Element }) =>
    isAuthenticated ? children : <Navigate to="/login" />;
  const RequireAdmin = ({ children }: { children: JSX.Element }) =>
    isAuthenticated && userRole === "admin" ? children : <Navigate to="/login" />;

  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route
            path="/dashboard"
            element={
              <RequireAuth>
                <Dashboard />
              </RequireAuth>
            }
          />
          <Route
            path="/courses"
            element={
              <RequireAuth>
                <CourseList />
              </RequireAuth>
            }
          />
          <Route
            path="/course/:id"
            element={
              <RequireAuth>
                <CourseDetail />
              </RequireAuth>
            }
          />
          <Route
            path="/admin"
            element={
              <RequireAdmin>
                <AdminPanel />
              </RequireAdmin>
            }
          />
          <Route
            path="/profile"
            element={
              <RequireAuth>
                <Profile />
              </RequireAuth>
            }
          />
          <Route
            path="/reset-password"
            element={<PasswordReset />}
          />
          <Route path="*" element={<Navigate to={isAuthenticated ? "/dashboard" : "/login"} />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
