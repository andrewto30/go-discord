package models

// Embed is a struct representing Discord embed object
type Embed struct {
	Username  string            `json:"username"`
	AvatarURL string            `json:"avatar_url"`
	Content   string            `json:"content"`
	Embeds    []EmbedComponents `json:"embeds"`
}

// EmbedComponents is a struct representing Discord embed object components
type EmbedComponents struct {
}
