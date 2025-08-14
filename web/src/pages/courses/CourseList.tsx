import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface Course {
  id: string;
  title: string;
  description: string;
  category: string;
  published: boolean;
}

const CourseList: React.FC = () => {
  const [courses, setCourses] = useState<Course[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCourses = async () => {
      setLoading(true);
      setError(null);
      try {
        const token = localStorage.getItem("token");
        const res = await axios.get("/courses", {
          headers: token ? { Authorization: `Bearer ${token}` } : {},
        });
        setCourses(res.data);
      } catch (err: any) {
        setError(
          err?.response?.data?.error ||
            "Failed to load courses. Please try again."
        );
      } finally {
        setLoading(false);
      }
    };
    fetchCourses();
  }, []);

  const handleCourseClick = (id: string) => {
    navigate(`/course/${id}`);
  };

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-50 py-8">
      <div className="w-full max-w-3xl bg-white rounded shadow p-6">
        <h2 className="text-2xl font-bold mb-6 text-center">Courses</h2>
        {loading && <div className="text-center">Loading...</div>}
        {error && (
          <div className="mb-4 text-red-600 text-center">{error}</div>
        )}
        {!loading && !error && courses.length === 0 && (
          <div className="text-center text-gray-500">No courses found.</div>
        )}
        <ul>
          {courses.map((course) => (
            <li
              key={course.id}
              className="mb-4 p-4 border rounded hover:bg-blue-50 cursor-pointer transition"
              onClick={() => handleCourseClick(course.id)}
            >
              <div className="flex justify-between items-center">
                <div>
                  <h3 className="text-lg font-semibold">{course.title}</h3>
                  <p className="text-gray-600">{course.description}</p>
                  <span className="text-xs text-gray-400">
                    Category: {course.category}
                  </span>
                </div>
                <span
                  className={`ml-4 px-2 py-1 rounded text-xs ${
                    course.published
                      ? "bg-green-100 text-green-700"
                      : "bg-yellow-100 text-yellow-700"
                  }`}
                >
                  {course.published ? "Published" : "Draft"}
                </span>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default CourseList;
