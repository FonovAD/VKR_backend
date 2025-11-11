# VKR


## API Endpoints

Приложение использует следующие эндпоинты бэкенда:
**Также в проекте настроен Swagger**

### Организации
- `GET /api/v1/organization` — список организаций
- `GET /api/v1/organization/:id` — организация по ID
- `POST /api/v1/organization` — создание организации
- `PUT /api/v1/organization/:id` — обновление организации
- `DELETE /api/v1/organization/:id` — удаление организации
- `GET /api/v1/organization/search/by-inn?inn=...` — поиск организации по ИНН

### Музеи
- `GET /api/v1/museum` — список музеев
- `GET /api/v1/museum/:id` — музей по ID
- `POST /api/v1/museum` — создание музея
- `PUT /api/v1/museum/:id` — обновление музея
- `DELETE /api/v1/museum/:id` — удаление музея
- `GET /api/v1/museum/search/by-inn?inn=...` — поиск музея по ИНН
- `GET /api/v1/museum/owner/:owner_id` — список музеев по организации

### Деятельность музеев
- `GET /api/v1/activity` — список записей
- `GET /api/v1/activity/search/by-inn?inn=...` — деятельность по ИНН
- `POST /api/v1/activity` — создание записи
- `PUT /api/v1/activity` — обновление записи
- `DELETE /api/v1/activity?inn=...&activity_type_id=...&visitor_category=...&year=...` — удаление записи

### Трудовые ресурсы
- `GET /api/v1/labor/organization/:org_id` — трудовые ресурсы по ID организации
- `GET /api/v1/labor/search/by-inn?inn=...` — трудовые ресурсы по ИНН
