import React, { useEffect, useState } from "react";
import axios from "axios";

// --- Types ---
interface User {
  id: string;
  name: string;
  email: string;
  role: string;
}

interface Course {
  id: string;
  title: string;
  description: string;
  category: string;
  published: boolean;
  createdBy: string;
}

interface Module {
  id: string;
  courseID: string;
  title: string;
  contentType: string;
  contentURL: string;
  orderIndex: number;
}

// --- Admin Panel ---
const AdminPanel: React.FC = () => {
  // Tabs: "users" | "courses" | "modules"
  const [tab, setTab] = useState<"users" | "courses" | "modules">("users");

  // Data
  const [users, setUsers] = useState<User[]>([]);
  const [courses, setCourses] = useState<Course[]>([]);
  const [modules, setModules] = useState<Module[]>([]);

  // For module CRUD, need to select a course
  const [selectedCourseId, setSelectedCourseId] = useState<string>("");

  // Loading and error
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // --- Fetch Data ---
  useEffect(() => {
    const token = localStorage.getItem("token");
    setError(null);
    setLoading(true);

    const fetchAll = async () => {
      try {
        if (tab === "users") {
          const res = await axios.get("/users", {
            headers: token ? { Authorization: `Bearer ${token}` } : {},
          });
          setUsers(res.data);
        } else if (tab === "courses") {
          const res = await axios.get("/courses", {
            headers: token ? { Authorization: `Bearer ${token}` } : {},
          });
          setCourses(res.data);
        } else if (tab === "modules" && selectedCourseId) {
          const res = await axios.get(`/course/${selectedCourseId}/modules`, {
            headers: token ? { Authorization: `Bearer ${token}` } : {},
          });
          setModules(res.data);
        }
      } catch (err: any) {
        setError(
          err?.response?.data?.error ||
            "Failed to load data. Please try again."
        );
      } finally {
        setLoading(false);
      }
    };

    fetchAll();
    // eslint-disable-next-line
  }, [tab, selectedCourseId]);

  // --- User CRUD ---
  const handleDeleteUser = async (id: string) => {
    if (!window.confirm("Delete this user?")) return;
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem("token");
      await axios.delete(`/api/user/${id}`, {
        headers: token ? { Authorization: `Bearer ${token}` } : {},
      });
      setUsers(users.filter((u) => u.id !== id));
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to delete user.");
    } finally {
      setLoading(false);
    }
  };

  // --- Course CRUD ---
  const [newCourse, setNewCourse] = useState<Partial<Course>>({});
  const handleCreateCourse = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem("token");
      const res = await axios.post(
        "/api/course",
        newCourse,
        { headers: token ? { Authorization: `Bearer ${token}` } : {} }
      );
      setCourses([...courses, res.data]);
      setNewCourse({});
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to create course.");
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteCourse = async (id: string) => {
    if (!window.confirm("Delete this course?")) return;
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem("token");
      await axios.delete(`/api/course/${id}`, {
        headers: token ? { Authorization: `Bearer ${token}` } : {},
      });
      setCourses(courses.filter((c) => c.id !== id));
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to delete course.");
    } finally {
      setLoading(false);
    }
  };

  // --- Module CRUD ---
  const [newModule, setNewModule] = useState<Partial<Module>>({});
  const handleCreateModule = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedCourseId) {
      setError("Select a course first.");
      return;
    }
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem("token");
      const res = await axios.post(
        "/api/module",
        { ...newModule, courseID: selectedCourseId },
        { headers: token ? { Authorization: `Bearer ${token}` } : {} }
      );
      setModules([...modules, res.data]);
      setNewModule({});
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to create module.");
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteModule = async (id: string) => {
    if (!window.confirm("Delete this module?")) return;
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem("token");
      await axios.delete(`/api/module/${id}`, {
        headers: token ? { Authorization: `Bearer ${token}` } : {},
      });
      setModules(modules.filter((m) => m.id !== id));
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to delete module.");
    } finally {
      setLoading(false);
    }
  };

  // --- Render ---
  return (
    <div className="min-h-screen bg-gray-50 py-8 px-4">
      <div className="max-w-5xl mx-auto bg-white rounded shadow p-6">
        <h1 className="text-3xl font-bold mb-6 text-center">Admin Panel</h1>
        <div className="flex justify-center mb-6">
          <button
            className={`px-4 py-2 rounded-l ${tab === "users" ? "bg-blue-600 text-white" : "bg-gray-200"}`}
            onClick={() => setTab("users")}
          >
            Users
          </button>
          <button
            className={`px-4 py-2 ${tab === "courses" ? "bg-blue-600 text-white" : "bg-gray-200"}`}
            onClick={() => setTab("courses")}
          >
            Courses
          </button>
          <button
            className={`px-4 py-2 rounded-r ${tab === "modules" ? "bg-blue-600 text-white" : "bg-gray-200"}`}
            onClick={() => setTab("modules")}
          >
            Modules
          </button>
        </div>
        {error && <div className="mb-4 text-red-600 text-center">{error}</div>}
        {loading && <div className="mb-4 text-center">Loading...</div>}

        {/* --- Users Tab --- */}
        {tab === "users" && (
          <div>
            <h2 className="text-xl font-semibold mb-4">All Users</h2>
            <table className="w-full mb-4 border">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 border">Name</th>
                  <th className="p-2 border">Email</th>
                  <th className="p-2 border">Role</th>
                  <th className="p-2 border">Actions</th>
                </tr>
              </thead>
              <tbody>
                {users.map((u) => (
                  <tr key={u.id}>
                    <td className="p-2 border">{u.name}</td>
                    <td className="p-2 border">{u.email}</td>
                    <td className="p-2 border">{u.role}</td>
                    <td className="p-2 border">
                      <button
                        className="text-red-600 hover:underline"
                        onClick={() => handleDeleteUser(u.id)}
                      >
                        Delete
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
            <div className="text-sm text-gray-500">
              (User creation is only available via registration.)
            </div>
          </div>
        )}

        {/* --- Courses Tab --- */}
        {tab === "courses" && (
          <div>
            <h2 className="text-xl font-semibold mb-4">All Courses</h2>
            <form onSubmit={handleCreateCourse} className="mb-6 flex flex-wrap gap-2">
              <input
                type="text"
                placeholder="Title"
                className="border px-2 py-1 rounded"
                value={newCourse.title || ""}
                onChange={e => setNewCourse({ ...newCourse, title: e.target.value })}
                required
              />
              <input
                type="text"
                placeholder="Description"
                className="border px-2 py-1 rounded"
                value={newCourse.description || ""}
                onChange={e => setNewCourse({ ...newCourse, description: e.target.value })}
                required
              />
              <input
                type="text"
                placeholder="Category"
                className="border px-2 py-1 rounded"
                value={newCourse.category || ""}
                onChange={e => setNewCourse({ ...newCourse, category: e.target.value })}
                required
              />
              <button
                type="submit"
                className="bg-green-600 text-white px-3 py-1 rounded"
                disabled={loading}
              >
                Add Course
              </button>
            </form>
            <table className="w-full mb-4 border">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 border">Title</th>
                  <th className="p-2 border">Description</th>
                  <th className="p-2 border">Category</th>
                  <th className="p-2 border">Published</th>
                  <th className="p-2 border">Actions</th>
                </tr>
              </thead>
              <tbody>
                {courses.map((c) => (
                  <tr key={c.id}>
                    <td className="p-2 border">{c.title}</td>
                    <td className="p-2 border">{c.description}</td>
                    <td className="p-2 border">{c.category}</td>
                    <td className="p-2 border">
                      {c.published ? (
                        <span className="text-green-600">Yes</span>
                      ) : (
                        <span className="text-yellow-600">No</span>
                      )}
                    </td>
                    <td className="p-2 border">
                      <button
                        className="text-red-600 hover:underline"
                        onClick={() => handleDeleteCourse(c.id)}
                      >
                        Delete
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        {/* --- Modules Tab --- */}
        {tab === "modules" && (
          <div>
            <h2 className="text-xl font-semibold mb-4">Modules</h2>
            <div className="mb-4">
              <label className="mr-2 font-medium">Select Course:</label>
              <select
                className="border px-2 py-1 rounded"
                value={selectedCourseId}
                onChange={e => setSelectedCourseId(e.target.value)}
              >
                <option value="">-- Select --</option>
                {courses.map((c) => (
                  <option key={c.id} value={c.id}>
                    {c.title}
                  </option>
                ))}
              </select>
            </div>
            {selectedCourseId && (
              <>
                <form onSubmit={handleCreateModule} className="mb-6 flex flex-wrap gap-2">
                  <input
                    type="text"
                    placeholder="Title"
                    className="border px-2 py-1 rounded"
                    value={newModule.title || ""}
                    onChange={e => setNewModule({ ...newModule, title: e.target.value })}
                    required
                  />
                  <input
                    type="text"
                    placeholder="Content Type"
                    className="border px-2 py-1 rounded"
                    value={newModule.contentType || ""}
                    onChange={e => setNewModule({ ...newModule, contentType: e.target.value })}
                    required
                  />
                  <input
                    type="text"
                    placeholder="Content URL"
                    className="border px-2 py-1 rounded"
                    value={newModule.contentURL || ""}
                    onChange={e => setNewModule({ ...newModule, contentURL: e.target.value })}
                  />
                  <input
                    type="number"
                    placeholder="Order"
                    className="border px-2 py-1 rounded"
                    value={newModule.orderIndex || ""}
                    onChange={e => setNewModule({ ...newModule, orderIndex: Number(e.target.value) })}
                    required
                  />
                  <button
                    type="submit"
                    className="bg-green-600 text-white px-3 py-1 rounded"
                    disabled={loading}
                  >
                    Add Module
                  </button>
                </form>
                <table className="w-full mb-4 border">
                  <thead>
                    <tr className="bg-gray-100">
                      <th className="p-2 border">Title</th>
                      <th className="p-2 border">Type</th>
                      <th className="p-2 border">URL</th>
                      <th className="p-2 border">Order</th>
                      <th className="p-2 border">Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {modules.map((m) => (
                      <tr key={m.id}>
                        <td className="p-2 border">{m.title}</td>
                        <td className="p-2 border">{m.contentType}</td>
                        <td className="p-2 border">
                          {m.contentURL ? (
                            <a
                              href={m.contentURL}
                              target="_blank"
                              rel="noopener noreferrer"
                              className="text-blue-600 underline"
                            >
                              Link
                            </a>
                          ) : (
                            "-"
                          )}
                        </td>
                        <td className="p-2 border">{m.orderIndex}</td>
                        <td className="p-2 border">
                          <button
                            className="text-red-600 hover:underline"
                            onClick={() => handleDeleteModule(m.id)}
                          >
                            Delete
                          </button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default AdminPanel;
