package types

// FileUploadStatus represents the status of a file upload.
//
// See: https://developers.notion.com/reference/file-upload
type FileUploadStatus string

const (
	// FileUploadStatusPending represents a file upload that is pending.
	//
	// See: https://developers.notion.com/reference/file-upload
	FileUploadStatusPending FileUploadStatus = "pending"

	// FileUploadStatusUploaded represents a file upload that has been uploaded.
	//
	// See: https://developers.notion.com/reference/file-upload
	FileUploadStatusUploaded FileUploadStatus = "uploaded"

	// FileUploadStatusExpired represents a file upload that has expired.
	//
	// See: https://developers.notion.com/reference/file-upload
	FileUploadStatusExpired FileUploadStatus = "expired"

	// FileUploadStatusFailed represents a file upload that has failed.
	//
	// See: https://developers.notion.com/reference/file-upload
	FileUploadStatusFailed FileUploadStatus = "failed"
)

// File represents a Notion file object which is one of:
// - `file_upload`: File uploaded via the File Upload API.
// - `external`: File linked from a public URL.
// - `file`: Notion-hosted file (uploaded via UI).
//
// See: https://developers.notion.com/reference/file-object
type File struct {
	Type       string                     `json:"type"`
	File       *NotionHostedFileType      `json:"file,omitempty"`
	FileUpload *NotionAPIUploadedFileType `json:"file_upload,omitempty"`
	External   *ExternalFileType          `json:"external,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Caption    []RichText                 `json:"caption,omitempty"`
}

// NotionHostedFileType represents a file manually uploaded in Notion UI.
type NotionHostedFileType struct {
	URL        string `json:"url"`
	ExpiryTime string `json:"expiry_time"`
}

// NotionAPIUploadedFileType represents a file uploaded via the File Upload API.
type NotionAPIUploadedFileType struct {
	ID string `json:"id"`
}

// ExternalFileType represents a file linked from a public URL.
type ExternalFileType struct {
	URL string `json:"url"`
}

// CreateFileUploadRequest represents a request to create a file upload.
//
// Note: This is Step 1 for uploading files.
//
// See: https://developers.notion.com/reference/create-a-file-upload
type CreateFileUploadRequest struct {
	Filename string `json:"filename" validate:"required,max=900bytes"`
	Size     int64  `json:"size" validate:"required,min=1,max=20mb"`
	Mode     string `json:"mode,omitempty" validate:"required,equals=single_part"`
}

// SendFileUploadSinglePartRequest represents a single-part file upload request.
//
// Note: This is Step 2.1 for uploading files < 20MB.
//
// See: https://developers.notion.com/docs/uploading-small-files
//
// See: https://developers.notion.com/reference/send-a-file-upload
type SendFileUploadSinglePartRequest struct {
	FileUploadID string `json:"-"` // URL path parameter
	PartNumber   *int   `json:"part_number,omitempty" validate:"required,min=1,max=1000"`
	// File data is sent as multipart/form-data under "file" key
}

// SendFileUploadMultiPartRequest represents a part in a multi-part file upload.
//
// Note: This is Step 2.2+ for uploading files > 20MB sent as multiple parts via multiple requests.
//
// See: https://developers.notion.com/docs/sending-larger-files
type SendFileUploadMultiPartRequest struct {
	PartNumber int    `json:"part_number" validate:"required,min=1,max=1000"`
	UploadURL  string `json:"upload_url" validate:"required,url"`
}

// SendExternalFileUploadRequest represents a request to upload a file from an external URL.
//
// Note: This is Step 2.3 for uploading files from an external URL.
//
// See: https://developers.notion.com/docs/importing-external-files#step-2-wait-for-the-import-to-complete
type SendExternalFileUploadRequest struct {
	Mode        string `json:"mode" validate:"required,equals=external_url"`
	ExternalURL string `json:"external_url" validate:"required,url"`
	Filename    string `json:"filename" validate:"required,max=900bytes"`
}

// ExternalFileUploadResponse represents the result of uploading a file from an external URL.
//
// Note: This is Step 2.3.1 for uploading files from an external URL.
//
// See: https://developers.notion.com/docs/importing-external-files#step-2-wait-for-the-import-to-complete
type ExternalFileUploadResponse struct {
	Status   string  `json:"status"`
	FileID   *string `json:"file_id,omitempty"`
	ErrorMsg *string `json:"error_msg,omitempty"`
}

// CompleteFileUploadRequest represents a request to complete a file upload.
//
// Note: This is Step 2.4 for uploading files.
//
// See: https://developers.notion.com/reference/complete-a-file-upload
type CompleteFileUploadRequest struct {
	FileUploadID string `json:"-"` // URL path parameter
}

// ListFileUploadsRequest represents a request to list file uploads.
//
// See: https://developers.notion.com/reference/list-file-uploads
type ListFileUploadsRequest struct {
	StartCursor *string `json:"start_cursor,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
}

// FileUploadResponse represents the response for both creating a file upload
// and retrieving file uploads.
//
// Note: This is for both single-part and multi-part file uploads.
//
// See: https://developers.notion.com/reference/retrieve-a-file-upload
// See: https://developers.notion.com/reference/file-upload
type FileUploadResponse struct {
	Object           string                      `json:"object"`
	ID               string                      `json:"id"`
	CreatedTime      string                      `json:"created_time"`
	LastEditedTime   string                      `json:"last_edited_time"`
	ExpiryTime       *string                     `json:"expiry_time,omitempty"`
	Status           FileUploadStatus            `json:"status"`
	Filename         *string                     `json:"filename,omitempty"`
	ContentType      *string                     `json:"content_type,omitempty"`
	ContentLength    *int64                      `json:"content_length,omitempty"`
	UploadURL        *string                     `json:"upload_url,omitempty"`
	CompleteURL      *string                     `json:"complete_url,omitempty"`
	FileImportResult *ExternalFileUploadResponse `json:"file_import_result,omitempty"`
}

// ListFileUploadsResponse represents the response from listing file uploads.
//
// See: https://developers.notion.com/reference/list-file-uploads
type ListFileUploadsResponse struct {
	Object     string               `json:"object"`
	Results    []FileUploadResponse `json:"results"`
	NextCursor *string              `json:"next_cursor,omitempty"`
	HasMore    bool                 `json:"has_more"`
}
