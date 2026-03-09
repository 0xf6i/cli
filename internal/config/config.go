package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Animation AnimationConfig `yaml:"animation"`
	Style     StyleConfig     `yaml:"style"`
	Content   ContentConfig   `yaml:"content"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type AnimationConfig struct {
	FramerateMs   int    `yaml:"framerate_ms"`
	TypingSpeedMs int    `yaml:"typing_speed_ms"`
	FramesDir     string `yaml:"frames_dir"`
}

type StyleConfig struct {
	HomeColor  string `yaml:"home_color"`
	AboutColor string `yaml:"about_color"`
	ErrorColor string `yaml:"error_color"`
}

type ContentConfig struct {
	Placeholder string `yaml:"placeholder"`
	AboutText   string `yaml:"about_text"`
}

func Load(path string) (Config, error) {
	cfg := defaults()

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("reading config: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parsing config: %w", err)
	}

	return cfg, nil
}

func defaults() Config {
	return Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: "2222",
		},
		Animation: AnimationConfig{
			FramerateMs:   100,
			TypingSpeedMs: 35,
			FramesDir:     "assets/frames",
		},
		Style: StyleConfig{
			HomeColor:  "#FFFFFF",
			AboutColor: "15",
			ErrorColor: "9",
		},
		Content: ContentConfig{
			Placeholder: "Try /home, /about, or whatever you feel like big dog.",
			AboutText:   "Hey, I'm Ludwig \u2014 a developer who builds things for the terminal and the web.\n\nI love Go, distributed systems, and developer tooling. When I'm not writing code,\nyou'll find me exploring new tech, reading about systems design, or tinkering with\nside projects like this SSH portfolio you're looking at right now.\n\nType /home to go back, or /quit to disconnect.",
		},
	}
}
