Web Forum Project README

Welcome to the Web Forum Project! This project aims to create a web forum where users can communicate, interact with posts and comments, and filter content. Below you'll find all the necessary information to understand, set up, and contribute to this project.
Objectives

This project's main objectives are as follows:

    Communication: Users can create posts and comments, associating categories with posts.
    Authentication: Users can register, log in, and have their own sessions using cookies.
    Likes and Dislikes: Registered users can like or dislike posts and comments.
    Filter: Users can filter posts by categories, created posts, and liked posts.
    SQLite: The project uses SQLite for data storage, including users, posts, comments, etc.

Getting Started

To get started with this project, follow these steps:
Prerequisites

    Install Go (Go programming language).
    Install Docker for containerization.

Installation

    Clone the project repository:

    bash

git clone https://github.com/yourusername/web-forum.git
cd web-forum

Build and run the Docker container:

bash

    docker-compose up --build

    Access the forum in your web browser at http://localhost:8080.

User Registration and Authentication

    Users can register with their email, username, and password.
    Email addresses must be unique; duplicates are not allowed.
    Passwords are securely encrypted before storage.

Communication

    Registered users can create posts and comments.
    Posts can be associated with one or more categories.
    Posts and comments are visible to all users, registered or not.

Likes and Dislikes

    Only registered users can like or dislike posts and comments.
    The number of likes and dislikes is visible to all users.

Filtering

    Users can filter posts by categories, created posts, and liked posts.
    Filtering by categories functions as subforums, focusing on specific topics.

Development Guidelines

    The codebase must follow best practices and coding standards.
    Handle website errors and HTTP status codes gracefully.
    Include unit tests where applicable.

Dependencies

This project utilizes the following Go packages and libraries:

    All standard Go packages
    sqlite3: SQLite database management
    bcrypt: Password encryption
    UUID: Universally unique identifier (for a bonus task)

Contributing

We welcome contributions to this project. To contribute, please follow these steps:

    Fork the repository.
    Create a feature branch: git checkout -b feature/your-feature-name.
    Make your changes and commit them.
    Push to your fork and submit a pull request.

License

This project is licensed under the MIT License.
Acknowledgments

Thank you for contributing to this project and helping us create a functional web forum!

If you have any questions or need assistance, please feel free to reach out to us. Happy coding!