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
1, 'mozete nas kontaktirati preko broja telefona', 'vrsimo razne usluge', 'SCITI', true, 501);
INSERT INTO public.company (
id, contact_info, description, name, validated, owner_id) VALUES (
2, 'mozete nas kontaktirati preko broja telefona', 'vise informacija mozete dobiti na nasem sajtu', 'TRIDERM', true, 501);


INSERT INTO public.company (
id, contact_info, description, name, validated, owner_id) VALUES (
3, 'nas sajt: hhttp....', 'Odlicna kompanija za sve usluge', 'IXIY', true, 500);
INSERT INTO public.company (
id, contact_info, description, name, validated, owner_id) VALUES (
4, 'mozete nas kontaktirati preko broja telefona', 'Pomoci cemo Vam u svakom trenutku', 'IGOR4', true, 500);

INSERT INTO public.comment(
id, company_id, user_id, text) VALUES(
1, 1, 501, 'Odlicna kompanija');

INSERT INTO public.comment(
id, company_id, user_id, text) VALUES(
2, 1, 501, 'Veoma prijatna kompanija');

INSERT INTO public.comment(
id, company_id, user_id, text) VALUES(
3, 2, 501, 'Odlicna kompanija');

INSERT INTO public.comment(
id, company_id, user_id, text) VALUES(
4, 2, 501, 'Usluge nam se nisu svidele');







