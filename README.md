# Novel API

This repository contains the backend API for Novel data.

## Overview
The Novel Backend API is designed to facilitate the management and distribution of novel data, such as title, author, genre, etc. It provides endpoints for CRUD (Create, Read, Update, Delete) operations on novels, authors and genres.

## Features

- **Author Management**: Enables CRUD operations for authors, providing a seamless way to manage author information.
- **Genre Classification**: Supports genre classification for novels, allowing users to categorize novels based on their genres.
- Novel: Supports description, relation with one or more genre and author, yiear of publication, language and others.
- **Python**: The backend API is implemented using Golang.
- **Gorm**: Utilizes the Gorm ORM.
- **gorilla/mux**: Utilizes mux to handle requests and routes.
- **Database PostgreSQL**: Uses a relational database to store novel-related data, ensuring data consistency and integrity.

## Getting Started
To get started with the Novel Backend API, follow these steps:

First, [download](https://go.dev/dl/) and install **Go**. Version `1.21` or higher is required.

**Clone the Repository**: Clone this repository to your local machine.

```bash
git clone https://github.com/felixrodrigo19/novel-api.git
cd novel-api
```

Install dependencies
```bash
go install
```

Run:
```bash
go run .\main.go
```

## Endpoints

**port**: 8000

### POST

- **Genre**: `/genre`
  **body**
  ```bash
  {
    "name": "Romance"
  }
  ```
- **Author**: `/author`
  **body**
  ```bash
  {
    "Name": "Author Name"
  }
  ```
- **Novel**: `/novel`
  **body**
  ```bash
  {
    "Title": "title",
    "Description": "Novel description",
    "Language": "Japanese",
    "Type": "Web Novel",
    "genres": [
        {
            "Id": 1
        }
    ],
    "Authors": [
        {
            "Id": 1
        }
    ],
    "Year": 2012
  }
  ```
