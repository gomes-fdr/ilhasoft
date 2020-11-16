# Projeto de teste IlhaSoft

Esta API foi desenvolvida como prova de conceito para empresa IlhaSoft.

**Atenção:** Está implementado em um repo PRIVADO do github, para evitar plágio.

## Sobre a API

Esta é uma API que segue o padrão Rest, portanto utilizaremos os *verbos* do protocolo http para realizar as operações de *CRUD* de uma *todo list*.

## Stack utilizada

* Golang como linguagem de desenvolvimento da API
* Dados gravados em memória(usa uma strutura do tipo `map`)
    

## End points

* POST /todos: Grava uma tarefa
* GET /todos: Lista todas as tarefas contidas
* GET /todos/{id} Lista uma tarefa específica
* DELETE /todo/{id}: Apaga tarefa por id

Todas as entradas e saídas de dados são no formato JSON.

## Tipo de dado

Uma tarefa contém as seguintes informações:

    {
        "id": "xxx",
        "description": "The task",
        "status": "done"
    }

## Testes

Para obter informações da API

    curl localhost:3030 -v
    *   Trying ::1:3030...
    * Connected to localhost (::1) port 3030 (#0)
    > GET / HTTP/1.1
    > Host: localhost:3030
    > User-Agent: curl/7.71.1
    > Accept: */*
    > 
    * Mark bundle as not supporting multiuse
    < HTTP/1.1 200 OK
    < Content-Type: application/json
    < Date: Sat, 14 Nov 2020 22:28:12 GMT
    < Content-Length: 56
    < 
    * Connection #0 to host localhost left intact
    {"version":"1.0.0","description":"Uma API de todo list"}


Para inserir uma tarefa:

    curl -X POST -H "Content-Type: application/json"  -d '{"Description":"Uma tarefa", "Status":"todo"}'  localhost:3030/todos -v

Para listar todas as tarefas:

    curl localhost:3030/todos -v

Listar uma tarefa

    curl -X GET localhost:3030/todo/{id} -v

Para apagar uma tarefa:

    curl -X DELETE localhost:3030/todo/{id} -v

## O que faltou

* Testes unitários
* Salvar numa base dados
* Um end-point para atualizar uma tarefa [PATCH {campo}/{valor}]
* Um end-point para buscar tarefas por status [GET {campo}/{valor}]
