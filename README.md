# ABC Cinema Schedule and Tickets

Welcome to ABC Cinema Schedule and Tickets project! This is a web application designed to offer an engaging and user-friendly platform for browsing movie schedules, purchasing tickets, and managing movie-related data. It includes features for both regular users and administrators, ensuring a comprehensive cinema experience.

## Features

### For Users
- **Browse Movies**: View movies and cinema schedules created by the admin.
  - Browse by genre
  - Browse by director
  - Filter movies by genre
  - Filter movies by ingredients
  - Search by movie title
- **User Accounts**: Register, log in, and manage user accounts.
- **Ticket Purchasing**: Buy tickets for movie schedules.
- **Movie Bookmarks**: Bookmark favorite movie schedules.
- **Movie Ratings**: Rate movies you have watched.

### For Admins
- **Admin Panel**: A separate interface for managing content.
- **Movie Schedule Creation**: 
  - Upload multiple images (select a featured image for the thumbnail)
  - Set one director of the movie
  - Set multiple stars (actors/actresses)
  - Set movie duration
  - Define movie genre
  - Set movie title and description
- **Manage Directors**: Create, read, update, and delete movie directors.
- **Manage Stars**: Create, read, update, and delete movie stars.

## Technical Requirements

- **Authentication**: JWT authentication system for secure access.
- **Backend**: 
  - Local instance of Hasura (Docker)
  - Postgres with custom functions, events, triggers, and computed/generated properties
- **Frontend**: 
  - Vue3 with Vite
  - Vue Apollo for GraphQL integration
  - Vee-validate for form validations
  - Nuxt 3
  - TailwindCSS for styling
- **Backend Language**: Golang

## Installation

### Prerequisites

- Docker
- Node.js
- hasura
- Golang

### Setup
1. **Clone the Repository**
   ```bash
   git clone https://github.com/Eyumul/project_menab.git
2. **Start the docker engine**
 start docker and open the hasura console
3. **Head to the backend directory**
 cd to backend
 cd to imageupload
 go run main.go 
 cd ..
 cd to go-proxy-server
 go run main.go
 cd ..
 cd to email-sender
 go run main.go
4. **Head to frontend directory**
 cd frontend
 npm install ( yarn install )
 npm run dev ( yarn dev )

## Usage

### Access the Application: 
Open your browser and go to http://localhost:3000 to view the application.

### Admin Panel: 
Access the admin panel at http://localhost:3000/admin (credentials required).

## Screenshots
![ABC Cinema site client side screenshot](https://github.com/user-attachments/assets/1abc6e37-3d80-4203-97f0-80054231505f)
![ABC Cinema site admin side screenshot](https://github.com/user-attachments/assets/f78de568-2238-4a6f-9e09-7586de6ecec6)
