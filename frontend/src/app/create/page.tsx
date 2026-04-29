import { DashboardShell } from "@/app/dashboard-shell";
import { NotificationComposer } from "@/components/NotificationComposer";

export default function CreateNotificationPage() {
  return (
    <DashboardShell>
      <div className="flex-1 px-6 py-8 lg:px-8">
        <NotificationComposer />
      </div>
    </DashboardShell>
  );
}
