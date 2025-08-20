import { Navigate } from "react-router-dom";
import jwt_decode from "jwt-decode";

/**
 * Guard that allows access only if the user has an admin role in JWT
 */
const RequireAdmin = ({ children }: { children: JSX.Element }) => {
    const token = localStorage.getItem("token");
    if (!token) return <Navigate to="/login" />;

    try {
        const decoded: any = jwt_decode(token);
        return decoded.role === "admin" ? children : <Navigate to="/login" />;
    } catch {
        return <Navigate to="/login" />;
    }
};

export default RequireAdmin;
