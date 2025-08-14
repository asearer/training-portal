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

  useEffect(() => {
    const fetchProfile = async () => {
      setLoading(true);
      setError(null);
      try {
        const token = localStorage.getItem("token");
        // Assume backend can get user info from token (e.g., /me or /user/:id)
        const res = await axios.get("/user/me", {
          headers: token ? { Authorization: `Bearer ${token}` } : {},
        });
        setUser(res.data);
        setForm({ name: res.data.name, email: res.data.email });
      } catch (err: any) {
        setError(
          err?.response?.data?.error ||
            "Failed to load profile. Please try again."
        );
      } finally {
        setLoading(false);
      }
    };
    fetchProfile();
  }, []);

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
      setError(
        err?.response?.data?.error ||
          "Failed to update profile. Please try again."
      );
    } finally {
      setLoading(false);
    }
  };

  // Password change handlers
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
      setPwError(
        err?.response?.data?.error ||
          "Failed to update password. Please try again."
      );
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-50">
        <div className="text-lg">Loading...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-50">
        <div className="bg-white p-8 rounded shadow-md w-full max-w-md text-center">
          <div className="text-red-600 mb-4">{error}</div>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-50 py-8">
      <div className="w-full max-w-md bg-white rounded shadow p-6">
        <h2 className="text-2xl font-bold mb-6 text-center">My Profile</h2>
        {success && (
          <div className="mb-4 text-green-600 text-center">{success}</div>
        )}
        {!editMode ? (
          <>
            <div className="mb-4">
              <div className="font-medium">Name:</div>
              <div>{user?.name}</div>
            </div>
            <div className="mb-4">
              <div className="font-medium">Email:</div>
              <div>{user?.email}</div>
            </div>
            <div className="mb-4">
              <div className="font-medium">Role:</div>
              <div>{user?.role}</div>
            </div>
            <button
              className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition mr-2"
              onClick={handleEdit}
            >
              Edit Profile
            </button>
            <button
              className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition ml-2"
              onClick={() => setPwMode(true)}
            >
              Change Password
            </button>
          </>
        ) : (
          <form onSubmit={handleSave}>
            <div className="mb-4">
              <label className="block mb-1 font-medium">Name</label>
              <input
                type="text"
                className="w-full border px-3 py-2 rounded"
                value={form.name}
                onChange={e => setForm({ ...form, name: e.target.value })}
                required
                autoFocus
              />
            </div>
            <div className="mb-4">
              <label className="block mb-1 font-medium">Email</label>
              <input
                type="email"
                className="w-full border px-3 py-2 rounded"
                value={form.email}
                onChange={e => setForm({ ...form, email: e.target.value })}
                required
              />
            </div>
            <button
              type="submit"
              className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition mr-2"
              disabled={loading}
            >
              Save
            </button>
            <button
              type="button"
              className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition ml-2"
              onClick={handleCancel}
              disabled={loading}
            >
              Cancel
            </button>
          </form>
        )}

        {/* Password change form */}
        {pwMode && (
          <form onSubmit={handlePwChange} className="mt-8">
            <h3 className="text-lg font-semibold mb-4 text-center">
              Change Password
            </h3>
            {pwError && (
              <div className="mb-4 text-red-600 text-center">{pwError}</div>
            )}
            {pwSuccess && (
              <div className="mb-4 text-green-600 text-center">{pwSuccess}</div>
            )}
            <div className="mb-4">
              <label className="block mb-1 font-medium">Current Password</label>
              <input
                type="password"
                className="w-full border px-3 py-2 rounded"
                value={pwForm.oldPassword}
                onChange={e =>
                  setPwForm({ ...pwForm, oldPassword: e.target.value })
                }
                required
                autoFocus
              />
            </div>
            <div className="mb-4">
              <label className="block mb-1 font-medium">New Password</label>
              <input
                type="password"
                className="w-full border px-3 py-2 rounded"
                value={pwForm.newPassword}
                onChange={e =>
                  setPwForm({ ...pwForm, newPassword: e.target.value })
                }
                required
              />
            </div>
            <div className="mb-6">
              <label className="block mb-1 font-medium">Confirm New Password</label>
              <input
                type="password"
                className="w-full border px-3 py-2 rounded"
                value={pwForm.confirm}
                onChange={e =>
                  setPwForm({ ...pwForm, confirm: e.target.value })
                }
                required
              />
            </div>
            <button
              type="submit"
              className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition mr-2"
              disabled={loading}
            >
              Change Password
            </button>
            <button
              type="button"
              className="bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition ml-2"
              onClick={() => setPwMode(false)}
              disabled={loading}
            >
              Cancel
            </button>
          </form>
        )}
      </div>
    </div>
  );
};

export default Profile;
