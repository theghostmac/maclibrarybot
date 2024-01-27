package commands

import "github.com/theghostmac/maclibrarybot/internal/searcher"

// Define the state of the conversation

type ConversationState int

const (
	StateNone ConversationState = iota
	StateAwaitingBookSearch
	// Add more states as needed
)

// UserSession holds the state of conversation with a specific user.
type UserSession struct {
	State             ConversationState
	LastSearchResults []searcher.Book
}

// Sessions holds the state for all conversations, keyed by user ID.
var Sessions = make(map[int64]*UserSession)

// GetSession retrieves the session for a user, creating it if necessary.
func GetSession(userID int64) *UserSession {
	if session, ok := Sessions[userID]; ok {
		return session
	}
	Sessions[userID] = &UserSession{State: StateNone}
	return Sessions[userID]
}
