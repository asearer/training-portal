import { ReactNode } from "react";
import { useNavigate } from "react-router-dom";

interface LayoutProps {
    children: ReactNode;
}

/**
 * Shared layout with navbar and logout button
 */
export default function Layout({ children }: LayoutProps) {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    };

    return (
        <div>
            {/* Navbar */}
            <nav className="bg-blue-600 text-white p-4 flex justify-between">
                <span className="font-bold">Training Portal</span>
                <div>
                    <button
                        onClick={handleLogout}
                        className="bg-red-500 px-3 py-1 rounded hover:bg-red-600 transition"
                    >
                        Logout
                    </button>
                </div>
            </nav>

            {/* Main content */}
            <main className="p-4">{children}</main>
        </div>
    );
}

