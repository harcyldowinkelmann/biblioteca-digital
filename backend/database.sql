CREATE TABLE usuario_tipo (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255) NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO usuario_tipo (nome) VALUES ('Comum');
INSERT INTO usuario_tipo (nome) VALUES ('Especial');

CREATE TABLE usuario (
  id SERIAL PRIMARY KEY, 
  nome VARCHAR(255) NULL,
  cpf VARCHAR(45) NULL,
  email VARCHAR(45) NULL,
  senha VARCHAR(255) NULL,
  usuario_tipo_id INTEGER REFERENCES usuario_tipo (id) 
    ON DELETE SET NULL 
    ON UPDATE CASCADE, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO usuario (id, nome, cpf, email, senha, usuario_tipo_id) VALUES (1, 'Harcyldo Winkelmann', '000.000.000-10', 'harcyldo.winkelmann060@academico.ufgd.edu.br', '$2a$10$bNRfGNjv31x9qLLif.9EOe7Vd5XFSAOlhJMi3a8f6nHhm0OY6ABF.', 1);