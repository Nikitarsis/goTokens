BEGIN;
DELETE FROM public.keys WHERE kid = $1;
DELETE FROM public.user_agents WHERE kid = $1;
DELETE FROM public.ips WHERE kid = $1;
COMMIT;