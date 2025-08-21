// src/components/Layout.tsx
import { ReactNode } from "react";
import { Link, useNavigate } from "react-router-dom";

interface LayoutProps {
    children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    };

    return (
        <div className="flex min-h-screen bg-gray-100">
            {/* Sidebar / Navbar */}
            <aside className="w-64 bg-blue-700 text-white flex flex-col p-4">
                <h1 className="text-2xl font-bold mb-6">Training Portal</h1>
                <nav className="flex flex-col gap-3">
                    <Link className="hover:bg-blue-600 rounded px-3 py-2" to="/dashboard">
                        Dashboard
                    </Link>
                    <Link className="hover:bg-blue-600 rounded px-3 py-2" to="/courses">
                        Courses
                    </Link>
                    <Link className="hover:bg-blue-600 rounded px-3 py-2" to="/profile">
                        Profile
                    </Link>
                    <Link className="hover:bg-blue-600 rounded px-3 py-2" to="/admin">
                        Admin Panel
                    </Link>
                </nav>
                <button
                    onClick={handleLogout}
                    className="mt-auto bg-red-500 hover:bg-red-600 px-3 py-2 rounded"
                >
                    Logout
                </button>
            </aside>

            {/* Main content */}
            <main className="flex-1 p-6">{children}</main>
        </div>
    );
}


