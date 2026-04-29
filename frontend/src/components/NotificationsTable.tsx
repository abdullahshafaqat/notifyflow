export type NotificationRow = {
  recipient: string;
  message: string;
  status: "sent" | "scheduled" | "failed" | "processing";
  retries: number;
  scheduledAt: string;
  statusTone?: string;
  lastError?: string;
};

type NotificationsTableProps = {
  rows: NotificationRow[];
};

export function NotificationsTable({ rows }: NotificationsTableProps) {
  return (
    <section className="rounded-3xl border border-white/70 bg-white p-6 shadow-[0_20px_60px_-32px_rgba(15,23,42,0.25)]">
      <div className="flex flex-col gap-2 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p className="text-sm font-medium uppercase tracking-[0.2em] text-indigo-600">
            Recent activity
          </p>
          <h2 className="mt-1 text-2xl font-semibold tracking-tight text-slate-950">
            Notifications
          </h2>
        </div>
        <p className="text-sm text-slate-500">Live data with 5-second refresh</p>
      </div>

      <div className="mt-6 overflow-hidden rounded-2xl border border-slate-200">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-slate-500">
            <tr>
              <th className="px-5 py-4 font-medium">Recipient</th>
              <th className="px-5 py-4 font-medium">Message</th>
              <th className="px-5 py-4 font-medium">Status</th>
              <th className="px-5 py-4 font-medium">Retries</th>              <th className="px-5 py-4 font-medium">Last Error</th>              <th className="px-5 py-4 font-medium">Scheduled At</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-200 bg-white">
            {rows.map((row) => (
              <tr key={`${row.recipient}-${row.message}`} className="hover:bg-slate-50/70">
                <td className="px-5 py-4 font-medium text-slate-900">{row.recipient}</td>
                <td className="px-5 py-4 text-slate-600">{row.message}</td>
                <td className="px-5 py-4">
                  <span className={`inline-flex rounded-full px-3 py-1 text-xs font-semibold ${row.statusTone ?? "bg-slate-100 text-slate-600 ring-1 ring-slate-200"}`}>
                    {row.status}
                  </span>
                </td>
                <td className="px-5 py-4 text-slate-600">{row.retries}</td>
                <td className="px-5 py-4 text-rose-600 max-w-[240px] truncate">{row.lastError ?? "-"}</td>
                <td className="px-5 py-4 text-slate-600">{row.scheduledAt}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
}
