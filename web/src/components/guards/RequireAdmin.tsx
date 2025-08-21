// src/components/guards/RequireAdmin.tsx
import { Navigate, useLocation } from "react-router-dom";
import jwt_decode from "jwt-decode";

interface RequireAdminProps {
    children: JSX.Element;
}

/**
 * Admin-only route guard
 * - Checks for token in localStorage
 * - Validates token expiration
 * - Verifies user role is "admin"
 */
const RequireAdmin = ({ children }: RequireAdminProps) => {
    const token = localStorage.getItem("token");
    const location = useLocation(); // capture current route

    if (!token) {
        console.warn("No token found, redirecting to login");
        return <Navigate to="/login" state={{ from: location }} replace />;
    }

    try {
        const decoded: any = jwt_decode(token);

        // check expiry
        if (decoded.exp && decoded.exp * 1000 < Date.now()) {
            console.warn("Token expired, redirecting to login");
            localStorage.removeItem("token");
            return <Navigate to="/login" state={{ from: location }} replace />;
        }

        // check role
        if (decoded.role !== "admin") {
            console.warn("Non-admin tried to access admin route, redirecting");
            return <Navigate to="/dashboard" replace />;
        }

        // ✅ authorized
        return children;
    } catch (err) {
        console.error("Token decode failed", err);
        return <Navigate to="/login" state={{ from: location }} replace />;
    }
};

export default RequireAdmin;

