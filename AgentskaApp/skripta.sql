INSERT INTO public.user_table (
id, activated, email, password, username) VALUES (
'500', true, 'a@gmail.com', '$2a$04$Vbug2lwwJGrvUXTj6z7ff.97IzVBkrJ1XfApfGNl.Z695zqcnPYra', 'anija');

INSERT INTO public.user_table (
id, activated, email, password, username) VALUES (
'501', true, 'alexmirkovic1719@gmail.com', '$2a$04$Vbug2lwwJGrvUXTj6z7ff.97IzVBkrJ1XfApfGNl.Z695zqcnPYra', 'alexmirkovic1719');

INSERT INTO public.user_table (
id, activated, email, password, username) VALUES (
'502', true, 'admin@gmail.com', '$2a$04$Vbug2lwwJGrvUXTj6z7ff.97IzVBkrJ1XfApfGNl.Z695zqcnPYra', 'admin');
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
501,3);
INSERT INTO public.user_roles(
user_id, role_id) VALUES (
502,1);


INSERT INTO public.company (
id, contact_info, description, name, validated, owner_id) VALUES (
1, 'neki tamo kontakt', 'losa kompanija1', 'IGOR', true, 501);
INSERT INTO public.company (
id, contact_info, description, name, validated, owner_id) VALUES (
2, 'caos', 'losa kompanija', 'IGOR2', true, 501);






