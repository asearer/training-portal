// src/App.tsx
import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";

// --- Pages (feature modules) ---
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import PasswordReset from "./pages/auth/PasswordReset";
import Dashboard from "./pages/dashboard/Dashboard";
import CourseList from "./pages/courses/CourseList";
import CourseDetail from "./pages/courses/CourseDetail";
import Profile from "./pages/profile/Profile";
import AdminPanel from "./pages/admin/AdminPanel";

// --- Guards (authorization checks) ---
import RequireAuth from "./components/guards/RequireAuth";
import RequireAdmin from "./components/guards/RequireAdmin";

// --- Layout (shared wrapper: navbar, logout) ---
import Layout from "./components/Layout";

// --- JWT decoder to extract user role and expiration ---
import jwt_decode from "jwt-decode";

/**
 * Utility function to get user role from JWT token stored in localStorage
 * - Returns role as string ("admin" or "user")
 * - Returns null if token is missing, invalid, or expired
 */
function getUserRole(): string | null {
    const token = localStorage.getItem("token");
    if (!token) return null;

    try {
        const decoded: any = jwt_decode(token);

        // JWT exp is in seconds; convert to milliseconds for comparison
        if (decoded.exp * 1000 < Date.now()) {
            localStorage.removeItem("token"); // remove expired token
            return null;
        }

        return decoded.role || null;
    } catch {
        // token is invalid or cannot decode
        return null;
    }
}

function App() {
    // Quick checks for authentication and role
    const isAuthenticated = !!localStorage.getItem("token");
    const userRole = getUserRole();

    return (
        <Router>
            {/* Outer container with a light background */}
            <div className="min-h-screen bg-gray-50">
                <Routes>
                    {/** --- Public routes --- */}
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/reset-password" element={<PasswordReset />} />

                    {/** --- Protected routes (user must be logged in) --- */}
                    <Route
                        path="/dashboard"
                        element={
                            <RequireAuth>
                                <Layout>
                                    <Dashboard />
                                </Layout>
                            </RequireAuth>
                        }
                    />
                    <Route
                        path="/courses"
                        element={
                            <RequireAuth>
                                <Layout>
                                    <CourseList />
                                </Layout>
                            </RequireAuth>
                        }
                    />
                    <Route
                        path="/course/:id"
                        element={
                            <RequireAuth>
                                <Layout>
                                    <CourseDetail />
                                </Layout>
                            </RequireAuth>
                        }
                    />
                    <Route
                        path="/profile"
                        element={
                            <RequireAuth>
                                <Layout>
                                    <Profile />
                                </Layout>
                            </RequireAuth>
                        }
                    />

                    {/** --- Admin-only route --- */}
                    <Route
                        path="/admin"
                        element={
                            <RequireAdmin>
                                <Layout>
                                    <AdminPanel />
                                </Layout>
                            </RequireAdmin>
                        }
                    />

                    {/** --- Catch-all route --- */}
                    <Route
                        path="*"
                        element={
                            <Navigate
                                to={
                                    isAuthenticated
                                        ? userRole === "admin"
                                            ? "/admin"
                                            : "/dashboard"
                                        : "/login"
                                }
                            />
                        }
                    />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
