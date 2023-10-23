
### Запуск:
___

#### Запуск сервиса
```shell
make docker-run
make run-sub
```
#### Запуск паблишера
```shell
make run-pub
```
#### Остановка Docker
```shell
make docker-stop
```

### Работа:
___
Конфиг сервиса расположен в `subscriber/config/config.yml` и `.env`.
Конфиг паблишера в `publisher/config/config.yml`

Доступ к ордерам из кэша осуществляется через ручку `localhost:/3000/orders/{id}`.
