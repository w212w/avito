## Структура проекта

В данном репозитории размещено тестовое задание "Avito". Имеются файлы проекта, также папка с заданием - "задание", в ней рамещено само задание и спецификация OpenAPI. В файле Readme указана структура проекта с описанием, а также начальные условия реализации сервиса. Проект организован следующим образом:

### `avito/`

Корневая директория проекта, содержащая все исходные файлы.

#### `handlers/`

В этой директории находятся обработчики HTTP-запросов.

- **`bids.go`**  
  Обработчик для операций, связанных с предложениями.
  
- **`ping.go`**  
  Обработчик для проверки доступности сервера.

- **`tenders.go`**  
  Обработчик для операций, связанных с тендерами.

#### `models/`

В этой директории находятся определения моделей данных.

- **`bid.go`**  
  Модель данных для предложений.

- **`tender.go`**  
  Модель данных для тендеров.

#### `repository/`

В этой директории содержатся файлы для работы с базой данных.

- **`bids_repository.go`**  

- **`postgres.go`**  
  Файл для инициализации подключения к PostgreSQL.

- **`tenders_repository.go`**  

#### `main.go`

Главный файл приложения, который запускает HTTP сервер и устанавливает обработчики для API.

## Задание
В папке "задание" размещена задача.





