export function Navbar() {
  return (
    <header className="sticky top-0 z-10 border-b border-white/60 bg-white/75 backdrop-blur">
      <div className="flex items-center gap-4 px-6 py-4 lg:px-8">
        <div className="flex-1">
          <label className="relative block max-w-xl">
            <span className="sr-only">Search</span>
            <input
              type="search"
              placeholder="Search notifications, emails, statuses"
              className="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-800 outline-none transition focus:border-indigo-400 focus:bg-white"
            />
          </label>
        </div>

        <div className="hidden items-center gap-3 md:flex">
          <button className="rounded-2xl border border-slate-200 bg-white px-4 py-3 text-sm font-medium text-slate-700 shadow-sm transition hover:border-slate-300 hover:bg-slate-50">
            Profile
          </button>
          <div className="flex h-11 w-11 items-center justify-center rounded-2xl bg-slate-950 text-sm font-semibold text-white shadow-lg shadow-slate-950/15">
            A
          </div>
        </div>
      </div>
    </header>
  );
}
