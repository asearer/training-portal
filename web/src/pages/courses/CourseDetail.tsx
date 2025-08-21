// src/pages/courses/CourseDetail.tsx
import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

interface Course {
    id: string;
    title: string;
    description: string;
    category: string;
    published: boolean;
}

interface Module {
    id: string;
    courseID: string;
    title: string;
    contentType: string;
    contentURL: string;
    orderIndex: number;
}

const CourseDetail: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [course, setCourse] = useState<Course | null>(null);
    const [modules, setModules] = useState<Module[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchCourseAndModules = async () => {
            setLoading(true);
            setError(null);
            try {
                const token = localStorage.getItem("token");
                const [courseRes, modulesRes] = await Promise.all([
                    axios.get(`/course/${id}`, {
                        headers: token ? { Authorization: `Bearer ${token}` } : {},
                    }),
                    axios.get(`/course/${id}/modules`, {
                        headers: token ? { Authorization: `Bearer ${token}` } : {},
                    }),
                ]);
                setCourse(courseRes.data);
                setModules(modulesRes.data);
            } catch (err: any) {
                setError(
                    err?.response?.data?.error ||
                    "Failed to load course details. Please try again."
                );
            } finally {
                setLoading(false);
            }
        };
        if (id) fetchCourseAndModules();
    }, [id]);

    if (loading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="text-lg font-medium text-gray-700">Loading course details...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="bg-white p-6 rounded shadow w-full max-w-2xl text-center">
                    <p className="text-red-600 font-semibold">{error}</p>
                    <button
                        className="mt-4 bg-gray-200 text-gray-800 px-4 py-2 rounded hover:bg-gray-300 transition"
                        onClick={() => navigate("/courses")}
                    >
                        Back to Courses
                    </button>
                </div>
            </div>
        );
    }

    return (
        <div className="flex flex-col items-center min-h-screen bg-gray-50 py-8 px-4">
            <div className="w-full max-w-2xl bg-white rounded-lg shadow p-6">
                <button
                    className="mb-4 text-blue-600 hover:underline"
                    onClick={() => navigate("/courses")}
                >
                    &larr; Back to Courses
                </button>

                {course && (
                    <>
                        <h2 className="text-2xl font-bold mb-2">{course.title}</h2>
                        <p className="mb-2 text-gray-600">{course.description}</p>
                        <div className="mb-4 text-sm text-gray-400">
                            Category: {course.category} |{" "}
                            {course.published ? (
                                <span className="text-green-600 font-medium">Published</span>
                            ) : (
                                <span className="text-yellow-600 font-medium">Draft</span>
                            )}
                        </div>

                        <h3 className="text-lg font-semibold mb-3">Modules</h3>
                        {modules.length === 0 ? (
                            <div className="text-gray-500">No modules found for this course.</div>
                        ) : (
                            <ul className="space-y-3">
                                {modules.map((mod) => (
                                    <li
                                        key={mod.id}
                                        className="p-3 border rounded-lg bg-gray-50 hover:bg-gray-100 transition"
                                    >
                                        <div className="flex justify-between items-center">
                                            <div>
                                                <div className="font-medium">{mod.title}</div>
                                                <div className="text-xs text-gray-500">
                                                    Type: {mod.contentType}
                                                </div>
                                            </div>
                                            {mod.contentURL && (
                                                <a
                                                    href={mod.contentURL}
                                                    target="_blank"
                                                    rel="noopener noreferrer"
                                                    className="ml-4 text-blue-600 underline text-sm"
                                                >
                                                    View Content
                                                </a>
                                            )}
                                        </div>
                                    </li>
                                ))}
                            </ul>
                        )}
                    </>
                )}
            </div>
        </div>
    );
};

export default CourseDetail;

