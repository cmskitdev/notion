package types

import (
	"encoding/json"
	"fmt"

	"github.com/cmskitdev/notion/id"
)

// UserID represents a unique identifier for a Notion user.
type UserID id.ID

// String returns the string representation of the UserID.
//
// Returns:
// - string: The string representation of the UserID.
func (p *UserID) String() string {
	return string(*p)
}

// ParseUserID parses a string as a Notion UserID, accepting both dashed and undashed formats.
// It returns the UserID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a UserID.
//
// Returns:
// - UserID: The canonical, dashed UserID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParseUserID(id string) (UserID, error) {
	userID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return UserID(userID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for UserID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (p *UserID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseUserID(s)
	if err != nil {
		return fmt.Errorf("invalid UserID in JSON: %w", err)
	}

	*p = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for UserID.
//
// Returns:
// - []byte: JSON-encoded UserID as a string.
// - error: Marshaling error, if any.
func (p UserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// Validate ensures the UserID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := userID.Validate(); err != nil {
//	    log.Printf("Invalid UserID: %v", err)
//	}
func (p UserID) Validate() error {
	_, err := ParseUserID(string(p))
	return err
}

// User represents a user in Notion, which can be either a person or a bot.
// Users are referenced throughout the API for creation, editing, and mention operations.
type User struct {
	BaseObject
	ID        UserID   `json:"id"`
	Type      UserType `json:"type"`
	Name      *string  `json:"name"`
	AvatarURL *string  `json:"avatar_url"`
	Person    *Person  `json:"person,omitempty"`
	Bot       *Bot     `json:"bot,omitempty"`
}

// UserType represents the type of user account.
type UserType string

const (
	// UserTypePerson represents a human user in Notion.
	UserTypePerson UserType = "person"
	// UserTypeBot represents a bot user in Notion.
	UserTypeBot UserType = "bot"
)

// Person represents a human user in Notion with email information.
type Person struct {
	Email *string `json:"email,omitempty"`
}

// Bot represents a bot user in Notion with owner and workspace information.
type Bot struct {
	Owner           *BotOwner        `json:"owner,omitempty"`
	WorkspaceName   *string          `json:"workspace_name,omitempty"`
	WorkspaceLimits *WorkspaceLimits `json:"workspace_limits,omitempty"`
}

// BotOwner represents the owner of a bot, which can be a user or workspace.
type BotOwner struct {
	Type      BotOwnerType `json:"type"`
	User      *User        `json:"user,omitempty"`
	Workspace *bool        `json:"workspace,omitempty"`
}

// BotOwnerType represents the type of bot owner.
type BotOwnerType string

const (
	// BotOwnerTypeUser represents a user as the owner of a bot.
	BotOwnerTypeUser BotOwnerType = "user"
	// BotOwnerTypeWorkspace represents a workspace as the owner of a bot.
	BotOwnerTypeWorkspace BotOwnerType = "workspace"
)

// WorkspaceLimits represents limits and restrictions for a bot's workspace.
type WorkspaceLimits struct {
	MaxFileUploadSizeInBytes int64 `json:"max_file_upload_size_in_bytes"`
}

// PartialUser represents a minimal user reference used in API responses
// where full user information is not needed.
type PartialUser struct {
	ID     UserID     `json:"id"`
	Object ObjectType `json:"object"`
}

// Validate ensures the User has valid required fields based on its type.
//
// Returns:
// - error: Validation error if required fields are missing or invalid, nil if valid.
//
// Example:
//
//	user := &User{Type: UserTypePerson}
//	if err := user.Validate(); err != nil {
//	    log.Printf("Invalid user: %v", err)
//	}
func (u *User) Validate() error {
	if u.Object != ObjectTypeUser {
		return fmt.Errorf("object type must be 'user', got: %s", u.Object)
	}

	switch u.Type {
	case UserTypePerson:
		if u.Person == nil {
			return fmt.Errorf("person field is required for user type 'person'")
		}
		if u.Bot != nil {
			return fmt.Errorf("bot field should be nil for user type 'person'")
		}
	case UserTypeBot:
		if u.Bot == nil {
			return fmt.Errorf("bot field is required for user type 'bot'")
		}
		if u.Person != nil {
			return fmt.Errorf("person field should be nil for user type 'bot'")
		}
		return u.Bot.Validate()
	default:
		return fmt.Errorf("unknown user type: %s", u.Type)
	}

	return nil
}

// Validate ensures the Bot has valid configuration.
//
// Returns:
// - error: Validation error if configuration is invalid, nil if valid.
func (b *Bot) Validate() error {
	if b.Owner != nil {
		return b.Owner.Validate()
	}
	return nil
}

// Validate ensures the BotOwner has the correct fields set for its type.
//
// Returns:
// - error: Validation error if required fields are missing, nil if valid.
func (bo *BotOwner) Validate() error {
	switch bo.Type {
	case BotOwnerTypeUser:
		if bo.User == nil {
			return fmt.Errorf("user field is required for bot owner type 'user'")
		}
		if bo.Workspace != nil {
			return fmt.Errorf("workspace field should be nil for bot owner type 'user'")
		}
		return bo.User.Validate()
	case BotOwnerTypeWorkspace:
		if bo.Workspace == nil || !*bo.Workspace {
			return fmt.Errorf("workspace must be true for bot owner type 'workspace'")
		}
		if bo.User != nil {
			return fmt.Errorf("user field should be nil for bot owner type 'workspace'")
		}
	default:
		return fmt.Errorf("unknown bot owner type: %s", bo.Type)
	}
	return nil
}

// IsBot returns true if the user is a bot.
func (u *User) IsBot() bool {
	return u.Type == UserTypeBot
}

// IsPerson returns true if the user is a person.
func (u *User) IsPerson() bool {
	return u.Type == UserTypePerson
}

// GetEmail returns the email address for person users.
//
// Returns:
// - string: Email address if available and user is a person, empty string otherwise.
func (u *User) GetEmail() string {
	if u.Type == UserTypePerson && u.Person != nil && u.Person.Email != nil {
		return *u.Person.Email
	}
	return ""
}

// GetDisplayName returns the display name for the user.
//
// Returns:
// - string: Display name if available, user ID if name is nil.
func (u *User) GetDisplayName() string {
	if u.Name != nil {
		return *u.Name
	}
	return u.ID.String()
}
