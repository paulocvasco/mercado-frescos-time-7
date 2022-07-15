# Mercado Fresco Team 7

Projeto  Mercado Frescos tem como objetivo implementar uma API REST, aplicando os conhecimentos adquiridos durante o BOOTCAMP-GO MELI.

### Swagger: https://furydocs.io/bootcamp-go-w1-s4-7-4/0.0.4/specs/

## Habilidades Desenvolvidas:
- Criar um CRUD com GO.
- Utilizar o gin.
- Realizar Unit Testes.

## Tecnologias Utilizadas:
- [Go](https://go.dev/doc/) 
- [Gin-Goninc](https://github.com/gin-gonic/gin)
- [Testify](https://github.com/stretchr/testify)
- [Swaggo](https://github.com/swaggo/swag)
## End Points Desenvolvidos:

1. endpoint: `sellers`<br/>
/api/v1/sellers `[GET]`: List all Seller.<br/>
/api/v1/sellers/:id `[GET]`: List a Seller.<br/> 
/api/v1/sellers `[POST]`: Create a Seller.<br/>
/api/v1/sellers/:id `[PATCH]`: Modify Seller.<br/>
/api/v1/sellers/:id `[DELETE]`: Delete Seller.<br/>

2. endpoint: `warehouses`<br/>
/api/v1/warehouses `[GET]`: List all Warehouse.<br/>
/api/v1/warehouses/:id `[GET]`: List a Warehouse.<br/>
/api/v1/warehouses `[POST]`: Create a Warehouse.<br/>
/api/v1/warehouses/:id`[PATCH]`: Modify Warehouse.<br/>
/api/v1/warehouses/:id`[DELETE]`: Delete Warehouse.<br/> 

3. endpoint: `sections`<br/>
/api/v1/sections `[GET]`: List all Section.<br/>
/api/v1/sections/:id `[GET]`: List a Section.<br/> 
/api/v1/sections `[POST]`: Create a Section.<br/>
/api/v1/sections/:id `[PATCH]`: Modify Section.<br/>
/api/v1/sections/:id `[DELETE]`: Delete Section.<br/> 

4. endpoint: `products`<br/>
/api/v1/products `[GET]`: List all Product.<br/>
/api/v1/products/:id`[GET]`: List a Product.<br/>
/api/v1/products `[POST]`: Create a Product.<br/>
/api/v1/products/:id `[PATCH]`: Modify Product.<br/>
/api/v1/products/:id `[DELETE]`: Delete Product.<br/>

5. endpoint: `employees`<br/>
/api/v1/employees `[GET]`: List all Employee.<br/>
/api/v1/employees/:id `[GET]`: List an Employee.<br/>
/api/v1/employees `[POST]`: Create an Employee.<br/>
/api/v1/employees/:id `[PATCH]`: Modify Employee.<br/>
/api/v1/employees/:id `[DELETE]`: Delete Employee.<br/>

6. endpoint: `buyers`<br/>
/api/v1/buyers `[GET]`: List all Buyers.<br/>
/api/v1/buyers/:id `[GET]`: List a Buyer.<br/>
/api/v1/buyers `[POST]`: Create a Buyer.<br/>
/api/v1/buyers/:id `[PATCH]`: Modify Buyers.<br/>
/api/v1/buyers/:id `[DELETE]`: Delete Buyer.<br/>

## Iniciando o Projeto:

```sh  
# Clone o repositorio
git clone git@github.com:paulocvasco/mercado-frescos-time-7.git

# Entre na pasta
cd mercado-frescos-time-7/go-web/cmd/server

# Instalar a aplicação
go install

# Se GOPATH estiver adicionado em PATH, iniciar o servidor:
server

```

## Rodando os Testes:

```sh  
# Na raiz do projeto executar o comando:
go test -v ./...

# Vizualizar a cobertura dos testes, na pasta raiz do projeto executar:
go test ./...  -coverprofile=coverage.out

# Vizulizar o report de coverage:
go tool cover -html=coverage.out

```

## Gerando a documentação:

```sh 

# Certifique que o Vendor esta OK 
go mod vendor

# Gerando a documentação
swagger generate spec -o ./swagger.yaml --scan-models\n

# Gerando gerando o server visualização e testes:
swagger serve -F=swagger swagger.yaml

```

Para faciliar os testes foi criado um script de criação da tabelas do banco a adição de alguns objetos
```sh 

# Na pasta raiz do projeto: 
./script.sh -c|-p --user [USER] --password [PASSWORD]

```

As flags -c e -p representam a opção criar o banco e popular as tabelas respectivamente.

Caso o user no mysql não tenha senha, passar a senha `nopass` no argumento.



