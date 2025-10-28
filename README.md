# go-jwt-auth
Implement JWT auth in REST API

## Setup

### Environment Configuration

Create a `.env` file in the root directory with the following configuration:

```env
# Database Configuration
DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_NAME=
DATABASE_USER=
DATABASE_PASSWORD=your_password_here

# JWT Configuration
JWT_TOKEN=your_jwt_secret_here
REFERESH_TOKEN=your_refresh_secret_here

# Google OAuth Configuration
GOOGLE_CLIENT_ID=your_google_client_id_here
GOOGLE_CLIENT_SECRET=your_google_client_secret_here
GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback
```

### Database Setup

Make sure you have MySQL running and create the database:

```sql
CREATE DATABASE go_jwt_auth;
```

### Running the Application

```bash
go run main.go
```
