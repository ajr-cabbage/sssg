# sssg
The Small, Shitty, (static) Site Generator

## Motivation
There are a bunch of wonderful site generators out there, but I want one that makes writing and publishing the focus. No fluff, no real features, just a way to take your thoughts and push it to your very own website. It was the dream of the internet when I was young, and it would be nice to capture that creativity without any crap coming with it. Those sites I made as a kid were crap. This site generator is also crap, and I love it dearly. No vibes, just handwritten code.

## Quick Start
- Requirements:
    + go 1.26
    + goldmark - ```go get github.com/yuin/goldmark```
- Clone the repo and ```go build .```
- See 'Configuration' for first time setup

## Usage
sssg <command> -<options> <parameters>
- ```sssg build```
    + no flags - build all markdown in the content directory
    + -f <filename.md> - build only the specified markdown file
- ```sssg config```
    + -set <key> <value> - modify config.json through via cli
- ```sssg serve```
    + no flags - serve your target directory over http on the port specified in config.json
    + -p <port-num> - override the default port
    + -d - start server deteched from current shell session
- ```ssssg theme```
    + <theme-name.css> - change themes live - no server restart requried

## Configuration
config.json - modify directly or edit single values via cli
- "content" - the source directory for your markdown files
- "target" - the destination directory for generated html
- "assets" - directory for html templates, css themes, images, etc.
- "theme" - css filename to use
- "port" - default http port

## Contributing
Glad to take contributions. Please fork and sumbit a pull request on the main branch. Please respect the initial goal of the project. No vibes please. Code with your fingers like a caveman. See TODO.md for items in the works. Mark TODO items in-progress with your handle to communicate what you're working on.
