Apartments Platform

Backend service for managing apartments, user filters, and real-time matching using PostgreSQL and Elasticsearch percolators.

This service allows:

- Creating apartments with attributes (price, rooms, district, city, etc.)
- Users to define filters for apartments
- Automatic percolation in Elasticsearch to find users whose filters match newly added apartments
- Optional notifications for matched users
  `
  Mental Model

- Source of Truth: PostgreSQL database - All permanent data about apartments and user filters is stored in the database. Ensures reliable persistence and backup.

- Elasticsearch: Apartments and user filters are indexed in Elasticsearch to support percolation queries. When a new apartment is added, Elasticsearch runs percolator queries to find all users whose filters match the apartment.

- Future improvements: If indexing an apartment in Elasticsearch fails, consider queuing the apartment in Redis or Kafka to retry indexing later (all of these are written as a comments in the code)

Architecture

The project is structured with clear separation of concerns, following SOLID principles.

We use modules to have clean architecture. Each layer has a single responsibility and can operate independently. Dependencies are injected via constructors, promoting loose coupling.

- Request layer: Defines the incoming request structures (DTOs). Handles validation and parsing of input data.
  Domain Layer

- Domain layer: Contains the core business logic and models. Defines interfaces for repositories, services, and other dependencies. Fully decoupled from transport or database implementations.

- Initialization - The init functions in Go start the necessary services, repositories, and connections (DB, Elasticsearch, caches). Ensures that all components are ready before handling requests.

API Endpoints

- POST /apartments – Create an apartment

body :

{
"title": "Apartment 1",
"price_per_month": 1000,
"room_numbers": 3,
"bedroom_numbers": 2,
"bathroom_numbers": 1,
"district": "Saburtalo",
"city": "Tbilisi"
}

- POST /filters/:userID – Create or update a user filter

body :

{
"min_price": 500,
"max_price": 1500,
"room_numbers": [2,3],
"bedroom_numbers": [1,2],
"bathroom_numbers": [1],
"city": "Tbilisi",
"district": "Vake"
}

- POST /users - Creates a user

body :

{
"email": "randomEmail@gmail.com"
}
