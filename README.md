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

1. Problem Statement

The system must notify users weekly about apartments that match their selected filters.

Constraints:

Millions of users (8M+)

Large apartment dataset

Avoid expensive SELECT queries per user

Weekly delivery must be fast and scalable

A naive solution (query apartments per user weekly) would result in N × M queries, which does not scale.

2. Solution

We use precomputation + read optimization.

Filters and apartments are stored in PostgreSQL (source of truth)

When a user creates or updates a filter:

Matching apartments are computed immediately

Results are stored in Cassandra as precomputed matches

Weekly notifications read directly from Cassandra (O(1) access)

This shifts cost from read-time to write-time, which is more predictable and scalable.

Besides, precompute job, because new apartments may be added after filters were created we use Precompute works on filters, not users. Filters define the matching logic, and each filter belongs to a user. The repository returns (userID, filter) pairs so the service stays independent from database schema.
