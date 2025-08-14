# Training Portal

A full-featured, scalable Learning Management System (LMS) built with Go (Fiber) for the backend and React + TypeScript for the frontend. This project is designed for organizations to deliver, manage, and track employee or student training, with a focus on modularity, extensibility, and modern best practices.

---

## 🚀 Features

- **User Authentication & Authorization** (JWT, role-based access)
- **Course Management** (CRUD, categories, publishing)
- **Module Management** (videos, PDFs, ordering)
- **User Enrollment** (enroll/unenroll, progress tracking)
- **Quizzes & Assessments** (auto-grading, feedback)
- **Certificates** (PDF/badge generation, download/share)
- **Admin Panel** (user, course, module, analytics management)
- **Notifications** (email, in-app, reminders)
- **Forums & Messaging** (discussions, direct messages)
- **Analytics Dashboard** (engagement, completion rates)
- **Bulk User Management** (CSV import/export)
- **Profile Management** (avatar, achievements, password reset)
- **Accessibility & Internationalization** (i18n-ready, WCAG)
- **Mobile Responsive** (Tailwind CSS, mobile-first)
- **DevOps Ready** (Docker, CI/CD, environment configs)

---

## 🗂️ Directory Structure

See [`directory-structure.md`](./directory-structure.md) for a detailed breakdown.

Key directories:

- `cmd/server/` — Backend entry point
- `internal/` — Clean architecture: domain, usecase, interface, repository
- `migrations/` — SQL migrations
- `configs/` — App configuration
- `web/` — React + TypeScript frontend
- `docs/` — API docs, architecture diagrams
- `scripts/` — DevOps scripts

---

## 🛠️ Getting Started

### Prerequisites

- Go 1.20+
- Node.js 18+
- PostgreSQL 13+
- (Optional) Docker

### Backend Setup

1. **Install dependencies:**
   ```sh
   cd training-portal
   go mod tidy
   ```

2. **Configure environment:**
   - Copy `configs/app.yaml` and/or `.env` and set DB/JWT values.

3. **Run migrations:**
   ```sh
   # Use your preferred migration tool or psql
   ```

4. **Start the server:**
   ```sh
   go run cmd/server/main.go
   ```

### Frontend Setup

1. **Install dependencies:**
   ```sh
   cd web
   npm install
   ```

2. **Configure API URL:**
   - Edit `.env` if backend is not on `localhost:3000`.

3. **Start the frontend:**
   ```sh
   npm run dev
   ```

---

## 🧩 Roadmap

See [`roadmap.md`](./roadmap.md) for planned features and phases, including:

- Enrollment & progress tracking
- Quizzes & certificates
- Bulk user management
- Analytics dashboard
- Notifications, forums, messaging
- Accessibility, i18n, mobile, security, DevOps

---

## 🧪 Testing

- Register/login as a user and admin
- Enroll in courses, complete modules, take quizzes
- Use the admin panel for CRUD and analytics
- Test notifications, forums, and messaging
- Try password reset and profile updates

---

## 📦 Deployment

- Dockerize backend and frontend for production
- Use CI/CD for automated testing and deployment
- Set secure environment variables for production
- Deploy to cloud platforms like AWS, GCP, or Azure

---

## 📚 Documentation

- [API Docs](./docs/)
- [Architecture Diagrams](./docs/)
- [Roadmap](./roadmap.md)
- [Directory Structure](./directory-structure.md)

---

_This project is under active development. See the roadmap for upcoming features!_
