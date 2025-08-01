package types

import (
	"encoding/json"
	"testing"
	"time"
)

// TestIDValidation tests the validation of different ID types.
func TestIDValidation(t *testing.T) {
	validID := "550e8400-e29b-41d4-a716-446655440000"
	invalidID := "invalid-id"

	tests := []struct {
		name      string
		id        interface{ Validate() error }
		wantError bool
	}{
		{"Valid PageID", PageID(validID), false},
		{"Invalid PageID", PageID(invalidID), true},
		{"Valid DatabaseID", DatabaseID(validID), false},
		{"Invalid DatabaseID", DatabaseID(invalidID), true},
		{"Valid BlockID", BlockID(validID), false},
		{"Invalid BlockID", BlockID(invalidID), true},
		{"Valid UserID", UserID(validID), false},
		{"Invalid UserID", UserID(invalidID), true},
		{"Valid PropertyID UUID", PropertyID(validID), false},
		{"Valid PropertyID String", PropertyID("valid_prop"), false},
		{"Empty PropertyID", PropertyID(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.id.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestTimestampMarshaling tests JSON marshaling and unmarshaling of timestamps.
func TestTimestampMarshaling(t *testing.T) {
	now := time.Now().UTC()
	timestamp := Timestamp{Time: now}

	// Test marshaling
	data, err := json.Marshal(timestamp)
	if err != nil {
		t.Fatalf("Failed to marshal timestamp: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Timestamp
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal timestamp: %v", err)
	}

	// Compare times (allowing for small differences due to precision)
	// JSON marshaling truncates to seconds, so we compare with second precision
	if abs(unmarshaled.Time.Sub(now)) > time.Second {
		t.Errorf("Unmarshaled time %v does not match original %v", unmarshaled.Time, now)
	}
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

// TestParentValidation tests validation of parent relationships.
func TestParentValidation(t *testing.T) {
	validPageID := PageID("550e8400-e29b-41d4-a716-446655440000")
	validDatabaseID := DatabaseID("550e8400-e29b-41d4-a716-446655441111")

	tests := []struct {
		name      string
		parent    *Parent
		wantError bool
	}{
		{
			"Valid page parent",
			&Parent{Type: ParentTypePage, PageID: &validPageID},
			false,
		},
		{
			"Valid database parent",
			&Parent{Type: ParentTypeDatabase, DatabaseID: &validDatabaseID},
			false,
		},
		{
			"Valid workspace parent",
			&Parent{Type: ParentTypeWorkspace, Workspace: boolPtr(true)},
			false,
		},
		{
			"Invalid page parent - missing ID",
			&Parent{Type: ParentTypePage},
			true,
		},
		{
			"Invalid workspace parent - missing workspace flag",
			&Parent{Type: ParentTypeWorkspace},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.parent.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestRichTextCreation tests creation of different rich text elements.
func TestRichTextCreation(t *testing.T) {
	// Test plain text
	richText := NewTextRichText("Hello, world!", &Annotations{Bold: true})
	if richText.Type != RichTextTypeText {
		t.Errorf("Expected type %s, got %s", RichTextTypeText, richText.Type)
	}
	if richText.PlainText != "Hello, world!" {
		t.Errorf("Expected plain text 'Hello, world!', got %s", richText.PlainText)
	}
	if richText.Annotations == nil || !richText.Annotations.Bold {
		t.Error("Expected bold annotation")
	}

	// Test link
	linkText := NewLinkRichText("Google", "https://google.com", nil)
	if linkText.Type != RichTextTypeText {
		t.Errorf("Expected type %s, got %s", RichTextTypeText, linkText.Type)
	}
	if linkText.Text == nil || linkText.Text.Link == nil {
		t.Error("Expected link to be set")
	} else if linkText.Text.Link.URL != "https://google.com" {
		t.Errorf("Expected URL 'https://google.com', got %s", linkText.Text.Link.URL)
	}

	// Test user mention
	userMention := NewUserMention("550e8400-e29b-41d4-a716-446655440000")
	if userMention.Type != RichTextTypeMention {
		t.Errorf("Expected type %s, got %s", RichTextTypeMention, userMention.Type)
	}
	if userMention.Mention == nil || userMention.Mention.Type != MentionTypeUser {
		t.Error("Expected user mention")
	}
}

// TestPropertyCreation tests creation of different property types.
func TestPropertyCreation(t *testing.T) {
	// Test title property
	titleProp := NewTitleProperty("My Title")
	if titleProp.Type != PropertyTypeTitle {
		t.Errorf("Expected type %s, got %s", PropertyTypeTitle, titleProp.Type)
	}
	if len(titleProp.Title) == 0 {
		t.Error("Expected title content")
	} else if ToPlainText(titleProp.Title) != "My Title" {
		t.Errorf("Expected title 'My Title', got %s", ToPlainText(titleProp.Title))
	}

	// Test number property
	value := 42.5
	numberProp := NewNumberProperty(&value)
	if numberProp.Type != PropertyTypeNumber {
		t.Errorf("Expected type %s, got %s", PropertyTypeNumber, numberProp.Type)
	}
	if numberProp.Number == nil || numberProp.Number.Number == nil {
		t.Error("Expected number value")
	} else if *numberProp.Number.Number != 42.5 {
		t.Errorf("Expected number 42.5, got %f", *numberProp.Number.Number)
	}

	// Test checkbox property
	checkboxProp := NewCheckboxProperty(true)
	if checkboxProp.Type != PropertyTypeCheckbox {
		t.Errorf("Expected type %s, got %s", PropertyTypeCheckbox, checkboxProp.Type)
	}
	if checkboxProp.Checkbox == nil || !*checkboxProp.Checkbox {
		t.Error("Expected checkbox to be checked")
	}
}

// TestBlockCreation tests creation of different block types.
func TestBlockCreation(t *testing.T) {
	richText := []RichText{*NewTextRichText("Hello, world!", nil)}

	// Test paragraph block
	paragraph := NewParagraphBlock(richText)
	if paragraph.Type != BlockTypeParagraph {
		t.Errorf("Expected type %s, got %s", BlockTypeParagraph, paragraph.Type)
	}
	if paragraph.Paragraph == nil {
		t.Error("Expected paragraph content")
	} else if ToPlainText(paragraph.Paragraph.RichText) != "Hello, world!" {
		t.Errorf("Expected content 'Hello, world!', got %s", ToPlainText(paragraph.Paragraph.RichText))
	}

	// Test heading block
	heading := NewHeading1Block(richText)
	if heading.Type != BlockTypeHeading1 {
		t.Errorf("Expected type %s, got %s", BlockTypeHeading1, heading.Type)
	}
	if heading.Heading1 == nil {
		t.Error("Expected heading content")
	}

	// Test to-do block
	todo := NewToDoBlock(richText, false)
	if todo.Type != BlockTypeToDo {
		t.Errorf("Expected type %s, got %s", BlockTypeToDo, todo.Type)
	}
	if todo.ToDo == nil {
		t.Error("Expected to-do content")
	} else if todo.ToDo.Checked {
		t.Error("Expected to-do to be unchecked")
	}

	// Test code block
	codeText := []RichText{*NewTextRichText("fmt.Println(\"Hello, world!\")", nil)}
	codeBlock := NewCodeBlock(codeText, "go")
	if codeBlock.Type != BlockTypeCode {
		t.Errorf("Expected type %s, got %s", BlockTypeCode, codeBlock.Type)
	}
	if codeBlock.Code == nil {
		t.Error("Expected code content")
	} else if codeBlock.Code.Language != "go" {
		t.Errorf("Expected language 'go', got %s", codeBlock.Code.Language)
	}
}

// TestPageOperations tests page creation and property manipulation.
func TestPageOperations(t *testing.T) {
	// Create a database parent
	dbID := DatabaseID("550e8400-e29b-41d4-a716-446655441111")
	parent := &Parent{Type: ParentTypeDatabase, DatabaseID: &dbID}

	// Create properties
	properties := map[string]Property{
		"Name":   *NewTitleProperty("Test Page"),
		"Status": *NewSelectProperty(stringPtr("In Progress"), colorPtr(ColorYellow)),
		"Done":   *NewCheckboxProperty(false),
	}

	// Create page
	page := NewPage(parent, properties)

	// Test basic properties
	if page.Object != ObjectTypePage {
		t.Errorf("Expected object type %s, got %s", ObjectTypePage, page.Object)
	}

	if !page.IsInDatabase() {
		t.Error("Expected page to be in database")
	}

	if page.IsTopLevel() {
		t.Error("Expected page not to be top-level")
	}

	// Test title retrieval
	title := page.PropertyAccessor.GetTitleProperty("Name")
	if title == nil || title[0].Text.Content != "Test Page" {
		t.Errorf("Expected title 'Test Page', got %s", title[0].Text.Content)
	}

	// Test property operations
	if _, exists := page.GetProperty("Status"); !exists {
		t.Error("Expected Status property to exist")
	}

	page.SetProperty("Priority", *NewSelectProperty(stringPtr("High"), colorPtr(ColorRed)))
	if _, exists := page.GetProperty("Priority"); !exists {
		t.Error("Expected Priority property to be set")
	}

	if !page.RemoveProperty("Done") {
		t.Error("Expected Done property to be removed")
	}

	if _, exists := page.GetProperty("Done"); exists {
		t.Error("Expected Done property to be removed")
	}
}

// TestDatabaseOperations tests database creation and property schema manipulation.
func TestDatabaseOperations(t *testing.T) {
	title := []RichText{*NewTextRichText("Task Database", nil)}

	properties := map[string]DatabaseProperty{
		"Name": {
			Type:  PropertyTypeTitle,
			Title: &TitleConfig{},
		},
		"Status": {
			Type: PropertyTypeSelect,
			Select: &SelectConfig{
				Options: []SelectOption{
					{Name: "To Do", Color: ColorRed},
					{Name: "In Progress", Color: ColorYellow},
					{Name: "Done", Color: ColorGreen},
				},
			},
		},
	}

	database := NewDatabase(title, properties)

	// Test basic properties
	if database.Object != ObjectTypeDatabase {
		t.Errorf("Expected object type %s, got %s", ObjectTypeDatabase, database.Object)
	}

	dbTitle := database.GetTitle()
	if dbTitle != "Task Database" {
		t.Errorf("Expected title 'Task Database', got %s", dbTitle)
	}

	// Test property operations
	if _, exists := database.GetProperty("Status"); !exists {
		t.Error("Expected Status property to exist")
	}

	if !database.HasTitleProperty() {
		t.Error("Expected database to have title property")
	}

	titlePropName := database.GetTitlePropertyName()
	if titlePropName != "Name" {
		t.Errorf("Expected title property name 'Name', got %s", titlePropName)
	}

	// Add new property
	database.SetProperty("Priority", DatabaseProperty{
		Type: PropertyTypeSelect,
		Select: &SelectConfig{
			Options: []SelectOption{
				{Name: "High", Color: ColorRed},
				{Name: "Low", Color: ColorBlue},
			},
		},
	})

	if _, exists := database.GetProperty("Priority"); !exists {
		t.Error("Expected Priority property to be set")
	}
}

// TestUserTypes tests user type creation and validation.
func TestUserTypes(t *testing.T) {
	// Test person user
	personUser := &User{
		BaseObject: BaseObject{
			Object: ObjectTypeUser,
		},
		ID:   UserID("550e8400-e29b-41d4-a716-446655440000"),
		Type: UserTypePerson,
		Name: stringPtr("John Doe"),
		Person: &Person{
			Email: stringPtr("john@example.com"),
		},
	}

	if !personUser.IsPerson() {
		t.Error("Expected user to be a person")
	}

	if personUser.IsBot() {
		t.Error("Expected user not to be a bot")
	}

	email := personUser.GetEmail()
	if email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got %s", email)
	}

	displayName := personUser.GetDisplayName()
	if displayName != "John Doe" {
		t.Errorf("Expected display name 'John Doe', got %s", displayName)
	}

	// Test bot user
	botUser := &User{
		BaseObject: BaseObject{
			Object: ObjectTypeUser,
		},
		ID:   UserID("550e8400-e29b-41d4-a716-446655440001"),
		Type: UserTypeBot,
		Name: stringPtr("My Bot"),
		Bot: &Bot{
			Owner: &BotOwner{
				Type: BotOwnerTypeUser,
				User: personUser,
			},
			WorkspaceName: stringPtr("My Workspace"),
		},
	}

	if !botUser.IsBot() {
		t.Error("Expected user to be a bot")
	}

	if botUser.IsPerson() {
		t.Error("Expected user not to be a person")
	}

	botEmail := botUser.GetEmail()
	if botEmail != "" {
		t.Errorf("Expected empty email for bot, got %s", botEmail)
	}
}

// TestJSONMarshaling tests JSON marshaling and unmarshaling of complex types.
func TestJSONMarshaling(t *testing.T) {
	// Create a complex page structure
	dbID := DatabaseID("550e8400-e29b-41d4-a716-446655441111")
	parent := &Parent{Type: ParentTypeDatabase, DatabaseID: &dbID}

	properties := map[string]Property{
		"Name":      *NewTitleProperty("Test Page"),
		"Status":    *NewSelectProperty(stringPtr("Done"), colorPtr(ColorGreen)),
		"Completed": *NewCheckboxProperty(true),
	}

	page := NewPage(parent, properties)
	page.ID = "550e8400-e29b-41d4-a716-446655440000"
	page.CreatedTime = time.Now().UTC()
	page.LastEditedTime = time.Now().UTC()

	// Marshal to JSON
	data, err := json.Marshal(page)
	if err != nil {
		t.Fatalf("Failed to marshal page: %v", err)
	}

	// Unmarshal from JSON
	var unmarshaled Page
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal page: %v", err)
	}

	// Verify key fields
	if unmarshaled.ID != page.ID {
		t.Errorf("Expected ID %s, got %s", page.ID, unmarshaled.ID)
	}

	if unmarshaled.Object != page.Object {
		t.Errorf("Expected object %s, got %s", page.Object, unmarshaled.Object)
	}

	if len(unmarshaled.Properties) != len(page.Properties) {
		t.Errorf("Expected %d properties, got %d", len(page.Properties), len(unmarshaled.Properties))
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func colorPtr(c Color) *Color {
	return &c
}
