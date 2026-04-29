'use client';

import { useEffect, useMemo, useState } from "react";

import {
  getNotifications,
  getStatusTone,
  type DashboardStats,
  type Notification,
} from "@/lib/api";

import { NotificationComposer } from "@/components/NotificationComposer";
import { NotificationsTable } from "@/components/NotificationsTable";
import { StatCard } from "@/components/StatCard";

type DashboardClientProps = {
  initialNotifications?: Notification[] | null;
};

function formatDate(value?: string) {
  if (!value) return "-";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return value;
  return new Intl.DateTimeFormat("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}

function deriveStats(notifications: Notification[]): DashboardStats {
  return notifications.reduce(
    (accumulator, notification) => {
      accumulator.total += 1;
      const normalized = notification.status?.toLowerCase();
      if (normalized === "sent") accumulator.sent += 1;
      if (normalized === "failed") accumulator.failed += 1;
      if (normalized === "scheduled") accumulator.scheduled += 1;
      return accumulator;
    },
    { total: 0, sent: 0, failed: 0, scheduled: 0 },
  );
}

export function DashboardClient({ initialNotifications }: DashboardClientProps) {
  const [notifications, setNotifications] = useState<Array<Notification>>(initialNotifications ?? []);
  const [isRefreshing, setIsRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [toast, setToast] = useState<string | null>(null);

  useEffect(() => {
    const interval = window.setInterval(async () => {
      setIsRefreshing(true);
      try {
        const nextNotifications = await getNotifications();
        setNotifications(nextNotifications);
        setError(null);
      } catch {
        setError("Live refresh failed. Showing the latest loaded data.");
      } finally {
        setIsRefreshing(false);
      }
    }, 5000);

    return () => window.clearInterval(interval);
  }, []);

  const stats = useMemo(() => deriveStats(notifications ?? []), [notifications]);

  const tableRows = (notifications ?? []).map((notification) => ({
    recipient: notification.to,
    message: notification.message,
    status: notification.status.toLowerCase() as "sent" | "scheduled" | "failed" | "processing",
    retries: notification.retry,
    scheduledAt: formatDate(notification.send_at),
    statusTone: getStatusTone(notification.status),
    lastError: notification.last_error ?? "",
  }));

  return (
    <div className="flex-1 px-6 py-8 lg:px-8">
      <section className="mb-8 rounded-3xl border border-white/70 bg-white/80 p-6 shadow-[0_20px_60px_-30px_rgba(15,23,42,0.2)] backdrop-blur">
        <div className="flex flex-col gap-3 md:flex-row md:items-end md:justify-between">
          <div>
            <p className="text-sm font-medium uppercase tracking-[0.24em] text-indigo-600">
              Dashboard
            </p>
            <h1 className="mt-2 text-3xl font-semibold tracking-tight text-slate-950 md:text-4xl">
              Notifications at a glance
            </h1>
            <p className="mt-2 max-w-2xl text-sm leading-6 text-slate-600 md:text-base">
              Monitor sent, scheduled, and failed notifications from one clean SaaS-style view.
            </p>
          </div>

          <div className="rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-600">
            <span className="font-medium text-slate-900">Workspace:</span> NotifyFlow
            {isRefreshing ? <span className="ml-2 text-indigo-600">Refreshing…</span> : null}
          </div>
        </div>
      </section>

      {error ? (
        <div className="mb-6 rounded-2xl border border-amber-200 bg-amber-50 px-4 py-3 text-sm text-amber-800">
          {error}
        </div>
      ) : null}

      {toast ? (
        <div className="mb-6 rounded-2xl border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-800">
          {toast}
        </div>
      ) : null}

      <section className="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
        <StatCard label="Total Notifications" value={stats.total.toString()} trend="Live" />
        <StatCard label="Sent" value={stats.sent.toString()} trend="Live" />
        <StatCard label="Failed" value={stats.failed.toString()} trend="Live" />
        <StatCard label="Scheduled" value={stats.scheduled.toString()} trend="Live" />
      </section>

      <section className="mt-8">
        <NotificationsTable rows={tableRows} />
      </section>

      <div className="mt-8">
        <NotificationComposer
          onSuccess={(message) => {
            setToast(message);
            setError(null);
            void getNotifications().then(setNotifications).catch(() => undefined);
          }}
          onError={(message) => {
            setToast(null);
            setError(message);
          }}
        />
      </div>
    </div>
  );
}
