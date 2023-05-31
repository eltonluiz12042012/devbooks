insert into usuarios(nome, nick, email, senha)
values
("Usuario 1", "User1","usuario1@user.com.br","$2a$10$7/lVrLHPNFYe36Wj4/7SmOvSLjdLuGAHULcV/NSKyQ4gyRR8Rv.I."),
("Usuario 2", "User2","usuario2@user.com.br","$2a$10$7/lVrLHPNFYe36Wj4/7SmOvSLjdLuGAHULcV/NSKyQ4gyRR8Rv.I."),
("Usuario 3", "User3","usuario3@user.com.br","$2a$10$7/lVrLHPNFYe36Wj4/7SmOvSLjdLuGAHULcV/NSKyQ4gyRR8Rv.I."),
("Usuario 4", "User4","usuario4@user.com.br","$2a$10$7/lVrLHPNFYe36Wj4/7SmOvSLjdLuGAHULcV/NSKyQ4gyRR8Rv.I."),
("Usuario 5", "User5","usuario5@user.com.br","$2a$10$7/lVrLHPNFYe36Wj4/7SmOvSLjdLuGAHULcV/NSKyQ4gyRR8Rv.I.");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(1,3),
(1,4),
(2,3),
(4,1),
(2,1),
(2,5),
(3,2),
(3,5)
(4,3),
(5,1),
(5,2);
