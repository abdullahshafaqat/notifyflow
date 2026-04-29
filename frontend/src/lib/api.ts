export type Notification = {
  id: string;
  to: string;
  message: string;
  status: "sent" | "scheduled" | "failed" | "processing" | string;
  retry: number;
  send_at?: string;
  last_error?: string;
};

export type DashboardStats = {
  total: number;
  sent: number;
  failed: number;
  scheduled: number;
};

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";

export async function getNotifications(): Promise<Notification[]> {
  const response = await fetch(`${API_BASE_URL}/notifications`, {
    cache: "no-store",
  });

  if (!response.ok) {
    throw new Error("Failed to load notifications");
  }

  const data = await response.json();
  return Array.isArray(data) ? data : [];
}

export async function createNotification(payload: {
  to: string;
  message: string;
  send_at?: string;
}) {
  const response = await fetch(`${API_BASE_URL}/send`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data?.message ?? "Failed to create notification");
  }

  return data as {
    id?: string;
    status: string;
    message: string;
  };
}

export function getStatusTone(status: Notification["status"]) {
  const normalized = status?.toLowerCase();

  switch (normalized) {
    case "sent":
      return "bg-emerald-50 text-emerald-700 ring-1 ring-emerald-200";
    case "failed":
      return "bg-rose-50 text-rose-700 ring-1 ring-rose-200";
    case "scheduled":
      return "bg-amber-50 text-amber-700 ring-1 ring-amber-200";
    case "processing":
      return "bg-sky-50 text-sky-700 ring-1 ring-sky-200";
    default:
      return "bg-slate-100 text-slate-600 ring-1 ring-slate-200";
  }
}
