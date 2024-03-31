## ПВЗ API

#### Запуск
```bash
docker-compose up --build
```

### Добавление ПВЗ


```bash
curl -k -X POST \
  https://127.0.0.1:9000/pick-up-point \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
  -d '{
    "name": "PickUpPoint1",
    "address": {
        "region": "Курская область",
        "city": "Курск",
        "street": "Студенческая",
        "house_num": "13А"
    },
    "phone_number": "88005553535"
}'
```
### Обновление информации о ПВЗ

```bash
curl -k -X PUT \
  https://127.0.0.1:9000/pick-up-point \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
  -d '{
    "id":1,
    "name": "PickUpPoint1",
    "address": {
        "region": "Курская область",
        "city": "Курск",
        "street": "Ленина",
        "house_num": "13А"
    },
    "phone_number": "89001001010"
}'

```

### Получение списка всех ПВЗ

```bash
curl -k -X GET   https://127.0.0.1:9000/pick-up-points   -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получение ПВЗ по его ID

```bash
curl -k -X GET \
  https://127.0.0.1:9000/pick-up-point/1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Удаление ПВЗ

```bash
curl -k -X DELETE \
  https://127.0.0.1:9000/pick-up-point/1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Принять заказ от курьера (добавить на ПВЗ)
```
curl -k -X POST \
https://127.0.0.1:9000/order \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "id": 12345,
    "client_id": 67890,
    "weight": 12.1,
    "price": 50.99,
    "storage_expiration_date": "2024-03-29T12:00:00Z",
    "package_type": "box"
}'
```

### Отдать заказ курьеру (Удалить с ПВЗ)
```
curl -k -X DELETE \
https://127.0.0.1:9000/order/12345 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получить список возвратов
```
curl -k -X GET \
https://127.0.0.1:9000/orders/returns/2 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получить список заказов клиента
```
curl -k -X GET \
https://127.0.0.1:9000//clients/{CLIENT_ID}/orders?num_of_last_orders=1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Отдать заказ клиенту
```
curl -k -X PUT \
https://127.0.0.1:9000/orders/issue \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "orders-ids": [12346, 12347]
}'
```

### Принять возврат от клиента
```
curl -k -X PUT \
https://127.0.0.1:9000/orders/return \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "order_id": 12345,
    "client_id": 67890
}'
```
