INSERT INTO public.user_table (
id, activated, email, password, username) VALUES (
'500', true, 'a@gmail.com', '$2a$04$Vbug2lwwJGrvUXTj6z7ff.97IzVBkrJ1XfApfGNl.Z695zqcnPYra', 'neki tamo');

INSERT INTO public.user_table (
id, activated, email, password, username) VALUES (
'501', true, 'alexmirkovic1719@gmail.com', '$2a$04$Vbug2lwwJGrvUXTj6z7ff.97IzVBkrJ1XfApfGNl.Z695zqcnPYra', 'alexmirkovic1719');

INSERT INTO public.role(
id, name) VALUES (
1,'ADMINISTRATOR');

INSERT INTO public.role(
id, name) VALUES (
2,'USER');

INSERT INTO public.role(
id, name) VALUES (
3,'OWNER');


INSERT INTO public.user_roles(
user_id, role_id) VALUES (
500,2);
INSERT INTO public.user_roles(
user_id, role_id) VALUES (
501,2);





