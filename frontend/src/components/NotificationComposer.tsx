'use client';

import { useState } from "react";

import { createNotification } from "@/lib/api";

type NotificationComposerProps = {
  onSuccess?: (message: string) => void;
  onError?: (message: string) => void;
};

export function NotificationComposer({ onSuccess, onError }: NotificationComposerProps) {
  const [form, setForm] = useState({ to: "", message: "", sendAt: "" });
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);

    try {
      const payload: { to: string; message: string; send_at?: string } = {
        to: form.to,
        message: form.message,
      };

      if (form.sendAt) {
        payload.send_at = new Date(form.sendAt).toISOString();
      }

      const response = await createNotification(payload);
      onSuccess?.(`Notification accepted${response.id ? ` • ID: ${response.id}` : ""}`);
      setForm({ to: "", message: "", sendAt: "" });
    } catch (submitError) {
      onError?.(submitError instanceof Error ? submitError.message : "Failed to create notification");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <section className="rounded-3xl border border-white/70 bg-white p-6 shadow-[0_20px_60px_-32px_rgba(15,23,42,0.25)]">
      <div className="flex flex-col gap-2 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p className="text-sm font-medium uppercase tracking-[0.2em] text-indigo-600">
            Quick create
          </p>
          <h2 className="mt-1 text-2xl font-semibold tracking-tight text-slate-950">
            Create Notification
          </h2>
        </div>
        <p className="text-sm text-slate-500">Connects to POST /send</p>
      </div>

      <form onSubmit={handleSubmit} className="mt-6 grid gap-4 lg:grid-cols-2">
        <label className="grid gap-2 text-sm font-medium text-slate-700">
          Email
          <input
            type="email"
            required
            value={form.to}
            onChange={(event) => setForm({ ...form, to: event.target.value })}
            className="rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none transition focus:border-indigo-400 focus:bg-white"
            placeholder="user@gmail.com"
          />
        </label>

        <label className="grid gap-2 text-sm font-medium text-slate-700 lg:col-span-2">
          Message
          <textarea
            required
            value={form.message}
            onChange={(event) => setForm({ ...form, message: event.target.value })}
            className="min-h-32 rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none transition focus:border-indigo-400 focus:bg-white"
            placeholder="Write your notification message"
          />
        </label>

        <label className="grid gap-2 text-sm font-medium text-slate-700">
          Schedule (optional)
          <input
            type="datetime-local"
            value={form.sendAt}
            onChange={(event) => setForm({ ...form, sendAt: event.target.value })}
            className="rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none transition focus:border-indigo-400 focus:bg-white"
          />
        </label>

        <div className="flex items-end justify-end lg:col-span-2">
          <button
            type="submit"
            disabled={isSubmitting}
            className="inline-flex items-center justify-center rounded-2xl bg-slate-950 px-5 py-3 text-sm font-semibold text-white shadow-lg shadow-slate-950/15 transition hover:bg-slate-800 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {isSubmitting ? "Sending…" : "Send notification"}
          </button>
        </div>
      </form>
    </section>
  );
}
