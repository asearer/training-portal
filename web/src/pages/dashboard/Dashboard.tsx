// src/pages/dashboard/Dashboard.tsx
import React from "react";

const Dashboard: React.FC = () => {
    // Dummy data for demonstration
    const stats = [
        { title: "Courses Enrolled", value: 5 },
        { title: "Completed Modules", value: 12 },
        { title: "Pending Assignments", value: 3 },
    ];

    return (
        <div className="min-h-screen bg-gray-50 p-6">
            <h1 className="text-3xl font-bold mb-6">Welcome to the Training Portal!</h1>

            <p className="mb-8 text-gray-700">
                Use the sidebar to access your courses, modules, and profile.
            </p>

            {/* Stats cards */}
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
                {stats.map((stat) => (
                    <div
                        key={stat.title}
                        className="bg-white rounded-lg shadow p-6 text-center hover:shadow-lg transition"
                    >
                        <p className="text-gray-500">{stat.title}</p>
                        <p className="text-2xl font-bold mt-2">{stat.value}</p>
                    </div>
                ))}
            </div>

            {/* Optional quick actions */}
            <div className="mt-8 flex flex-wrap gap-4">
                <button className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700 transition">
                    Browse Courses
                </button>
                <button className="bg-green-600 text-white px-6 py-2 rounded hover:bg-green-700 transition">
                    View Assignments
                </button>
                <button className="bg-purple-600 text-white px-6 py-2 rounded hover:bg-purple-700 transition">
                    Profile Settings
                </button>
            </div>
        </div>
    );
};

export default Dashboard;
