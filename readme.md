curl -i -H "X-Trace-Id: trace-123" "http://127.0.0.1:8080/v1/compare/products?ids=1,2,3"









todos:
IA para gerar base de dados - ok
IA para models de produtos - ok
carregar dados em memoria - ok
camada de repository para consulta ao banco - ok
camada de negocio para filtrar itens - ok
definir forma de comparar valor via json - ok
ter um timeout interno para busca dos dados (3s) - ok
tratamento de exceção e padronização de mensagem de erro
ids duplicados
lista vazia
todos itens nao encontrados
algum item nao encontrado
nao ter alguma info deve eliminar a comparação?
retornar comparacao se tiver ao menos 2 itens a serem comparados
status codes
logs e telemetria - ok
docker-compose
testes unitarios - ok