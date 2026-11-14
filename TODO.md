# TODO
- [x] Implement basic asset copy, html generation and file operations
    + [x] Markdown -> HTML function: start with the most naive implementation. Can always reimplement with extensions, options, context, etc.
    + [x] Recursive directory copy
    + [x] Asset copy
- [x] Implement guts of "init" command
    + [x] detect and create directory structure with default configuration
    + [x] generate directories using existing config.json
    + [x] copy default html template
    + [x] figure out how to save these on host machine during install, but for now, just get the directories and config made
- [x] Early CLI implementaion
    + [x] "init" command
    + [x] "build" command
    + [x] "config" command
    + [x] "help" command
      + [x] main help
      + [x] command help
    + [x] skeleton "serve"
    + [x] "theme" command
- [x] HTTP Server
    + [x] only GET on root of target folder - just doing file server to start
    + [x] wire in "serve" command
    + **ABANDONED** implement "detatch" flag - not advised for futur production use
- [x] Refine build functionality
    + [x] test and probably fix hyperlink functionality. May need to do some path replacement at build
        * used crude strings.CutPrefix implementation.
        * Should probably handle it at the file server level in the mux.
    + [x] implement extract title
    + [x] Get the title in to the directory tree and refine "printTree" function. 
    + [x] generate an index.html in target directory (will need to learn a little html)
      + need to record and store documents and paths generated during the build step.
      + This will be a shared memory resource that will need mutex protection.
      + Use this record to generate hyperlinks in index.html
    + [x] Make sidebar deterministic
- [x] Layout, CSS, and theming
    + Will need to learn enough CSS to make a few shitty themes
    + [x] Start with a generic one so there's something to test with
    + [x] Add title bar and side bar to default and index templates

That should be enough to get started. Check back later.
