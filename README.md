# MEDODS PROJECT

## Описание

Этот проект представляет собой часть сервиса аутентификации, который предоставляет два REST-маршрута для работы с токенами: выдачу пары Access и Refresh токенов и выполнение операции обновления с использованием Refresh токена.

#####Эндпоинт: /access
Этот маршрут принимает параметр запроса user_id, который является идентификатором пользователя в формате GUID и генерирует для него пару access и refresh токенов.

#####Эндпоинт: /refresh
Этот маршрут принимает два параметра user_id и refresh_token, пользователь предоставляет refresh токен, который затем проверяется на валидность, а затем используется для обновления пары токенов.
___

## Запуск

Использовать эту команду для запуска сервера и базы данных -
-- docker-compose up -d