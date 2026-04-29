import { DashboardShell } from "@/app/dashboard-shell";

export default function NotificationsPage() {
  return (
    <DashboardShell>
      <div className="flex-1 px-6 py-8 lg:px-8">
        <section className="rounded-3xl border border-white/70 bg-white p-6 shadow-[0_20px_60px_-32px_rgba(15,23,42,0.25)]">
          <p className="text-sm font-medium uppercase tracking-[0.2em] text-indigo-600">
            Notifications
          </p>
          <h1 className="mt-3 text-3xl font-semibold tracking-tight text-slate-950">
            Notifications page
          </h1>
          <p className="mt-3 text-sm leading-6 text-slate-500">
            This page will show the full notifications list once we wire route-level data.
          </p>
        </section>
      </div>
    </DashboardShell>
  );
}
