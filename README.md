# Apartments Platform

Hybrid system using PostgreSQL and Cassandra to deliver apartment listings
to millions of users using precomputed filters.

## Tech Stack

- Go
- PostgreSQL
- Cassandra
- Docker
- Kubernetes

📐 Apartments Matching Platform – Design Overview

We use precomputation + read optimization.

Filters and apartments are stored in PostgreSQL (source of truth)

When a user creates or updates a filter:

Matching apartments are computed immediately

Results are stored in Cassandra as precomputed matches

Weekly notifications read directly from Cassandra (O(1) access)

This shifts cost from read-time to write-time, which is more predictable and scalable.

Besides, we use precompute job, because new apartments may be added after filters were created. Precompute works on filters, not users. Filters define the matching logic, and each filter belongs to a user. The repository returns (userID, filter) pairs so the service stays independent from database schema.

---

System Architecture
Components

1. API Layer (Echo-based REST API) - Handles HTTP requests from clients.

Routes include /users, /apartments, /filters, /health.

2. Domain Layer (Business Logic) - All business logic is delegated to domain services.

3. Storage Layer

PostgreSQL: Stores users, apartments, and filters.

Cassandra: Stores precomputed matches for fast retrieval.

4. Jobs Layer

Precompute Job: Runs periodically or triggered to precompute matches for new apartments.

Weekly Notification Job: Sends weekly updates to users based on precomputed matches.

Key Design Principles

Separation of concerns: API → Handlers → Services → Repositories → Database.

Precomputation: Shifts heavy computation to write-time.

Scalable storage: Cassandra for read-optimized precomputed matches.
