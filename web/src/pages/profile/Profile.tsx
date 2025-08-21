// src/pages/profile/Profile.tsx
import React, { useEffect, useState } from "react";
import axios from "axios";

interface User {
    id: string;
    name: string;
    email: string;
    role: string;
}

const Profile: React.FC = () => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const [editMode, setEditMode] = useState(false);
    const [form, setForm] = useState({ name: "", email: "" });
    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);

    // Password change state
    const [pwMode, setPwMode] = useState(false);
    const [pwForm, setPwForm] = useState({ oldPassword: "", newPassword: "", confirm: "" });
    const [pwError, setPwError] = useState<string | null>(null);
    const [pwSuccess, setPwSuccess] = useState<string | null>(null);

    // Fetch user profile
    useEffect(() => {
        const fetchProfile = async () => {
            setLoading(true);
            setError(null);
            try {
                const token = localStorage.getItem("token");
                const res = await axios.get("/user/me", {
                    headers: token ? { Authorization: `Bearer ${token}` } : {},
                });
                setUser(res.data);
                setForm({ name: res.data.name, email: res.data.email });
            } catch (err: any) {
                setError(err?.response?.data?.error || "Failed to load profile. Please try again.");
            } finally {
                setLoading(false);
            }
        };
        fetchProfile();
    }, []);

    // Edit handlers
    const handleEdit = () => {
        setEditMode(true);
        setSuccess(null);
        setError(null);
    };

    const handleCancel = () => {
        if (user) setForm({ name: user.name, email: user.email });
        setEditMode(false);
        setSuccess(null);
        setError(null);
    };

    const handleSave = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);
        setSuccess(null);
        setLoading(true);
        try {
            const token = localStorage.getItem("token");
            await axios.put(
                `/api/user/${user?.id}`,
                { name: form.name, email: form.email },
                { headers: token ? { Authorization: `Bearer ${token}` } : {} }
            );
            setUser((u) => (u ? { ...u, name: form.name, email: form.email } : u));
            setSuccess("Profile updated successfully.");
            setEditMode(false);
        } catch (err: any) {
            setError(err?.response?.data?.error || "Failed to update profile. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    // Password change handler
    const handlePwChange = async (e: React.FormEvent) => {
        e.preventDefault();
        setPwError(null);
        setPwSuccess(null);

        if (!pwForm.oldPassword || !pwForm.newPassword || !pwForm.confirm) {
            setPwError("All fields are required.");
            return;
        }
        if (pwForm.newPassword !== pwForm.confirm) {
            setPwError("New passwords do not match.");
            return;
        }
        setLoading(true);
        try {
            const token = localStorage.getItem("token");
            await axios.put(
                `/api/user/${user?.id}/password`,
                { oldPassword: pwForm.oldPassword, newPassword: pwForm.newPassword },
                { headers: token ? { Authorization: `Bearer ${token}` } : {} }
            );
            setPwSuccess("Password updated successfully.");
            setPwForm({ oldPassword: "", newPassword: "", confirm: "" });
            setPwMode(false);
        } catch (err: any) {
            setPwError(err?.response?.data?.error || "Failed to update password. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="text-lg font-medium text-gray-700">Loading profile...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md text-center">
                    <p className="text-red-600 font-semibold">{error}</p>
                </div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-gray-50 flex justify-center py-8 px-4">
            <div className="w-full max-w-md bg-white rounded-lg shadow p-6">
                <h2 className="text-2xl font-bold mb-6 text-center">My Profile</h2>

                {success && <p className="mb-4 text-green-600 text-center font-medium">{success}</p>}

                {!editMode ? (
                    <>
                        <div className="space-y-3">
                            <div>
                                <span className="font-semibold">Name:</span> {user?.name}
                            </div>
                            <div>
                                <span className="font-semibold">Email:</span> {user?.email}
                            </div>
                            <div>
                                <span className="font-semibold">Role:</span> {user?.role}
                            </div>
                        </div>

                        <div className="mt-6 flex justify-center gap-3 flex-wrap">
                            <button
                                className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition"
                                onClick={handleEdit}
                            >
                                Edit Profile
                            </button>
                            <button
                                className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition"
                                onClick={() => setPwMode(true)}
                            >
                                Change Password
                            </button>
                        </div>
                    </>
                ) : (
                    <form onSubmit={handleSave} className="space-y-4">
                        <div>
                            <label className="block mb-1 font-medium">Name</label>
                            <input
                                type="text"
                                className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={form.name}
                                onChange={(e) => setForm({ ...form, name: e.target.value })}
                                required
                                autoFocus
                            />
                        </div>

                        <div>
                            <label className="block mb-1 font-medium">Email</label>
                            <input
                                type="email"
                                className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={form.email}
                                onChange={(e) => setForm({ ...form, email: e.target.value })}
                                required
                            />
                        </div>

                        <div className="flex justify-center gap-3 flex-wrap mt-4">
                            <button
                                type="submit"
                                className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition"
                                disabled={loading}
                            >
                                Save
                            </button>
                            <button
                                type="button"
                                className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition"
                                onClick={handleCancel}
                                disabled={loading}
                            >
                                Cancel
                            </button>
                        </div>
                    </form>
                )}

                {/* Password Change Form */}
                {pwMode && (
                    <form onSubmit={handlePwChange} className="mt-8 space-y-4">
                        <h3 className="text-lg font-semibold text-center mb-4">Change Password</h3>

                        {pwError && <p className="text-red-600 text-center font-medium">{pwError}</p>}
                        {pwSuccess && <p className="text-green-600 text-center font-medium">{pwSuccess}</p>}

                        <div>
                            <label className="block mb-1 font-medium">Current Password</label>
                            <input
                                type="password"
                                className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={pwForm.oldPassword}
                                onChange={(e) => setPwForm({ ...pwForm, oldPassword: e.target.value })}
                                required
                                autoFocus
                            />
                        </div>

                        <div>
                            <label className="block mb-1 font-medium">New Password</label>
                            <input
                                type="password"
                                className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={pwForm.newPassword}
                                onChange={(e) => setPwForm({ ...pwForm, newPassword: e.target.value })}
                                required
                            />
                        </div>

                        <div>
                            <label className="block mb-1 font-medium">Confirm New Password</label>
                            <input
                                type="password"
                                className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={pwForm.confirm}
                                onChange={(e) => setPwForm({ ...pwForm, confirm: e.target.value })}
                                required
                            />
                        </div>

                        <div className="flex justify-center gap-3 flex-wrap mt-4">
                            <button
                                type="submit"
                                className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition"
                                disabled={loading}
                            >
                                Change Password
                            </button>
                            <button
                                type="button"
                                className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition"
                                onClick={() => setPwMode(false)}
                                disabled={loading}
                            >
                                Cancel
                            </button>
                        </div>
                    </form>
                )}
            </div>
        </div>
    );
};

export default Profile;

