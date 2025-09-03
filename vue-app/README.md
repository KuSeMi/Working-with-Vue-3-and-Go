# Vue Book Management System

A modern Vue.js 3 frontend application for managing books and users with authentication and admin functionality. This single-page application (SPA) provides an intuitive interface for browsing books, user management, and administrative tasks.

## 🚀 Features

- **User Authentication**: Secure login system with token-based authentication
- **Book Management**: Browse, view, and manage books
- **Admin Panel**: Administrative interface for managing books and users
- **Responsive Design**: Modern UI with reusable form components
- **Client-side Routing**: Fast navigation with Vue Router
- **Code Quality**: ESLint integration for consistent code style

## 🛠 Technology Stack

- **Frontend Framework**: Vue.js 3.2.13 with Composition API
- **Routing**: Vue Router 4.5.1
- **Build Tool**: Vue CLI 5.0.0
- **Code Quality**: ESLint with Vue 3 essential rules
- **Notifications**: Notie 4.3.1
- **Deployment**: Docker with Caddy web server

## 📋 Prerequisites

Before running this application, make sure you have:

- **Node.js** (version 14 or higher)
- **npm** (Node Package Manager)
- **Docker** (optional, for containerized deployment)

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd vue-app
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

## 🚀 Development

### Start development server
```bash
npm run serve
```
The application will be available at `http://localhost:8080` with hot reload enabled.

### Build for production
```bash
npm run build
```
This creates optimized static files in the `dist/` directory.

### Lint and fix code
```bash
npm run lint
```

## 🏗 Project Structure

```
src/
├── components/
│   ├── forms/                 # Reusable form components
│   │   ├── CheckInput.vue     # Checkbox input component
│   │   ├── FormTag.vue        # Form wrapper component
│   │   ├── SelectInput.vue    # Select dropdown component
│   │   └── TextInput.vue      # Text input component
│   ├── Body.vue               # Main content component
│   ├── Header.vue             # Application header
│   ├── Footer.vue             # Application footer
│   ├── Book*.vue              # Book-related components
│   ├── Login*.vue             # Authentication components
│   ├── User*.vue              # User management components
│   ├── security.js            # Authentication utilities
│   └── store.js               # Application state management
├── router/
│   └── index.js               # Application routes
├── App.vue                    # Root component
└── main.js                    # Application entry point
```

## 🔐 Authentication

The application includes a security layer that:
- Manages user authentication tokens
- Protects routes with authentication checks
- Handles login/logout functionality

## 📱 Available Routes

- `/` - Home page
- `/login` - User login
- `/books` - Browse all books
- `/books/:bookName` - View specific book details
- `/admin/books` - Admin: Manage books
- `/admin/books/:bookId` - Admin: Edit specific book
- `/admin/users` - Admin: Manage users
- `/admin/users/:userId` - Admin: Edit specific user

## 🐳 Docker Deployment

The application includes Docker configuration for easy deployment:

1. **Build the application**
   ```bash
   npm run build
   ```

2. **Start with Docker Compose**
   ```bash
   docker-compose up -d
   ```

The application will be available at `http://localhost:80` served by Caddy web server.

## 🔧 Configuration

### Browser Support
- Modern browsers (>1% market share)
- Last 2 versions of major browsers
- Excludes Internet Explorer 11

### ESLint Configuration
- Vue 3 essential rules
- Babel parser for modern JavaScript syntax
- Customizable rules in `package.json`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin feature/new-feature`
5. Submit a pull request

## 📝 Development Guidelines

- Follow Vue 3 Composition API patterns
- Use ESLint rules for code consistency
- Create reusable components when possible
- Maintain separation of concerns
- Write meaningful commit messages

## 🔍 Troubleshooting

### Common Issues

1. **Port already in use**: Change the port in `vue.config.js` or stop other services
2. **Build errors**: Clear `node_modules` and reinstall dependencies
3. **Authentication issues**: Check if backend API is running and accessible

### Getting Help

- Check the [Vue CLI documentation](https://cli.vuejs.org/config/)
- Review [Vue 3 documentation](https://vuejs.org/guide/)
- Check [Vue Router documentation](https://router.vuejs.org/)

## 📄 License

This project is private and not licensed for public use.

## 🔗 Related Projects

This frontend application is designed to work with a Go backend service for complete functionality.
