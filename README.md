# Chat-Service

Tech Stack
	•	Backend: Go (Gorilla WebSockets, PostgreSQL driver, MongoDB driver)
	•	Database: PostgreSQL (Users, Friendships), MongoDB (Chat Messages)
	•	WebSockets: Real-time communication between users
	•	Docker: Containerized deployment with PostgreSQL & MongoDB
	•	Authentication: JWT-based authentication for secure API access
	•	Caching (Optional): Redis for session tracking & WebSocket timeouts

Flow of User Actions
	1.	User signs up → Stored in PostgreSQL.
	2.	User searches for a friend → Query users table. ??
	3.	User sends friend request → Insert into friendships table.
	4.	Friend accepts request → Friendship updated in DB.
	5.	User sends first message → WebSocket connection opens.
	6.	Messages exchanged in real-time → Stored in MongoDB.
	7.	If inactive for X minutes → WebSocket disconnects.
	8.	User can retrieve chat history from MongoDB anytime.


Middleware
- add logging (Use case: Debugging slow requests.)
- add rate limiting(Use case: Prevents spamming friend requests/messages.)
- add caching
- Add CORS middleware (Use case: Enables frontend (React, Vue, etc.) to call your API.)
- Add web socket management (Use case: Ensures users don’t open too many WebSocket connections.)

Tasks to do before chat system is implemented:
- User CRUD operations
- Friendship system
- Authentication



Things to do for chat system:
- Message persistence
- Presence tracking(Optional)
- File upload system
- Typing indicators(Optional)

Implementation Steps:
- Implement WebSocket server in Go
- Create MongoDB schema for message history
- Integrate with user service for authentication
- Add Redis for online status tracking
- Implement S3-compatible file storage
- Create message queue for notifications

FUTURE:
- Implement Google/GitHub OAuth
- Authorization(ROLE BASED)
- Configure Content Security Policy
- Add request validation layers
- Perform vulnerability scanning
- Set up HTTPS with Let's Encrypt
- GITHUB ACTIONS
- GO AIR-VERSE
- System Architecture

