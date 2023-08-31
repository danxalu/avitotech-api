# Запустить

### бд+сервер в фоновом режиме
``` shell
docker compose up db server -d
```
---


## Ручки
ask `/api/user/check` method `POST`
```json
{
  "user_id": 6
}
```
ask `/api/user` method `POST`
```json
{
  "user_name": "denchik"
}
```
ans `/api/user/check` method `POST`
```json
{
  "segment": [
    "AVITO_DISCOUNT_30", 
    "AVITO_DISCOUNT_50"
  ],
}
```

ask `/api/user/change` method `POST`
```json
{
  "user_id": 6,
  "add_user": [2, 4, 6],
  "delete_user": [1, 3, 5]
}
```
ans `/api/user/change` method `POST`
```json
{
  "user_id": 6,
  "segment": [1, 2, 3, 4],
}
```
ask `/api/segment` method `POST`
```json
{
  "slug": "AVITO_DISCOUNT_30"
}
```
ans `/api/segment` method `POST`
```json
{
  "segment_id": 2,
  "slug": "AVITO_DISCOUNT_50",
}
```
ask `/api/segment` method `DELETE`
```json
{
  "slug": "AVITO_DISCOUNT_50",
}
```
ans `/api/segment` method `DELETE`
```json
{
  "segment_id": 2,
  "slug": "AVITO_DISCOUNT_50",
}
```
ans `/api/user` method `POST`
```json
{
  "user_id": 6,
}
```
