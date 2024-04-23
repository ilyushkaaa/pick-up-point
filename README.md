## ПВЗ API

#### Запуск
```bash
make app_start
```

#### Юнит тесты
```bash
make unit_tests_run
```

#### Интеграционные тесты
```bash
make integration_tests_run
```

### Добавление ПВЗ


```bash
curl -k -X POST \
  https://127.0.0.1:9010/api/v1/pick-up-point \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
  -d '{
    "name": "PickUpPoint1",
    "address": {
        "region": "Курская область",
        "city": "Курск",
        "street": "Студенческая",
        "houseNum": "13А"
    },
    "phoneNumber": "88005553535"
}'
```
### Обновление информации о ПВЗ

```bash
curl -k -X PUT \
  https://127.0.0.1:9010/api/v1/pick-up-point \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
  -d '{
    "id":1,
    "name": "PickUpPoint1",
    "address": {
        "region": "Курская область",
        "city": "Курск",
        "street": "Ленина",
        "houseNum": "13А"
    },
    "phoneNumber": "89001001010"
}'

```

### Получение списка всех ПВЗ

```bash
curl -k -X GET   https://127.0.0.1:9010/api/v1/pick-up-points   -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получение ПВЗ по его ID

```bash
curl -k -X GET \
  https://127.0.0.1:9010/api/v1/pick-up-point/1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Удаление ПВЗ

```bash
curl -k -X DELETE \
  https://127.0.0.1:9010/api/v1/pick-up-point/1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Принять заказ от курьера (добавить на ПВЗ)
```bash
curl -k -X POST \
https://127.0.0.1:9010/api/v1/order \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "id": 12345,
    "clientId": 67890,
    "weight": 12.1,
    "price": 50.99,
    "storageExpirationDate": "2024-04-29T12:00:00Z",
    "packageType": "box",
    "pickUpPointId": 1
}'
```

### Отдать заказ курьеру (Удалить с ПВЗ)
```bash
curl  -X DELETE \
http://127.0.0.1:9010/api/v1/order/1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получить список возвратов
```bash
curl -k -X GET \
https://127.0.0.1:9010/api/v1/orders/returns/2 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Получить список заказов клиента
```bash
curl -k -X GET \
https://127.0.0.1:9010/api/v1/clients/{CLIENT_ID}/orders?num_of_last_orders=1 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx'
```

### Отдать заказ клиенту
```bash
curl -k -X PUT \
https://127.0.0.1:9010/api/v1/orders/issue \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "ordersIds": [12346, 12347]
}'
```

### Принять возврат от клиента
```bash
curl -k -X PUT \
https://127.0.0.1:9010/api/v1/orders/return \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46cGFzc3dvcmQx' \
-d '{
    "orderId": 12345,
    "clientId": 67890
}'
```
