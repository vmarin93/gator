# Gator 🐊

Gator is a command-line RSS feed aggregator built in Go. It allows users to follow their favorite blogs, news sites, and podcasts, collecting posts into a PostgreSQL database for easy browsing directly from the terminal.

This project was developed as part of the [Boot.dev](https://www.boot.dev) curriculum to practice Go programming, SQL integration, and building long-running services.

## Learning Goals

- **Go & PostgreSQL**: Integration of a Go application with a relational database.
- **Type-Safe SQL**: Using `sqlc` to generate Go code from raw SQL queries.
- **Concurrency & Services**: Implementing a service that runs continuously to perform background tasks.

## Features

- **User Management**: Register and login to manage your feed subscriptions.
- **Feed Aggregation**: Add RSS feeds and follow/unfollow feeds added by others.
- **Background Scrapping**: A long-running service that periodically fetches new posts from followed feeds.
- **Post Browsing**: View summaries and descriptions of collected posts.

## Prerequisites

To run Gator, you need the following installed on your system:

- **Go**: [Install Go](https://go.dev/doc/install) (latest version recommended).
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/).

## Installation & Configuration

Gator requires a configuration file located at `~/.gatorconfig.json` to store your database connection string and the current active user.

1. Create the file:
   ```bash
   touch ~/.gatorconfig.json
   ```

2. Add your database URL (and an empty string for the user):
   ```json
   {
     "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
     "current_user_name": ""
   }
   ```
   *Replace `username`, `password`, and `gator` with your PostgreSQL credentials and database name.*

Once that's done, clone the repository so you have access to the SQL migration files:

```bash
git clone https://github.com/vmarin93/gator.git
cd gator/sql/schema
```

Then run the database migrations with Goose:

```bash
goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```
   *Replace the connection string with your actual credentials.*

And lastly, we can build the executable:
```bash
go build
```

## Usage

Once configured, you can start using Gator. Here are some of the available commands:

### User Commands
- `gator register <name>`: Create a new user account and log in.
- `gator login <name>`: Switch to an existing user account.
- `gator users`: List all registered users.
- `gator reset`: Clear all users and their data from the database.

### Feed Commands
- `gator addfeed <name> <url>`: Add a new RSS feed to the system (requires login).
- `gator feeds`: List all feeds added by users.
- `gator follow <url>`: Follow an existing feed (requires login).
- `gator unfollow <url>`: Unfollow a feed (requires login).
- `gator following`: List all feeds the current user is following.

### Aggregation and Browsing
- `gator agg <duration>`: Start the long-running aggregator service. Example: `gator agg 1m` to fetch feeds every minute.
- `gator browse [limit]`: View the latest posts from feeds you follow. Default limit is 2.
