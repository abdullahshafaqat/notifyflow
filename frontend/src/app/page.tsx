import { DashboardClient } from "@/components/DashboardClient";
import { getNotifications } from "@/lib/api";

export default async function Home() {
  const notifications = await getNotifications().catch(() => []);

  return <DashboardClient initialNotifications={notifications} />;
}
