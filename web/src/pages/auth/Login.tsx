import { useNavigate } from "react-router-dom";

/**
 * Dummy Login Page
 * - Lets you "log in" as either a normal user or an admin
 * - Stores a fake JWT token in localStorage
 * - Redirects to dashboard or admin panel after login
 */
export default function Login() {
    const navigate = useNavigate();

    // Example JWTs (with exp claim set far in the future)
    const userToken = JSON.stringify({
        role: "user",
        exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24, // valid for 24h
    });

    const adminToken = JSON.stringify({
        role: "admin",
        exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24, // valid for 24h
    });

    // Encode JSON as base64-like string to simulate a JWT
    const encodeToken = (payload: string) =>
        btoa("header") + "." + btoa(payload) + "." + btoa("signature");

    const handleLogin = (role: "user" | "admin") => {
        const token = role === "admin" ? encodeToken(adminToken) : encodeToken(userToken);
        localStorage.setItem("token", token);

        // Redirect based on role
        navigate(role === "admin" ? "/admin" : "/dashboard");
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
            <div className="bg-white shadow-md rounded-lg p-8 w-80">
                <h1 className="text-xl font-bold mb-6 text-center">Dummy Login</h1>

                <button
                    onClick={() => handleLogin("user")}
                    className="w-full bg-blue-500 text-white px-4 py-2 rounded-lg mb-4 hover:bg-blue-600"
                >
                    Login as User
                </button>

                <button
                    onClick={() => handleLogin("admin")}
                    className="w-full bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600"
                >
                    Login as Admin
                </button>
            </div>
        </div>
    );
}
