package model

type StartWorkFlowResponse struct {
	Server           string `json:"server"`
	Task             string `json:"task"`
	RemainingCredits int    `json:"remaining_credits"`
}

type UploadWorkFlowResponse struct {
	ServerFilename string `json:"server_filename"`
}

type ProcessWorkFlowResponse struct {
	DownloadFilename string `json:"download_filename"`
	FileSize         int    `json:"file_size"`
	OutputFileSize   int    `json:"output_filesize"`
	OutputFileNumber int    `json:"output_filenumber"`
	OutputExtentions string `json:"output_extentions"`
	Timer            string `json:"timer"`
	Status           string `json:"status"`
}

type File struct {
	ServerFilename string  `json:"server_filename"`
	Filename       string  `json:"filename"`
	Rotate         *int    `json:"rotate,omitempty"`
	Password       *string `json:"password,omitempty"`
}

type MetaData struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Subject      string `json:"subject"`
	Keywords     string `json:"keywords"`
	Creator      string `json:"creator"`
	Producer     string `json:"producer"`
	CreationDate string `json:"creation_date"`
	ModDate      string `json:"mod_date"`
	Trapped      string `json:"trapped"`
}

type EditPdfRequest struct {
	Elements []Element `json:"elements"`
	Text     []Text    `json:"text"`
	Image    Image     `json:"image"`
	Svg      Svg       `json:"svg"`
}

type Image struct {
	ServerFilename string `json:"server_filename"`
}

type Svg struct {
	ServerFilename string `json:"server_filename"`
}

type Text struct {
	Text          string `json:"text"`
	TextAlign     string `json:"text_align"`
	FontFamily    string `json:"font_family"`
	FontSize      int    `json:"font_size"`
	FontStyle     string `json:"font_style"`
	FontColor     string `json:"font_color"`
	LetterSpacing int    `json:"letter_spacing"`
	UnderlineText bool   `json:"underline_text"`
}

type Element struct {
	Type        string      `json:"type"`
	Pages       *string     `json:"pages,omitempty"`
	Zindex      string      `json:"zindex"`
	Dimensions  *Dimension  `json:"dimensions,omitempty"`
	Coordinates *Coordinate `json:"coordinates,omitempty"`
	Rotation    *int        `json:"rotation,omitempty"`
	Opacity     *int        `json:"opacity,omitempty"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Number int `json:"number,omitempty"`
type ProcessWorkFlowRequest struct {
	// task fields
	Task              string     `json:"task"`
	Tool              string     `json:"tool"`
	Files             []File     `json:"files"`
	Metas             []MetaData `json:"metas,omitempty"`
	IgnoreErrors      *bool      `json:"ignore_errors,omitempty"`
	IgnorePassword    *bool      `json:"ignore_password,omitempty"`
	OutputFilename    *string    `json:"output_filename,omitempty"`
	PackagedFilename  *string    `json:"packaged_filename,omitempty"`
	FileEncryptionKey *string    `json:"file_encryption_key,omitempty"`
	TryPdfRepair      *bool      `json:"try_pdf_repair,omitempty"`
	TryImageRepair    *bool      `json:"try_image_repair,omitempty"`
	CustomInt         *int       `json:"custom_int,omitempty"`
	CustomString      *string    `json:"custom_string,omitempty"`
	Webhook           *string    `json:"webhook,omitempty"`

	// split fields
	SplitMode       *string `json:"split_mode,omitempty"`
	Range           *string `json:"range,omitempty"`
	FixedRange      *string `json:"fixed_range,omitempty"`
	RemovePages     *string `json:"remove_pages,omitempty"`
	RemovePasswords *bool   `json:"remove_passwords,omitempty"`

	// compress
	CompressLevel *int `json:"compress_level,omitempty"`

	// PDF OCR
	OcrLanguage *string `json:"ocr_language,omitempty"`

	// PDF to JPG
	PdfjpgMode *string `json:"pdfjpg_mode,omitempty"`

	// Image to PDF
	Orientation *string `json:"orientation,omitempty"`
	Margin      *int    `json:"margin,omitempty"`
	PageSize    *string `json:"page_size,omitempty"`

	// Image to PDF use and split use
	MeregeAfter *bool `json:"merege_after,omitempty"`

	// Page numbers
	FacingPages *bool   `json:"facing_pages,omitempty"`
	FirstCover  *bool   `json:"first_cover,omitempty"`
	Page        *string `json:"page,omitempty"`

	StartingNumber               *int    `json:"starting_number,omitempty"`
	VerticalPosition             *string `json:"vertical_position,omitempty"`
	HorizontalPosition           *string `json:"horizontal_position,omitempty"`
	VerticalPositionAdjustment   *int    `json:"vertical_position_adjustment,omitempty"`
	HorizontalPositionAdjustment *int    `json:"horizontal_position_adjustment,omitempty"`
	FontFamily                   *string `json:"font_family,omitempty"`
	FontSize                     *int    `json:"font_size,omitempty"`
	FontColor                    *string `json:"font_color,omitempty"`
	Text                         *string `json:"text,omitempty"`

	// Watermark
	Mode           *string `json:"mode,omitempty"`
	Image          *string `json:"image,omitempty"`
	Moasic         *bool   `json:"moasic,omitempty"`
	Rotation       *int    `json:"rotation,omitempty"`
	FontStyle      *string `json:"font_style,omitempty"`
	Transparent    *int    `json:"transparent,omitempty"`
	Layer          *string `json:"layer,omitempty"`
	Password       *string `json:"password,omitempty"`
	Conformance    *string `json:"conformance,omitempty"`
	AllowDowngrade *bool   `json:"allow_downgrade,omitempty"`
	Detailed       *bool   `json:"detailed,omitempty"`
}
