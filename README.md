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