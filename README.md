# SSO Сервис Авторизации

## Описание

Этот проект представляет собой сервис авторизации с использованием Single Sign-On (SSO), реализованный на языке Go. Он использует транспортировку через gRPC, proto файлы хранятся в проекте protos. Сервис использует JWT (JSON Web Token) для аутентификации пользователей и управления сессиями. Все учетные записи пользователей хранятся в базе данных PostgreSQL, а пароли перед сохранением хэшируются для повышения безопасности.


### Основные функции:

Сервис предоставляет три ключевые функции для работы с пользователями:

1. **Register**: 
   - **RPC**: `Register (RegisterRequest) returns (RegisterResponse)`
   - Позволяет новым пользователям регистрироваться в системе, создавая учетную запись и сохраняя необходимые данные.

2. **Login**: 
   - **RPC**: `Login (LoginRequest) returns (LoginResponse)`
   - Позволяет пользователям аутентифицироваться, проверяя их учетные данные и возвращая JWT токен для дальнейшего использования.

3. **IsAdmin**: 
   - **RPC**: `IsAdmin (IsAdminRequest) returns (IsAdminResponse)`
   - Проверяет, является ли текущий пользователь администратором, предоставляя соответствующий ответ.
  
### Протоколы и Сообщения

Сервис использует следующие сообщения для обмена данными:

- **RegisterRequest**:
  ```protobuf
  message RegisterRequest {
    string email = 1; // Email of the user to register
    string password = 2; // Password of the user to register
  }
  ```

- **RegisterResponse**:
  ```protobuf
  message RegisterResponse {
    int64 user_id = 1; // User ID of the registered user
  }
  ```

- **LoginRequest**:
  ```protobuf
  message LoginRequest {
    string email = 1; // Email of the user to login
    string password = 2; // Password of the user to login
    int32 app_id = 3; // ID of the app to login to
  }
  ```

- **LoginResponse**:
  ```protobuf
  message LoginResponse {
    string token = 1; // Auth token of the logged user
  }
  ```

- **IsAdminRequest**:
  ```protobuf
  message IsAdminRequest {
    int64 user_id = 1; // User ID to validate
  }
  ```

- **IsAdminResponse**:
  ```protobuf
  message IsAdminResponse {
    bool is_admin = 1; // Indicates whether the user is an admin
  }
  ```

### Технологии:

- **Язык программирования**: Go
- **Транспортировка**: gRPC
- **Файлы протоколов**: Протоколы описаны и хранятся в проекте protos.
- **Аутентификация**: JWT (JSON Web Token)
- **База данных**: PostgreSQL (учетные записи пользователей)
