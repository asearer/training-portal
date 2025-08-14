import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const PasswordReset: React.FC = () => {
  const [step, setStep] = useState<"request" | "reset" | "done">("request");
  const [email, setEmail] = useState("");
  const [token, setToken] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // Step 1: Request password reset (send email)
  const handleRequest = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);
    setLoading(true);
    try {
      await axios.post("/password-reset/request", { email });
      setSuccess("If your email exists, a reset link or code has been sent.");
      setStep("reset");
    } catch (err: any) {
      setError(
        err?.response?.data?.error ||
          "Failed to request password reset. Please try again."
      );
    } finally {
      setLoading(false);
    }
  };

  // Step 2: Set new password using token/code
  const handleReset = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    if (!token || !newPassword || !confirm) {
      setError("All fields are required.");
      return;
    }
    if (newPassword !== confirm) {
      setError("Passwords do not match.");
      return;
    }
    setLoading(true);
    try {
      await axios.post("/password-reset/confirm", {
        email,
        token,
        newPassword,
      });
      setSuccess("Password reset successful! Redirecting to login...");
      setStep("done");
      setTimeout(() => {
        navigate("/login");
      }, 1500);
    } catch (err: any) {
      setError(
        err?.response?.data?.error ||
          "Failed to reset password. Please check your code and try again."
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded shadow-md w-full max-w-sm">
        <h2 className="text-2xl font-bold mb-6 text-center">Password Reset</h2>
        {error && (
          <div className="mb-4 text-red-600 text-center">{error}</div>
        )}
        {success && (
          <div className="mb-4 text-green-600 text-center">{success}</div>
        )}

        {step === "request" && (
          <form onSubmit={handleRequest}>
            <div className="mb-4">
              <label className="block mb-1 font-medium">Email</label>
              <input
                type="email"
                className="w-full border px-3 py-2 rounded"
                value={email}
                onChange={e => setEmail(e.target.value)}
                required
                autoFocus
              />
            </div>
            <button
              type="submit"
              className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
              disabled={loading}
            >
              {loading ? "Sending..." : "Send Reset Link"}
            </button>
            <div className="mt-4 text-center text-sm">
              <span
                className="text-blue-600 hover:underline cursor-pointer"
                onClick={() => navigate("/login")}
              >
                Back to Login
              </span>
            </div>
          </form>
        )}

        {step === "reset" && (
          <form onSubmit={handleReset}>
            <div className="mb-4">
              <label className="block mb-1 font-medium">Reset Code / Token</label>
              <input
                type="text"
                className="w-full border px-3 py-2 rounded"
                value={token}
                onChange={e => setToken(e.target.value)}
                required
                autoFocus
              />
            </div>
            <div className="mb-4">
              <label className="block mb-1 font-medium">New Password</label>
              <input
                type="password"
                className="w-full border px-3 py-2 rounded"
                value={newPassword}
                onChange={e => setNewPassword(e.target.value)}
                required
              />
            </div>
            <div className="mb-6">
              <label className="block mb-1 font-medium">Confirm Password</label>
              <input
                type="password"
                className="w-full border px-3 py-2 rounded"
                value={confirm}
                onChange={e => setConfirm(e.target.value)}
                required
              />
            </div>
            <button
              type="submit"
              className="w-full bg-green-600 text-white py-2 rounded hover:bg-green-700 transition"
              disabled={loading}
            >
              {loading ? "Resetting..." : "Reset Password"}
            </button>
            <div className="mt-4 text-center text-sm">
              <span
                className="text-blue-600 hover:underline cursor-pointer"
                onClick={() => setStep("request")}
              >
                Back to Request
              </span>
            </div>
          </form>
        )}

        {step === "done" && (
          <div className="text-center">
            <div className="mb-4">Password reset successful!</div>
            <button
              className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
              onClick={() => navigate("/login")}
            >
              Go to Login
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default PasswordReset;
