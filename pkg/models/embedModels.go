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
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Color       int       `json:"color"`
	Footer      Footer    `json:"footer"`
	Image       Image     `json:"image"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Author      Author    `json:"author"`
	Field       Field     `json:"fields"`
}

type Footer struct {
}

type Image struct {
}

type Thumbnail struct {
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Field struct {
}
