import Link from "next/link";

const navigation = [
  { label: "Dashboard", href: "/", active: true },
  { label: "Notifications", href: "/notifications" },
  { label: "Scheduled", href: "/notifications?filter=scheduled" },
  { label: "Failures", href: "/notifications?filter=failed" },
  { label: "Create Notification", href: "/create" },
];

export function Sidebar() {
  return (
    <aside className="hidden w-72 shrink-0 border-r border-slate-200/80 bg-white/85 px-5 py-6 shadow-[0_0_40px_-30px_rgba(15,23,42,0.25)] backdrop-blur lg:flex lg:flex-col">
      <div className="mb-8">
        <div className="inline-flex h-11 w-11 items-center justify-center rounded-2xl bg-slate-950 text-sm font-semibold text-white shadow-lg shadow-slate-950/20">
          NF
        </div>
        <div className="mt-4">
          <p className="text-sm font-medium uppercase tracking-[0.2em] text-indigo-600">
            Project
          </p>
          <h2 className="mt-1 text-2xl font-semibold tracking-tight text-slate-950">
            NotifyFlow
          </h2>
          <p className="mt-2 text-sm leading-6 text-slate-500">
            Clean delivery ops for scheduled and instant notifications.
          </p>
        </div>
      </div>

      <nav className="flex flex-1 flex-col gap-2">
        {navigation.map((item) => (
          <Link
            key={item.label}
            href={item.href}
            className={`rounded-2xl px-4 py-3 text-sm font-medium transition-colors ${
              item.active
                ? "bg-slate-950 text-white shadow-lg shadow-slate-950/15"
                : "text-slate-600 hover:bg-slate-100 hover:text-slate-950"
            }`}
          >
            {item.label}
          </Link>
        ))}
      </nav>

      <div className="mt-8 rounded-3xl border border-slate-200 bg-slate-50 p-4">
        <p className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-400">
          Delivery Status
        </p>
        <div className="mt-3 flex items-center gap-3">
          <span className="h-2.5 w-2.5 rounded-full bg-emerald-500" />
          <p className="text-sm font-medium text-slate-800">Worker online</p>
        </div>
        <p className="mt-2 text-sm leading-6 text-slate-500">
          Scheduler and delivery pipeline are connected.
        </p>
      </div>
    </aside>
  );
}
