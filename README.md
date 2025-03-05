# ASCII Art Web

A web-based ASCII art generator that converts user-inputted text into ASCII art using different artistic banners. This project allows users to select various banners and view the corresponding ASCII art through a web interface.

## Table of Contents

- [Project Structure](#project-structure)
- [Features](#features)
- [Usage](#usage)
- [Installation](#installation)
- [Endpoints](#endpoints)
- [Error Handling](#error-handling)
- [Technologies Used](#technologies-used)
- [Authors](#authors)

## Project Structure

```
ascii-art-web/
├── banners/          # Contains ASCII art banners
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── static/           # Contains CSS and static files
│   ├── main.css
│   ├── output.css
│   └── error.css
│
├── templates/        # HTML templates for pages
│   ├── home.html
│   ├── output.html
│   └── error.html
│
├── server.go              # Server entry point
├── routes.go              # HTTP handlers
├── ascii_generator.go     # ASCII art conversion logic
├── utils.go               # Helper functions
└── README.md              # Project documentation
```

## Features

- Generate ASCII art from user input.
- Supports three banners: `standard`, `shadow`, `thinkertoy`.
- Handles form submission via `POST` requests.
- Error handling for invalid requests, missing banners, and more.
- Displays appropriate HTTP status codes.
- User-friendly interface with responsive design.

## Usage

1. Open the application in your browser: `http://localhost:8080`
2. Enter your desired text.
3. Select a banner.
4. Click the "Generate" button to produce ASCII art.

## Installation

Ensure you have Go installed (v1.20+ recommended).

1. Clone the repository:

```bash
$ git clone https://learn.zone01oujda.ma/git/mennaas/ascii-art-web
$ cd ascii-art-web
```

2. Run the server:

```bash
$ go run main.go
```

3. Open a browser and navigate to:

```
http://localhost:8080
```

## Endpoints

| HTTP Method | Endpoint      | Description                  |
| ----------- | ------------- | ---------------------------- |
| `GET`       | `/`           | Main home page.              |
| `POST`      | `/ascii-art`  | Generate ASCII art.          |
| `POST`      | `/ascii-art1` | Regenerate from output page. |
| `GET`       | `/static/*`   | Serves static CSS files.     |

## Error Handling

The application returns appropriate HTTP status codes:

| Status Code                 | Meaning                               |
| --------------------------- | ------------------------------------- |
| `200 OK`                    | Successful request.                   |
| `400 Bad Request`           | Invalid input or exceeding limits.    |
| `404 Not Found`             | Invalid endpoint or missing resource. |
| `405 Method Not Allowed`    | Unsupported HTTP method.              |
| `500 Internal Server Error` | Server-side error.                    |

### Input Validation

- Only printable ASCII characters (32-126) are accepted.
- Text length is limited to **1000 characters**.
- Empty input triggers an error notification.

## Technologies Used

- **Go** (for backend server)
- **HTML/CSS** (for frontend user interface)
- **net/http** (for handling HTTP requests)
- **html/template** (for rendering HTML pages)

## Authors

- [En-naas Marouane] - Initial development
- [El Mahmoudi Abderrahman] - Initial development


