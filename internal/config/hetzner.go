package config

import "os"

func HetznerApiToken() (string, error) {
	value := os.Getenv("HCLOUD_TOKEN")

	return value, nil
}