# Exemplos de Uso da API

## 🔍 Endpoint: GET /vulneravel/get

Busca um elemento pelo ID no banco de dados.

### Como usar:

#### 1. Via Navegador (mais simples)
Abra o navegador e acesse:
```
http://localhost:8080/vulneravel/get?id=1
```

#### 2. Via cURL (linha de comando)
```bash
curl http://localhost:8080/vulneravel/get?id=1
```

#### 3. Via Postman/Insomnia
- Método: `GET`
- URL: `http://localhost:8080/vulneravel/get`
- Query Parameters:
  - `id`: `1` (ou qualquer ID que exista no banco)

#### 4. Via JavaScript (fetch)
```javascript
fetch('http://localhost:8080/vulneravel/get?id=1')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Erro:', error));
```

### Resposta de Sucesso (200 OK):
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@example.com",
  "address": "Rua A, 123",
  "phone": "11999999999",
  "credit_card": "1234-5678-9012-3456"
}
```

### Resposta de Erro (400 Bad Request):
```json
{
  "error": "id é obrigatório"
}
```

### Resposta de Erro (500 Internal Server Error):
```json
{
  "error": "elemento não encontrado com id: 1"
}
```

## 🚀 Passos para testar:

1. **Inicie a aplicação:**
   ```bash
   go run app/main.go
   ```

2. **Popule o banco com dados de exemplo (se ainda não fez):**
   ```bash
   go run scripts/seed_database.go
   ```

3. **Teste a API:**
   - Abra o navegador em: `http://localhost:8080/vulneravel/get?id=1`
   - Ou use cURL: `curl http://localhost:8080/vulneravel/get?id=1`

## ⚠️ Nota sobre Segurança

Esta API está **intencionalmente vulnerável** a SQL Injection. Por exemplo:
- `http://localhost:8080/vulneravel/get?id=1 OR 1=1` pode retornar todos os registros
- `http://localhost:8080/vulneravel/get?id=1; DROP TABLE elements;--` pode deletar a tabela

Isso é para fins de demonstração de vulnerabilidades. Em produção, sempre use prepared statements!
