insert into usuarios (nome, nick, email, senha)
values
("Alexandre", "Bueno", "alexandre@gmail.com","$2a$10$s6VtloXxPS9ZJE08P4P0R.VWW0/VR0h0a14BQbRWVvOP30po0CDve"),
("Felipe", "Felipe", "felipe@gmail.com","$2a$10$s6VtloXxPS9ZJE08P4P0R.VWW0/VR0h0a14BQbRWVvOP30po0CDve"),
("Zanza", "Zanza", "zanza@gmail.com","$2a$10$s6VtloXxPS9ZJE08P4P0R.VWW0/VR0h0a14BQbRWVvOP30po0CDve");

insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3),
(2, 3);