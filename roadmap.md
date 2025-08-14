# Training Portal Roadmap

This document outlines the planned features, phases, and architectural stubs for the continued development of the Training Portal. As the project grows, this roadmap will help guide implementation, prioritize work, and communicate progress.

---

## 🚩 **Phases & Features**

### **Phase 1: Learning Core**
- [ ] **Course Enrollment**
  - Users can enroll/unenroll in courses.
  - Only enrolled users can access course content.
- [ ] **Progress Tracking**
  - Track modules completed, quiz scores, and overall course progress.
  - Display progress bars and completion status.
- [ ] **Quizzes & Assessments**
  - Add quizzes to modules/courses.
  - Auto-grade multiple choice/short answer.
  - Show results and feedback.
- [ ] **Certificates**
  - Generate certificates (PDF or digital badge) upon course completion.
  - Allow users to download/share certificates.

---

### **Phase 2: Admin & Analytics**
- [ ] **Bulk User Management**
  - Import/export users via CSV.
  - Assign users to courses in bulk.
- [ ] **Analytics Dashboard**
  - Visualize user engagement, course popularity, completion rates.
  - Export analytics data.
- [ ] **Content Management**
  - Rich text/markdown editor for course/module content.
  - Drag-and-drop module reordering.
- [ ] **Role Management**
  - Assign/revoke admin or trainer roles.
  - Fine-grained permissions (e.g., trainers can only edit their own courses).

---

### **Phase 3: Social & UX**
- [ ] **Notifications**
  - Email or in-app notifications for new courses, assignments, deadlines.
  - Reminders for incomplete courses.
- [ ] **Forums/Discussions**
  - Discussion forums or Q&A per course/module.
- [ ] **Messaging**
  - Direct messaging between users and trainers.
- [ ] **Search & Filtering**
  - Search courses by title, category, instructor.
  - Filter by enrolled, completed, recommended.
- [ ] **Profile Enhancements**
  - Avatar upload.
  - View course history, achievements, certificates.

---

### **Phase 4: Technical & Quality**
- [ ] **Accessibility & Internationalization**
  - Ensure WCAG accessibility compliance.
  - Add i18n support for multiple languages.
- [ ] **Mobile Responsiveness**
  - Ensure all pages are mobile-friendly.
  - Consider mobile app (React Native/Flutter).
- [ ] **Security Enhancements**
  - Two-factor authentication (2FA).
  - Account lockout after repeated failed logins.
  - Audit logs for admin actions.
- [ ] **DevOps & Deployment**
  - Dockerize frontend and backend.
  - CI/CD pipelines for automated testing/deployment.
  - Environment-specific configs for staging/production.

---

## 🏗️ **Stubs & Directory Structure**

### **Backend Stubs**
- `internal/domain/enrollment/` — Enrollment models/interfaces
- `internal/domain/progress/` — Progress tracking models
- `internal/domain/quiz/` — Quiz/question/answer models
- `internal/domain/certificate/` — Certificate models
- `internal/domain/notification/` — Notification models
- `internal/domain/forum/` — Forum/discussion models
- `internal/domain/message/` — Messaging models
- `internal/domain/analytics/` — Analytics models
- `internal/domain/role/` — Role/permission models

- `internal/usecase/enrollment/` — Enrollment logic
- `internal/usecase/progress/` — Progress logic
- `internal/usecase/quiz/` — Quiz logic
- `internal/usecase/certificate/` — Certificate logic
- `internal/usecase/notification/` — Notification logic
- `internal/usecase/forum/` — Forum logic
- `internal/usecase/message/` — Messaging logic
- `internal/usecase/analytics/` — Analytics logic
- `internal/usecase/role/` — Role logic

- `internal/interface/http/handler/enrollment.go`
- `internal/interface/http/handler/progress.go`
- `internal/interface/http/handler/quiz.go`
- `internal/interface/http/handler/certificate.go`
- `internal/interface/http/handler/notification.go`
- `internal/interface/http/handler/forum.go`
- `internal/interface/http/handler/message.go`
- `internal/interface/http/handler/analytics.go`
- `internal/interface/http/handler/role.go`

- `internal/interface/repository/postgres/enrollment.go`
- `internal/interface/repository/postgres/progress.go`
- `internal/interface/repository/postgres/quiz.go`
- `internal/interface/repository/postgres/certificate.go`
- `internal/interface/repository/postgres/notification.go`
- `internal/interface/repository/postgres/forum.go`
- `internal/interface/repository/postgres/message.go`
- `internal/interface/repository/postgres/analytics.go`
- `internal/interface/repository/postgres/role.go`

- `migrations/` — Add new tables for all new features

---

### **Frontend Stubs**
- `web/src/pages/enrollment/` — Enrollment UI (enroll/unenroll, list enrolled)
- `web/src/pages/progress/` — Progress bars, completion status
- `web/src/pages/quiz/` — Quiz taking, results, feedback
- `web/src/pages/certificate/` — Certificate download/share
- `web/src/pages/analytics/` — Admin analytics dashboard
- `web/src/pages/notifications/` — Notification center
- `web/src/pages/forum/` — Course/module forums
- `web/src/pages/messages/` — Messaging UI
- `web/src/pages/search/` — Search/filter UI
- `web/src/pages/profile/` — Avatar upload, achievements, history

- `web/src/components/` — Shared UI components (progress bar, quiz, certificate, notification, etc.)
- `web/src/services/` — API logic for new features
- `web/src/types/` — Types/interfaces for new features
- `web/src/hooks/` — Custom hooks for notifications, progress, etc.
- `web/src/utils/` — Utility functions (i18n, accessibility, etc.)
