# TODO
- [x] Implement basic asset copy, html generation and file operations
    + [x] Markdown -> HTML function: start with the most naive implementation. Can always reimplement with extensions, options, context, etc.
    + [x] Recursive directory copy
    + [x] Asset copy
- [ ] Implement guts of "init" command
    + [x] detect and create directory structure with default configuration
    + [x] generate directories using existing config.json
    + [ ] copy default html template
    + [ ] figure out how to save these on host machine during install, but for now, just get the directories and config made
- [ ] Early CLI implementaion
    + [x] "init" command
    + [x] "build" command
    + [x] "config" command
    + [ ] "help" command
      + [x] main help
      + [ ] command help
    + [ ] skeleton "serve"
- [ ] HTTP Server
    + [ ] only GET on root of target folder
    + [ ] wire in "serve" command

That should be enough to get started. Check back later.
