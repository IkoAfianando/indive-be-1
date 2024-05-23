# My Golang Fiber Application

This is a Golang application built with the Fiber web framework. It provides authentication functionalities including registration, login, and OAuth authentication with Google

## Installation

### Prerequisites

Before getting started, ensure you have the following installed on your machine:

- Go (version >= 1.16): [Installation Guide](https://golang.org/doc/install)
- MySQL: [Installation Guide](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
- Mailtrap account for email testing: [Mailtrap](https://mailtrap.io/)

### Steps

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/your-repository.git
   cd your-repository
   ```

2. **Install dependencies**:

   ```bash
    go mod tidy
    ```

3. **Setup environment variables**:

   Create a `.env` file in the root directory of the project. Add the following environment variables:

   ```env
    GOOGLE_CLIENT_ID=
    GOOGLE_CLIENT_SECRET=
    SMTP_USER=
    SMTP_PASSWORD=
    SMTP_URL=
    SMTP_PORT=
    MYSQL_URL=
   ```

    Replace the values with your own. You can get the `GOOGLE_CLIENT_ID` and `GOOGLE_CLIENT_SECRET` by creating a new project on the Google Developer Console. For the `SMTP` variables, you can get them from your Mailtrap account. The `MYSQL_URL` should be in the format `username:password@tcp(
4. **Run the application**:

   ```bash
   go run main.go
   ```

   The application should be running on `http://localhost:3000`.

## API Endpoints

The following endpoints are available:

- `POST /register`: Register a new user
- `POST /verify-email`: Verify a user's email
- `POST /login`: Login a user
- `GET /api/auth/google`: Redirect to Google OAuth
- `GET /api/auth/google/callback`: Callback URL for Google OAuth
- `GET /api/dashboard`: Get user's dashboard


## License
This project is open source and available under the [MIT License](LICENSE).
```
Feel free to reach out to me if you have any questions or suggestions. You can also contribute to the project by creating a pull request. Thank you for reading.
```
