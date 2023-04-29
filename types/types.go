package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/shlex"
	"github.com/mitchellh/mapstructure"
)

type PageType string

const (
	DetailPage PageType = "detail"
	ListPage   PageType = "list"
	FormPage   PageType = "form"
)

type Page struct {
	Type    PageType `json:"type"`
	Title   string   `json:"title,omitempty"`
	Actions []Action `json:"actions,omitempty"`

	// Form page
	SubmitAction *Action `json:"submitAction,omitempty"`

	// Detail page
	Preview *Preview `json:"preview,omitempty"`

	// List page
	ShowPreview bool `json:"showPreview,omitempty"`
	EmptyView   *struct {
		Text    string   `json:"text,omitempty"`
		Actions []Action `json:"actions,omitempty"`
	} `json:"emptyView,omitempty"`
	Items []ListItem `json:"items,omitempty"`
}

type Preview struct {
	HighLight string   `json:"highlight,omitempty"`
	Text      string   `json:"text,omitempty"`
	Command   *Command `json:"command,omitempty"`
}

func (p *Preview) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw := raw.(type) {
	case string:
		p.Text = raw
	case map[string]interface{}:
		if err := mapstructure.Decode(raw, p); err != nil {
			return err
		}
	default:
		return errors.New("invalid preview")
	}

	return nil
}

type ListItem struct {
	Id          string   `json:"id,omitempty"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle,omitempty"`
	Preview     *Preview `json:"preview,omitempty"`
	Accessories []string `json:"accessories,omitempty"`
	Actions     []Action `json:"actions,omitempty"`
}

type FormInputType string

const (
	TextFieldInput FormInputType = "textfield"
	TextAreaInput  FormInputType = "textarea"
	DropDownInput  FormInputType = "dropdown"
	CheckboxInput  FormInputType = "checkbox"
)

type DropDownItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Input struct {
	Name        string        `json:"name"`
	Type        FormInputType `json:"type"`
	Title       string        `json:"title"`
	Placeholder string        `json:"placeholder,omitempty"`
	Default     any           `json:"default,omitempty"`

	// Only for dropdown
	Items []DropDownItem `json:"items,omitempty"`

	// Only for checkbox
	Label             string `json:"label,omitempty"`
	TrueSubstitution  string `json:"trueSubstitution,omitempty"`
	FalseSubstitution string `json:"falseSubstitution,omitempty"`
}

type ActionType string

const (
	CopyAction   = "copy"
	OpenAction   = "open"
	PushAction   = "push"
	RunAction    = "run"
	ExitAction   = "exit"
	ReloadAction = "reload"
)

type OnSuccessType string

var (
	OpenOnSuccess   OnSuccessType = "open"
	PushOnSuccess   OnSuccessType = "push"
	ExitOnSuccess   OnSuccessType = "exit"
	ReloadOnSuccess OnSuccessType = "reload"
	CopyOnSuccess   OnSuccessType = "copy"
)

type Action struct {
	Title  string     `json:"title,omitempty"`
	Type   ActionType `json:"type"`
	Key    string     `json:"key,omitempty"`
	Inputs []Input    `json:"inputs,omitempty"`

	// copy
	Text string `json:"text,omitempty"`

	// open
	Target string `json:"target,omitempty"`

	// push
	Page string `json:"page,omitempty"`

	// run
	Command   *Command      `json:"command,omitempty"`
	Confirm   bool          `json:"confirm,omitempty"`
	OnSuccess OnSuccessType `json:"onSuccess,omitempty"`
}

type Command struct {
	Args  []string `json:"args"`
	Input string   `json:"input,omitempty"`
	Dir   string   `json:"dir,omitempty"`
}

func (c Command) Cmd() *exec.Cmd {
	cmd := exec.Command(c.Args[0], c.Args[1:]...)
	cmd.Dir = c.Dir
	cmd.Stdin = strings.NewReader(c.Input)

	return cmd

}

func (c Command) Run() error {
	err := c.Cmd().Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return fmt.Errorf("command exited with %d: %s", exitErr.ExitCode(), string(exitErr.Stderr))
	} else if err != nil {
		return err
	}

	return nil

}

func (c Command) Output() ([]byte, error) {
	output, err := c.Cmd().Output()

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return nil, fmt.Errorf("command exited with %d: %s", exitErr.ExitCode(), string(exitErr.Stderr))
	} else if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Command) UnmarshalJSON(data []byte) error {
	var sa []string
	if err := json.Unmarshal(data, &sa); err == nil {
		if len(sa) == 0 {
			return fmt.Errorf("empty command")
		}
		c.Args = sa
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		args, err := shlex.Split(s)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return fmt.Errorf("empty command")
		}

		c.Args = args
		return nil
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err == nil {
		if err := mapstructure.Decode(m, c); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("invalid command")
}
