# Importador API 🗂️

## 📌 Sobre  
A **Importador API** permite o upload e processamento de arquivos, além de autenticação de usuários via **JWT**.  

## 🚀 Tecnologias  
- **Golang** (Gin, GORM)  
- **PostgreSQL/MySQL**  
- **JWT** para autenticação  
- **Swagger** para documentação  

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
PORT=8080
DB_HOST=localhost
DB_USER=seu_usuario
DB_PASS=sua_senha
DB_NAME=importador
JWT_SECRET=seu_segredo_super_secreto
```

4️⃣ Gere as tabelas do banco, com o script na pasta sql

5️⃣  Execute o servidor:
```sh
go run main.go
```
    a api esta disponivel em:
```sh
    http://localhost:8080
```

📌 Endpoints








