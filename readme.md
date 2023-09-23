## Bexs Dev Test

Este repositório contém a implementação do teste técnico presente em [README_BACKEND_SR.md.pdf](./README_BACKEND_SR.md.pdf).

### Dependências

[Docker](https://www.docker.com/) com suporte a [compose](https://docs.docker.com/compose/) ou [Podman](https://podman.io/) com [compose](https://github.com/containers/podman-compose).

### Execução

1. Faça o clone deste repositório.
2. Execute o comando da dependência escolhida acima, sendo `docker compose up -d` ou `podman-compose up -d`.
3. Acesse `http://127.0.0.1:9007/api/status` e veja o retorno do status.

### Rotas

Abaixo documentação de todas as rotas disponíveis no projeto.

#### Status

<details>
 <summary><code>GET</code> <code><b>/api/status</b></code> <code> - </code> <code>Status da aplicação</code></summary>

##### Respostas

> | Status code | Resposta |
> |-----------|-----------|
> | `200` | Status com a hora da resposta |

##### Exemplo

> ```bash
> curl --request GET --url 'http://127.0.0.1:9007/api/status'
> ```

</details>

#### Parceiros

<details>
 <summary><code>POST</code> <code><b>/api/partners</b></code> <code> - </code> <code>Criação de um novo parceiro</code></summary>

##### Parâmetros

> | nome | Tipo | Descrição |
> |-----------|-----------|-----------|
> | id | string | Identificador do parceiro |
> | trading_name | string | Nome do parceiro |
> | document | string | Documento do parceiro |
> | currency | string | Unidade monetária do parceiro, sendo opções "USD", "EUR" ou "GBP" |

##### Respostas

> | Status code | Resposta |
> |-----------|-----------|
> | `201` | Parceiro criado com sucesso |
> | `400` | Um dos dados informados acima estão incorretos |
> | `422` | O parceiro já existe com o ID ou documento informado |
> | `500` | Erro interno do servidor |

##### Exemplo

> ```bash
> curl --request POST \
>   --url http://127.0.0.1:9007/api/partners \
>   --header 'Content-Type: application/json' \
>   --data '{
>   "id": "12",
>   "trading_name": "Pato",
>   "document": "12345",
>   "currency": "USD"
> }'
> ```

</details>

#### Pagamentos

<details>
 <summary><code>GET</code> <code><b>/api/payments</b></code> <code> - </code> <code>Listagem de pagamentos</code></summary>

##### Parâmetros

> | nome | Tipo | Descrição |
> |-----------|-----------|-----------|
> | offset | integer | Offset para paginação, sendo opcional. Default: 0 |
> | limit | integer | Limit para paginação, sendo opcional. Default: 10 |

##### Respostas

> | Status code | Resposta |
> |-----------|-----------|
> | `200` | Lista de pagamentos |
> | `400` | Um dos dados informados acima estão incorretos |
> | `500` | Erro interno do servidor |

##### Exemplo

> ```bash
> curl --request GET --url 'http://127.0.0.1:9007/api/payments?offset=0&limit=10'
> ```

</details>

<details>
 <summary><code>POST</code> <code><b>/api/payments</b></code> <code> - </code> <code>Criação de um novo pagamento</code></summary>

##### Parâmetros

> | nome | Tipo | Descrição |
> |-----------|-----------|-----------|
> | partner_id | string | Identificador do parceiro |
> | amount | string | Valor do pagamento, contendo duas casas decimais separadas por ponto "." e sem separador de milhar |
> | consumer.name | string | Nome do consumidor |
> | consumer.national_id | string | Documento do consumidor |

##### Respostas

> | Status code | Resposta |
> |-----------|-----------|
> | `201` | Pagamento criado com sucesso |
> | `400` | Um dos dados informados acima estão incorretos |
> | `422` | O parceiro informado não existe |
> | `422` | Duplicidade de criação de pagamento |
> | `500` | Erro interno do servidor |

##### Exemplo

> ```bash
> curl --request POST \
>   --url http://127.0.0.1:9007/api/payments \
>   --header 'Content-Type: application/json' \
>   --data '{
>   "partner_id": "10",
>   "amount": "99.05",
>   "consumer": {
>     "name": "Oliver Tsubasa",
>     "national_id": "30243434597"
>   }
> }'
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/api/payments/:id</b></code> <code> - </code> <code>Visualização de um pagamento</code></summary>

##### Parâmetros

> | nome | Tipo | Descrição |
> |-----------|-----------|-----------|
> | id | string | Identificador do pagamento |

##### Respostas

> | Status code | Resposta |
> |-----------|-----------|
> | `200` | Pagamento visualizado |
> | `404` | O pagamento não existe pelo identificador informado |
> | `500` | Erro interno do servidor |

##### Exemplo

> ```bash
> curl --request GET --url 'http://127.0.0.1:9007/api/payments/01HB205RPJ3DS4C6ZFPYTX6P4Z'
> ```

</details>
