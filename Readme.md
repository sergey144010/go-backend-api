# go-backend-api

## Only see example:
```bash
sudo docker run --rm --name=go-backend-api -p 8888:80 itlab77/go-backend-api
```

http://localhost:8888/v1/questions/1

## With your data
```bash
sudo docker run -d --rm --name=go-backend-api -p 8888:80 -v "$PWD":/data itlab77/go-backend-api
```

```bash
echo '{"name":"One"}' >> customers.get.json
```

http://localhost:8888/customers

## Advanced usage

Examples map files to route:

- customers.get.json - GET /customers
- customers.post.json - POST /customers
- customers.put.json - PUT /customers
- customers.patch.json - PATCH /customers
- customers.delete.json - DELETE /customers

- v1.customers.1.get.json - GET /v1/customers/1
