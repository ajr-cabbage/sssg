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
- [x] Early CLI implementaion
    + [x] "init" command
    + [x] "build" command
    + [x] "config" command
    + [x] "help" command
      + [x] main help
      + [x] command help
    + [x] skeleton "serve"
- [ ] HTTP Server
    + [x] only GET on root of target folder - just doing file server to start
    + [x] wire in "serve" command
    + [ ] implement "detatch" flag
      + want to be able to detatch using command flag or with keystroke if flag is forgotten
      + do it natively in the application, so this can be cross platform
- [ ] Refine build functionality
    + [ ] test and probably fix hyperlink functionality. May need to do some path replacement at build
    + [ ] generate an index.html in target directory (will need to learn a little html)
      + need to record and store documents and paths generated during the build step.
      + This will be a shared memory resource that will need mutex protection.
      + Use this record to generate hyperlinks in index.html
    + Brainstorm incremental build problem.
      + Specify directory or file?
      + Keep a database describing the state of the content directory and only update the modified parts?
- [ ] Theming and CSS
    + Will need to learn enough CSS to make a few shitty themes
    + Start with a generic one so there's something to test with

That should be enough to get started. Check back later.
