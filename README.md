# Chat-Service

Flow of User Actions
	1.	User signs up → Stored in PostgreSQL.
	2.	User searches for a friend → Query users table.
	3.	User sends friend request → Insert into friendships table.
	4.	Friend accepts request → Friendship updated in DB.
	5.	User sends first message → WebSocket connection opens.
	6.	Messages exchanged in real-time → Stored in MongoDB.
	7.	If inactive for X minutes → WebSocket disconnects.
	8.	User can retrieve chat history from MongoDB anytime.

WHAT IS A MIDDLEWARE ?
// Use middleware like
<!-- router.POST("/friends/:id", ratelimit.PerUser(10, time.Minute), SendFriendRequest) -->

Tasks to do:
- User CRUD operations
- Friendship system
- Authentication

Real-Time Collaboration
Tech Stack: Go, WebSockets, MongoDB

- WebSocket chat rooms

Message persistence

Presence tracking

File upload system

Typing indicators

Implementation Steps:

Implement WebSocket server in Go

Create MongoDB schema for message history

Integrate with user service for authentication

Add Redis for online status tracking

Implement S3-compatible file storage

Create message queue for notifications



FUTURE:

mplement Google/GitHub OAuth

Add CORS middleware

Configure Content Security Policy

Add request validation layers

Perform vulnerability scanning

Set up HTTPS with Let's Encrypt


Create PostgreSQL schema for users/roles

Implement password hashing with Argon2

Build REST API endpoints for user management

Add Redis integration for rate limiting

Create middleware for JWT verification

Implement audit logging to MongoDB


GITHUB ACTIONS



System Architecture

Tech Stack
	•	Backend: Go (Gorilla WebSockets, PostgreSQL driver, MongoDB driver)
	•	Database: PostgreSQL (Users, Friendships), MongoDB (Chat Messages)
	•	WebSockets: Real-time communication between users
	•	Docker: Containerized deployment with PostgreSQL & MongoDB
	•	Authentication: JWT-based authentication for secure API access
	•	Caching (Optional): Redis for session tracking & WebSocket timeouts