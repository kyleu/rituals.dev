package socket

// Client Messages
const (
	ClientCmdPing = "ping"

  ClientCmdConnect = "connect"
  ClientCmdUpdateSession = "update-session"

	ClientCmdGetActions = "get-actions"
	ClientCmdGetSprints = "get-sprints"
	ClientCmdSetSprint = "set-sprint"

	ClientCmdUpdateProfile = "update-profile"

  ClientCmdAddStory = "add-story"
  ClientCmdUpdateStory = "update-story"
  ClientCmdRemoveStory = "remove-story"
  ClientCmdSetStoryStatus = "set-story-status"
  ClientCmdSubmitVote = "submit-vote"

  ClientCmdAddReport = "add-report"
  ClientCmdUpdateReport = "update-report"
  ClientCmdRemoveReport = "remove-report"

  ClientCmdAddFeedback = "add-feedback"
  ClientCmdUpdateFeedback = "update-feedback"
  ClientCmdRemoveFeedback = "remove-feedback"
)

// Server Messages
const (
	ServerCmdError = "error"
  ServerCmdPong = "pong"

  ServerCmdSessionJoined = "session-joined"
  ServerCmdSessionUpdate = "session-update"
	ServerCmdSprintUpdate = "sprint-update"

	ServerCmdActions = "actions"
	ServerCmdSprints = "sprints"

  ServerCmdMemberUpdate = "member-update"
  ServerCmdOnlineUpdate = "online-update"

  ServerCmdStoryUpdate = "story-update"
  ServerCmdStoryRemove = "story-remove"
  ServerCmdStoryStatusChange = "story-status-change"
  ServerCmdVoteUpdate = "vote-update"

  ServerCmdReportUpdate = "report-update"
  ServerCmdReportRemove = "report-remove"

  ServerCmdFeedbackUpdate = "feedback-update"
  ServerCmdFeedbackRemove = "feedback-remove"
)
