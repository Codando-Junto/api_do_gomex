# api_do_gomex

Este é um projeto simples em Go que implementa uma API básica com duas rotas: uma para a página inicial (`/`) e outra para verificação de saúde (`/health`). O objetivo é demonstrar como estruturar um projeto Go com boas práticas, incluindo separação de responsabilidades em pacotes como `handlers` e `config`.

## Pré-requisitos

Certifique-se de ter o seguinte instalado em sua máquina:

- [Go](https://golang.org/dl/) (versão 1.18 ou superior)
- Git

## Como executar o projeto localmente

Siga os passos abaixo para rodar o projeto localmente:

### 1. Clone o repositório

```bash
git clone https://github.com/Codando-Junto/api_do_gomex.git
cd api_do_gomex
```

### 2. Instale as dependências
Este projeto não possui dependências externas além da biblioteca padrão do Go. Certifique-se de que o Go está configurado corretamente.

### 3. Execute o servidor
Para iniciar o servidor, execute o comando:

```bash
go run main.go
```

Se tudo estiver configurado corretamente, você verá a seguinte mensagem no terminal:

```bash
Server is running on :8080
```

### 3.1. Usando ngrok para expor a porta na internet

O ngrok é uma ferramenta que cria túneis seguros para expor localmente servidores ou aplicações à internet, permitindo acesso remoto por meio de URLs públicas.

Com o ngrok você pode expor a porta 8080 usando seu código da sua máquina e ele pode ser acessado por outra pessoa, pela internet.

**Observação importante**: Lembre-se que a internet é um local inseguro, e expor suas portas por muito tempo sem a devida segurança é um risco muito alto. Esse aplicativo deve ser usado somente para testes simples e por pouco período.

Você precisa se [cadastrar no ngrok](https://dashboard.ngrok.com/signup) e depois [instalar ele](https://dashboard.ngrok.com/signup) na sua estação de trabalho, que será onde vc vai rodar ele.

Você pode criar um domínio fixo pra você, para usar o mesmo domínio sempre [nesse link](https://dashboard.ngrok.com/domains) e nesse mesmo link você já pode pegar o link para iniciar o ngrok. Com meu domínio é assim:

```bash
ngrok http --url=relieved-intimate-labrador.ngrok-free.app 8080
```

Com o comando acima o ngrok iniciará um tunel mandando tudo que ele receber nesse domínio na porta padrão http (80) e mandará tudo para sua porta 8080 na localhost.

### 4. Teste as rotas
Você pode testar as rotas disponíveis usando um navegador, curl ou ferramentas como Postman.

#### Rota raiz /  

Descrição: Retorna uma mensagem de boas-vindas em formato JSON.
Exemplo de resposta:
```json
{
  "message": "Olá mundo!"
}
```

### Rota /health
Descrição: Retorna o status de saúde da aplicação.

Exemplo de resposta:
Código de status: 200 OK
Corpo da resposta: vazio


### Como personalizar

#### Adicionar novas rotas

Crie novos arquivos na pasta handlers e registre as rotas no arquivo main.go.

#### Alterar a porta do servidor

Modifique o valor do campo Addr no arquivo config/server.go.
Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests para melhorias no projeto.

### Licença
Este projeto está sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes. ```