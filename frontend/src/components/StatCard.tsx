type StatCardProps = {
  label: string;
  value: string;
  trend: string;
};

export function StatCard({ label, value, trend }: StatCardProps) {
  return (
    <article className="rounded-3xl border border-white/70 bg-white p-6 shadow-[0_20px_60px_-32px_rgba(15,23,42,0.25)]">
      <p className="text-sm font-medium text-slate-500">{label}</p>
      <div className="mt-4 flex items-end justify-between gap-4">
        <h3 className="text-3xl font-semibold tracking-tight text-slate-950">{value}</h3>
        <span className="rounded-full bg-emerald-50 px-3 py-1 text-xs font-semibold text-emerald-700">
          {trend}
        </span>
      </div>
    </article>
  );
}
