# Job Queue System in Go

A concurrent job queue system built with Go, Gin, GORM, and PostgreSQL. It simulates job processing using worker pools and channel-based queues, enabling scalable background task execution.

---

## Features

* REST API to create and retrieve jobs
* Asynchronous job processing using goroutines and channels
* Job status tracking: `todo`, `in_progress`, `done`, `failed`
* Result field for processed output
* GORM ORM for DB access
* PostgreSQL as the database

---

## Project Structure

```bash
job-queue/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── api/                 # HTTP handlers
│   ├── db/                  # DB initialization
│   ├── model/               # Job model
│   ├── queue/               # JobQueue and worker logic
│   └── service/             # Business logic (job creation, updates)
├── go.mod
└── README.md
```

---

## Tech Stack

* [Go](https://golang.org/) – Concurrency-friendly backend
* [Gin](https://github.com/gin-gonic/gin) – Web framework
* [GORM](https://gorm.io/) – ORM
* [PostgreSQL](https://www.postgresql.org/) – Relational database

---

## Setup Instructions

1. **Clone the Repo**

```bash
git clone https://github.com/flames31/job-queue.git
cd job-queue
```

2. **Install Dependencies**

```bash
go mod tidy
```

3. **Set up `.env` file**

```
DB_USER_="<user_name>"
DB_PASSWORD="<user_password>"

```

4. **Run PostgreSQL Locally or using below docker image**

```bash
docker run --name jobqueue-db -p 5432:5432 -e POSTGRES_PASSWORD=yourpass -e POSTGRES_USER=youruser -e POSTGRES_DB=jobqueue -d postgres
```

5. **Run the App**

```bash
go run ./cmd/server
```

---

## API Endpoints

### `POST /jobs`

Create a new job.

```json
{
  "description": "Process this task"
}
```

### `GET /jobs`

List all jobs.

### `GET /jobs/:id`

Get a specific job by ID.

---

## How It Works

* When a job is created, it’s saved to the DB with `status = todo`
* It's then sent to the job queue channel
* Workers (goroutines) read from the channel, mark the job `in_progress`, simulate work, and then update the status to `done` or `failed`

---

## Future Enhancements

* Retry logic for failed jobs
* Scheduled/Delayed jobs
* Job priority queue
* Metrics (prometheus) & logging
* Web UI to visualize job states

---
