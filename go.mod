module movieBlog/cli

go 1.24.5

require movieBlog/cli/cliParser v0.0.0-00010101000000-000000000000

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	golang.org/x/sys v0.0.0-20181122145206-62eef0e2fa9b // indirect
)

replace movieBlog/cli/cliParser => ./parser_cli
