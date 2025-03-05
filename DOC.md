

# Introdução

Ao decidir a arquitetura para este projeto Go, optei por seguir os princípios da Clean Architecture, mas seguindo com um projeto mais simples e objetivo. 

# Estrutura do Projeto

O projeto está organizado nas seguintes pastas:

- config/: Configurações gerais do projeto, como variáveis de ambiente e configuração de banco de dados.

- controller/: Contém os handlers HTTP responsáveis por lidar com as requisições e respostas da API.

- middleware/: Implementação de middlewares para autenticação e segurança.

- model/: Definição das estruturas de dados (structs) que representam as entidades do sistema.

- repository/: Responsável pela camada de acesso a dados, abstraindo a interação com bancos de dados.

- router/: Configuração das rotas e endpoints da API.

- service/: Contém a lógica de negócio e processamento das regras do sistema.

- sql/: Scripts e migrações para o banco de dados.

- temp/: Pasta temporária utilizada para armazenamento de arquivos durante a execução do sistema.

- util/: Funções utilitárias usadas em diversas partes do projeto.

# Bibliotecas Utilizadas

- Gin -> Framework web para construir APIs REST em Go.

- GORM -> ORM para facilitar a interação com bancos de dados.

- jwt-go -> Biblioteca para geração e validação de tokens JWT.

- godotenv -> Leitura de variáveis de ambiente a partir de arquivos .env.

# Arquivos principais:

- .env: Arquivo para configuração de variáveis de ambiente.

- .gitignore: Lista de arquivos e pastas ignorados pelo Git.

- DOC.md: Documentação do projeto.

- go.mod / go.sum: Gerenciamento de dependências do Go.

- LICENSE: Licença do projeto.

- main.go: Ponto de entrada da aplicação.

- README.md: Instruções iniciais sobre o projeto.

# Camadas da Clean Architecture

1. Controller (Interface Adapters)

- Responsável por lidar com as requisições HTTP, chamando os serviços adequados e formatando as respostas. Utilizei framework Gin.

2. Service (Use Cases)

- Implementa a lógica de negócio do sistema, separando regras de negócio da infraestrutura. Esta camada interage com os repositórios para buscar ou manipular dados.

3. Repository (Data Layer)

- Faz a interação com o banco de dados ou outras fontes de dados externas. Utilizei GORM para facilitar a manipulação de dados.

4. Model (Entities)

- Define as estruturas de dados utilizadas pelo sistema. Essas structs representam os objetos de negócio do projeto.

# Vantagens da Clean Architecture

- Desacoplamento

- A separação em camadas permite que os componentes do sistema sejam substituídos ou alterados sem impactar o restante do código.

# Testabilidade

- A modularização torna o projeto altamente testável, permitindo a criação de testes unitários para cada camada separadamente.

- Manutenção e Escalabilidade

- A estrutura clara facilita a manutenção e a evolução do sistema, garantindo que o projeto possa crescer sem comprometer a organização do código.

# Conclusão

A adoção da Clean Architecture neste projeto visa garantir um código limpo, modular e de fácil manutenção. Com essa estrutura, podemos escalar e evoluir a aplicação sem dificuldades, garantindo uma base sólida para o desenvolvimento.

