# Privatenote

Сервис приватных записей, которые живут время указанное при их создании и доступны по ссылке только указанное время.

## Технологический стек

### Backend
go, redis
### Frontend
vue

## Api
- service/ttl - массив времени жизни записи
- note/create - добавляем заметку
- note/list - список заметок
- url/{shorturl} - просмотр заметки
- user/register - регистрация пользователя
- user/login - авторизация пользователя
- note/get/{shorturl} - посмотреть запись в ui
- note/edit/{shorturl} - обновляем данные  в ui

## Схема данных

### Пользователи
- usersm - коллекция email:pass(md5)
### Заметки
Создаются просто hmap в redis со значениями
- User        string `json:"email"`
- Number      string `json:"number"`
- Text        string `json:"text"`
- Ttl         string `json:"ttl"`
- Title       string `json:"title"`
- Description string `json:"description"`

### В корень пишется email пользователя с uid note и ttl