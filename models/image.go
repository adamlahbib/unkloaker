package models

type Image struct {
	Name              string
	Tag               string `default:"latest"`
	RepoPath          string
	DockerfileContext string `default:"."`
}

type ImageRequest struct {
	Image Image
}

type ImageResponse struct {
	ImageNameWithTag string
	Success          bool
}
