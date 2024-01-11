package utils

import "github.com/manifoldco/promptui"

func GetDefaultTemplate() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "\U0001F336 {{ . | red | cyan }}",
	}
}
