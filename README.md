# Gator - RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator written in Go. It allows you to manage RSS feeds, follow/unfollow feeds, and browse posts from your followed feeds.

## Prerequisites

Before you can run Gator, you need to have the following installed:

1. **Go** (version 1.24.2 or later)
   - Download and install from [https://golang.org/dl/](https://golang.org/dl/)
   - Verify installation: `go version`

2. **PostgreSQL**
   - Download and install from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
   - Make sure PostgreSQL is running and you have a database created
   - You'll need the database connection URL

## Installation

Install Gator using Go's built-in package manager:

```bash
go install github.com/mjossany/Gator@latest
```

After installation, make sure your `$GOPATH/bin` (or `$GOBIN`) is in your system's `$PATH` so you can run `gator` from anywhere.

## Configuration

Before using Gator, you need to set up a configuration file:

1. Create a `.gatorconfig.json` file in your home directory
2. Add your database connection details:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace the `db_url` with your actual PostgreSQL connection string:
- `username`: Your PostgreSQL username
- `password`: Your PostgreSQL password
- `localhost:5432`: Your PostgreSQL host and port
- `gator`: Your database name

The `current_user_name` field will be automatically set when you register or login.

## Usage

### Basic Commands

#### User Management
```bash
# Register a new user
gator register <username>

# Login as an existing user
gator login <username>

# List all users
gator users

# Reset the database (removes all data)
gator reset
```

#### Feed Management
```bash
# Add a new RSS feed (requires login)
gator addfeed <feed_name> <feed_url>

# List all feeds
gator feeds

# Follow an existing feed (requires login)
gator follow <feed_url>

# List feeds you're following (requires login)
gator following

# Unfollow a feed (requires login)
gator unfollow <feed_url>
```

#### Content Aggregation
```bash
# Start the aggregator to fetch posts from feeds
gator agg <duration>

# Browse posts from your followed feeds (requires login)
gator browse [limit]
```

### Example Workflow

1. **Register a user:**
   ```bash
   gator register john
   ```

2. **Add an RSS feed:**
   ```bash
   gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
   ```

3. **Start the aggregator (in a separate terminal):**
   ```bash
   gator agg 1m
   ```
   This will fetch new posts every minute.

4. **Browse your feed posts:**
   ```bash
   gator browse 10
   ```
   This will show the latest 10 posts from your followed feeds.

## Features

- **User Management**: Register and switch between multiple users
- **Feed Management**: Add, follow, and unfollow RSS feeds
- **Automatic Aggregation**: Continuously fetch new posts from your followed feeds
- **Post Browsing**: View the latest posts from your feeds
- **PostgreSQL Storage**: All data is stored in a PostgreSQL database
- **Concurrent-Safe**: Built with Go's excellent concurrency features

## Development

If you want to contribute or run the project in development mode:

1. Clone the repository:
   ```bash
   git clone https://github.com/mjossany/Gator.git
   cd Gator
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run in development mode:
   ```bash
   go run . <command> [args...]
   ```

Note: For production use, always use the installed `gator` binary rather than `go run .`

## Database Schema

Gator uses PostgreSQL with the following main tables:
- `users`: Store user information
- `feeds`: Store RSS feed information
- `posts`: Store individual posts from feeds
- `feed_follows`: Track which users follow which feeds

The database schema is managed using SQL migrations in the `sql/schema/` directory.

## License

This project is open source. Feel free to contribute!
