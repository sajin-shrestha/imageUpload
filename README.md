# Image Upload Application

This repository contains a web application for uploading images, consisting of a frontend built with React and Vite, and a backend implemented in Golang.

## Features

- **Image Upload**: Securely upload images.
<!-- - **Responsive Design**: Mobile-friendly UI for seamless usage. -->
- **Backend API**: Golang-based API for image handling and storage.

## Technologies Used

### Frontend

- **React**: JavaScript library for building user interfaces.
- **Vite**: Next-generation frontend tooling for React applications.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **Axios**: Promise-based HTTP client for making requests to the backend API.

### Backend

- **Golang**: Programming language for building scalable backend services.
- **Gorilla Mux**: Powerful HTTP router and dispatcher for Golang web servers.
- **GORM**: ORM library for Golang, used for database interactions.
- **PostgreSQL**: Relational database for storing image metadata.

## Folder Structure

- **frontend/**: Contains the React frontend code.
- **backend/**: Contains the Golang backend code.

## Getting Started

To run the application locally, follow these steps:

### Prerequisites

- Node.js (with npm or yarn)
- Golang
- PostgreSQL

### Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/image-upload.git
   cd image-upload

   ```

2. **Create a .env file in root folder with following data:**

   ```bash
    DB_HOST=localhost
    DB_USER=<database-username>
    DB_PASSWORD=<database-password>
    DB_NAME=<database-name>
    DB_PORT=5432
   ```
