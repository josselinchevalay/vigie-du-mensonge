import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/admin/users/$userTag')({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/admin/users/$userTag"!</div>;
}
