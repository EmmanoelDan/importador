# Importador API 🗂️

## 📌 Sobre  
A **Importador API** permite o upload e processamento de arquivos, além de autenticação de usuários via **JWT**. Seu objetivo é facilitar a importação e gerenciamento de dados com performance, eficiência e segurança.

## 🚀 Tecnologias  
- **Golang** (Gin, GORM)  
- **PostgreSQL/MySQL**  
- **JWT** para autenticação  
- **GODOTENV** para variaveis de ambiente

## 📜 Instalação  

1️⃣ Clone o repositório:  
```sh
git clone https://github.com/seu-usuario/importador.git
cd importador
```

2️⃣ Instale as dependências:
```sh
go mod tidy
```

3️⃣ Configure o arquivo .env com as variáveis de ambiente:
```sh
DB_PORT=8080
DB_HOST=localhost
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=importador
JWT_SECRET=seu_segredo_super_secreto
```

4️⃣ Gere as tabelas do banco, com o script na pasta **./sql** arquivo **normalizacao.sql**

5️⃣  Execute o servidor:
```sh
go run main.go
```

```sh
    http://localhost:8080
```

## 📌 Endpoints

OBS: POST Register -> POST Sign -> POST Import Files -> GET Bellings

- **URL:** `/register`
- **Método:** `POST`

### Body:

<pre>
{
  "username": "admin",
  "password": "123456"
}
</pre>

### Resposta (201 created):

<pre>
{
    "message": "Create user successfully",
    "user": "admin"
}
</pre>

- **URL:** `/login`
- **Método:** `POST`

### Body:

<pre>

{
  "username": "admin",
  "password": "123456"
}
</pre>

### Resposta (200 OK):

<pre>
{
  token": "jwt_token_aqui"
}
</pre>


- **URL:** `/import_file`
- **Método:** `POST`
- **Tipo de Conteúdo:** `multipart/form-data`
- **Autenticação:** Bearer Token (JWT)
<pre>
{
    "message": "Upload successfully imported"
}
</pre>

- **URL:** `/billings?page=1&pageSize=10`
- **Método:** `GET`
- **Tipo de Conteúdo:** `multipart/form-data`
- **Autenticação:** Bearer Token (JWT)
- **Params** pagination: page e pageSize. 



📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.






