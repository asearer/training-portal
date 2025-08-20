import { Navigate } from "react-router-dom";

/**
 * Guard that allows access only if a JWT token exists in localStorage
 */
const RequireAuth = ({ children }: { children: JSX.Element }) => {
    const token = localStorage.getItem("token");
    return token ? children : <Navigate to="/login" />;
};

export default RequireAuth;

