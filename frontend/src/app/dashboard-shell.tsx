import { Navbar } from "@/components/Navbar";
import { Sidebar } from "@/components/Sidebar";

export function DashboardShell({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <div className="min-h-screen bg-[linear-gradient(180deg,#f8fafc_0%,#eef2ff_100%)] text-slate-900">
      <div className="mx-auto flex min-h-screen max-w-[1600px]">
        <Sidebar />
        <main className="flex min-w-0 flex-1 flex-col">
          <Navbar />
          {children}
        </main>
      </div>
    </div>
  );
}
