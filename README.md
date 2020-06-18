# Desafio Bravo - HURB

## Sobre o Projeto
O projeto é um conversor de moedas, para isso foi criado uma API capaz de realizar a conversão, criação, atualização, exclusão e listagem de moedas. Além disso foi criado um FrontEnd que permite de forma intuitiva a utilização das funções básica da API.

O sistema possuí um serviço de atualização de 5 em 5 minutos (~podendo ser alterado no index.js da API~) das moedas cadastradas. Esse serviço é possível graças a uma API externa e aberta que fornece informações atualizadas de diversas moedas. A API externa pode ser acessada [aqui](https://economia.awesomeapi.com.br).

É importante ressaltar que o valor base utilizado para conversão foi o Dólar Americano (lastro), então para inclusão de uma nova moeda, seu valor deve ser baseado no valor de 1 dólar. Exemplo: Para adicionar a moeda Euro, deve informar o valor de x Euros equivalente a 1 dólar.

A arquitetura utilizada foi um monolito que utiliza as seguintes tecnologias: MongoDB - NodeJs (JavaScript) - VueJs (Framework JavaScript).

O projeto não foi colocado em um contâiner Docker devido as dificuldades encontradas devido a incompatibilidade entre o Docker e a versão do meu Windows (Windows Home).

Foi utilizado utilizado o framework VueJs no frontend e para inicialização do projeto foi utilizado seu [CLI](https://cli.vuejs.org/).

## Requisitos necessários

É necessário a instalação do NodeJs 12.X, MongoDB e GIT para inicialização do projeto.

-   **Instalação do NodeJs:** [Acesse aqui](https://nodejs.org/en/download/) e realize o download do NodeJs 12.X
-   **Instalação do MongoDB:** Realize o download do instalador [aqui](https://www.mongodb.com/try/download/community) ou caso já tenha instalado o NodeJs use o comando `npm install mongodb --save`.
-   **Instalação do GIT:** Realize o download [aqui](https://git-scm.com/downloads) e realize a instalção.

Realizado as configurações vamos para a próxima etapa.

## Incializando o projeto

Acesse o [repositório](https://github.com/lmaiaa/challenge-bravo) do meu desafio

Execute o comando

```
git clone https://github.com/lmaiaa/challenge-bravo.git
```

Em seguida acesse a pasta clonada e instale as dependências

```
cd challenge-bravo

npm install
```

Após isso o projeto está pronto para ser inicializado.
Para inicializar execute o comando:

```
npm start
```

## Rotas

A API está por padrão sendo executada na URL: **http://localhost:3333**

-   GET /coins   -> Retorna uma lista de moedas cadastras na API

-   GET coin/conversion?from=*Moeda1*&to=*Moeda2*&amount=*Valor*  -> Realiza a conversão entre moedas cadastradas
-   POST /coin  -> Cria uma nova moeda
-   PUT /coin/*:name*/*:value*  -> Edita uma moeda (lastro dólar americano)
-   DELETE /coin/*:name*  -> Exclui uma moeda
