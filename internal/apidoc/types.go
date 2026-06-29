package apidoc

// Shared request/response types for Swagger generation.

type ErrorResp struct {
	Error string `json:"error" example:"invalid credentials"`
}

type User struct {
	ID             string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email          string `json:"email" example:"user@example.com"`
	Role           string `json:"role" example:"student" enums:"student,mentor,admin"`
	GithubUsername string `json:"github_username,omitempty" example:"octocat"`
	DisplayName    string `json:"display_name,omitempty" example:"Jane Doe"`
}

type RegisterReq struct {
	Email          string `json:"email" example:"user@example.com"`
	Password       string `json:"password" example:"secret123"`
	Role           string `json:"role,omitempty" example:"student" enums:"student,mentor,admin"`
	GithubUsername string `json:"github_username,omitempty" example:"octocat"`
	DisplayName    string `json:"display_name,omitempty" example:"Jane Doe"`
}

type LoginReq struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

type TokenResp struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type RefreshResp struct {
	Token string `json:"token"`
}

type Listing struct {
	ID                   string `json:"id"`
	MentorID             string `json:"mentor_id"`
	MentorDisplayName    string `json:"mentor_display_name,omitempty"`
	MentorGithubUsername string `json:"mentor_github_username,omitempty"`
	OSSProjectName       string `json:"oss_project_name" example:"kubernetes"`
	OSSRepoURL           string `json:"oss_repo_url" example:"https://github.com/kubernetes/kubernetes"`
	Description          string `json:"description"`
	PriceCents           int    `json:"price_cents" example:"9900"`
	DurationWeeks        int    `json:"duration_weeks" example:"8"`
	TotalSlots           int    `json:"total_slots" example:"5"`
	FilledSlots          int    `json:"filled_slots" example:"2"`
	Status               string `json:"status" example:"active"`
}

type CreateListingReq struct {
	OSSProjectName string `json:"oss_project_name" example:"kubernetes"`
	OSSRepoURL     string `json:"oss_repo_url" example:"https://github.com/kubernetes/kubernetes"`
	Description    string `json:"description"`
	PriceCents     int    `json:"price_cents" example:"9900"`
	DurationWeeks  int    `json:"duration_weeks" example:"8"`
	TotalSlots     int    `json:"total_slots" example:"5"`
}

type UpdateListingReq struct {
	Description   string `json:"description,omitempty"`
	Status        string `json:"status,omitempty" example:"active"`
	PriceCents    int    `json:"price_cents,omitempty"`
	DurationWeeks int    `json:"duration_weeks,omitempty"`
}

type Profile struct {
	ID             string `json:"id"`
	Email          string `json:"email,omitempty"`
	Role           string `json:"role"`
	GithubUsername string `json:"github_username,omitempty"`
	DisplayName    string `json:"display_name,omitempty"`
	Bio            string `json:"bio,omitempty"`
}

type UpdateProfileReq struct {
	DisplayName    string `json:"display_name,omitempty"`
	Bio            string `json:"bio,omitempty"`
	GithubUsername string `json:"github_username,omitempty"`
}

type Enrollment struct {
	ID        string `json:"id"`
	ListingID string `json:"listing_id"`
	Status    string `json:"status" example:"pending"`
}

type CreateEnrollmentReq struct {
	ListingID string `json:"listing_id"`
}

type ActivateEnrollmentReq struct {
	CheckoutSessionID string `json:"checkout_session_id,omitempty"`
}

type StatusResp struct {
	Status string `json:"status" example:"active"`
}

type ContributionReq struct {
	PRURL string `json:"pr_url" example:"https://github.com/org/repo/pull/42"`
}

type Contribution struct {
	ID             string `json:"id"`
	PRURL          string `json:"pr_url"`
	GithubVerified bool   `json:"github_verified"`
}

type Session struct {
	ID            string `json:"id"`
	ListingID     string `json:"listing_id"`
	ScheduledAt   string `json:"scheduled_at" example:"2026-06-28T15:00:00Z"`
	JitsiRoomName string `json:"jitsi_room_name"`
	JitsiURL      string `json:"jitsi_url"`
	Status        string `json:"status" example:"scheduled"`
}

type CreateSessionReq struct {
	ListingID   string `json:"listing_id"`
	ScheduledAt string `json:"scheduled_at" example:"2026-06-28T15:00:00Z"`
}

type UpdateSessionReq struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
	Status      string `json:"status,omitempty"`
}

type JoinSessionResp struct {
	JitsiURL string `json:"jitsi_url"`
	Room     string `json:"room"`
}

type ProgressReq struct {
	EnrollmentID string `json:"enrollment_id"`
	Note         string `json:"note,omitempty"`
	PRURL        string `json:"pr_url,omitempty"`
}

type ProgressEntry struct {
	ID           string `json:"id"`
	EnrollmentID string `json:"enrollment_id"`
	Note         string `json:"note,omitempty"`
	PRURL        string `json:"pr_url,omitempty"`
	CreatedAt    string `json:"created_at"`
}

type MentorApplyReq struct {
	GithubUsername string `json:"github_username,omitempty" example:"octocat"`
}

type MentorApplication struct {
	ID        string                 `json:"id"`
	UserID    string                 `json:"user_id"`
	Status    string                 `json:"status" example:"pending"`
	GithubData map[string]interface{} `json:"github_data,omitempty"`
	CreatedAt string                 `json:"created_at,omitempty"`
}

type ReviewApplicationReq struct {
	Status string `json:"status" enums:"approved,rejected"`
}

type CheckoutReq struct {
	ListingID    string `json:"listing_id"`
	StudentID    string `json:"student_id"`
	MentorID     string `json:"mentor_id"`
	EnrollmentID string `json:"enrollment_id"`
	AmountCents  int    `json:"amount_cents" example:"9900"`
	SuccessURL   string `json:"success_url" example:"http://localhost/checkout/success"`
	CancelURL    string `json:"cancel_url" example:"http://localhost/checkout/cancel"`
}

type CheckoutResp struct {
	CheckoutURL string `json:"checkout_url"`
	SessionID   string `json:"session_id"`
}

type PayoutSummary struct {
	TotalGrossCents        int64 `json:"total_gross_cents"`
	TotalMentorPayoutCents int64 `json:"total_mentor_payout_cents"`
	TotalPlatformFeeCents  int64 `json:"total_platform_fee_cents"`
	TransactionCount       int64 `json:"transaction_count"`
}

type LedgerEntry struct {
	ID                string `json:"id"`
	EventType         string `json:"event_type"`
	GrossCents        int    `json:"gross_cents"`
	PlatformFeeCents  int    `json:"platform_fee_cents"`
	MentorPayoutCents int    `json:"mentor_payout_cents"`
	CreatedAt         string `json:"created_at"`
}

type ConnectOnboardReq struct {
	ReturnURL  string `json:"return_url" example:"http://localhost/dashboard/mentor"`
	RefreshURL string `json:"refresh_url" example:"http://localhost/dashboard/mentor"`
}

type ConnectOnboardResp struct {
	OnboardingURL string `json:"onboarding_url"`
	AccountID     string `json:"account_id"`
}

type ConnectStatus struct {
	Connected bool   `json:"connected"`
	AccountID string `json:"account_id,omitempty"`
}

type HealthResp struct {
	Status  string `json:"status" example:"ok"`
	Service string `json:"service" example:"gateway"`
}
