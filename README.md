# NotifyFlow

A minimal SaaS-style notification delivery system with a Go backend (HTTP API, scheduler, gRPC worker) and a Next.js frontend. Uses MongoDB for persistence and Resend for email delivery.

**Contents**
- **Backend**: `backend/` — API server, scheduler, gRPC worker, database access, Resend sender
- **Frontend**: `frontend/` — Next.js (App Router) dashboard and create form
- **Proto**: `proto/` — gRPC contract for worker

**Quick links**
- Server bootstrap: [backend/cmd/server/main.go](backend/cmd/server/main.go)
- Worker bootstrap: [backend/cmd/worker/main.go](backend/cmd/worker/main.go)
- Scheduler: [backend/internal/scheduler/scheduler.go](backend/internal/scheduler/scheduler.go)
- Send API handler: [backend/internal/api/send_notification.go](backend/internal/api/send_notification.go)
- Notification model: [backend/internal/models/notification.go](backend/internal/models/notification.go)
- Resend sender: [backend/internal/email/resend_sender.go](backend/internal/email/resend_sender.go)
- Frontend dashboard: [frontend/src/components/DashboardClient.tsx](frontend/src/components/DashboardClient.tsx)
- Frontend API client: [frontend/src/lib/api.ts](frontend/src/lib/api.ts)

**Prerequisites**
- Go 1.20+ (or your system Go)
- Node 18+ and pnpm (for frontend)
- MongoDB running locally (or update `MONGO_URI`)
- Resend account & API key (optional for real sends)

## Environment / Configuration
Create a `.env` file in `backend/` (or set env vars). Important variables:

- `MONGO_URI` — MongoDB connection string (default: `mongodb://localhost:27017`)
- `SERVER_PORT` — API server port (default: `8080`)
- `RESEND_API_KEY` — Resend API key (required to send email via Resend)
- `EMAIL_FROM` — Verified `from` email (Resend sandbox restricts recipients)
- `MAX_RETRIES` — Number of retries for scheduled jobs (default: 3)
- `RETRY_BACKOFF_MS` — Backoff in ms between retry attempts (default: 50)

Example (`backend/.env`):

```env
MONGO_URI=mongodb://localhost:27017
SERVER_PORT=8080
RESEND_API_KEY=re_xxx
EMAIL_FROM=onboarding@yourdomain.com
MAX_RETRIES=3
RETRY_BACKOFF_MS=25
```

## Running Locally
Start MongoDB first (e.g., `mongod` or Docker).

Start the gRPC Worker (handles delivery):

```bash
cd backend
go run cmd/worker/main.go
```

Start the API server (HTTP + scheduler):

```bash
cd backend
go run cmd/server/main.go
```

Start the frontend (dev):

```bash
cd frontend
pnpm install
pnpm dev
# Open http://localhost:3000
```

Build commands:

```bash
# Backend: builds via go tooling
cd backend
go build ./...

# Frontend: production build
cd frontend
pnpm build
```

## API Endpoints
- `POST /send` — Create and send (or schedule) a notification.
  - Payload: `{ "to": "user@example.com", "message": "text", "send_at": "optional ISO timestamp" }`
  - Response: `{ id, status, message }`
- `GET /notifications` — List notifications (used by dashboard polling)
- `GET /notification?id=...` — Get by ID
- `GET /failed` — Get failed notifications

See router definition: [backend/internal/api/routes.go](backend/internal/api/routes.go)

## How it works (flow)
- Immediate send:
  - API validates payload and stores a notification with `status=processing`.
  - API calls the gRPC worker, which performs delivery and updates status (`sent`/`failed`) and records `last_error` when present.
- Scheduled send:
  - API stores `status=scheduled` and `send_at`.
  - Scheduler polls `GetScheduledDue` (Mongo) and attempts delivery via gRPC.
  - Scheduler retries with backoff up to `MAX_RETRIES`, updates `retry` and `last_error`, then sets `failed` on exhaustion.

## UI
- Dashboard polls `GET /notifications` every 5s, displays status, retry count, and last error.
- Create form posts to `POST /send` and shows acceptance message and ID.

Files: [frontend/src/components/NotificationsTable.tsx](frontend/src/components/NotificationsTable.tsx)

## Observability & Logs
- Important errors are logged to stdout. Noisy discovery logs were trimmed during development.
- Use Mongo to inspect stored documents in `notifications` collection.

## Common issues & troubleshooting
- Resend sandbox: if Resend reports `You can only send testing emails to your own email address ...`, either use the verified test recipient or verify a domain at Resend to send to arbitrary recipients. The app records that error in `last_error` and will show it in the UI.
- CORS: The API allows the frontend dev origin (http://localhost:3000) by default in `backend/cmd/server/main.go`. If you change frontends or hosts, update CORS there.
- If the frontend throws hydration mismatches during dev, try disabling interfering browser extensions or ensure server/date code is stable between server and client renders.

## Developer notes
- Database helpers: `backend/internal/db/*` (save, update, queries)
- Service layer: `backend/internal/service` — business logic, status updates
- Worker vs Scheduler: Worker does immediate retries for jobs dispatched to it; Scheduler orchestrates scheduled jobs and will retry across attempts.
- `last_error` is persisted on failures to aid troubleshooting and is shown in the UI.

## Next improvements (suggested)
- Add `Cancel scheduled` endpoint + UI control.
- Add `Retry now` manual action for failed jobs.
- Add tests for scheduler/worker flows.
- Add a README screenshots folder and short demo GIFs for GitHub.

