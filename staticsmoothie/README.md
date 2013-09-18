Staticsmoothie
==============

> There are only two hard things in Computer Science: 
> cache invalidation and naming things.
> 
> -- Phil Karlton 

`Staticsmootie` is an app to serve very simple "websites". It is a Go programm which parses the incoming request URL, checks whether a coresponding path on the filesystem exists and renders a markdown + amber layout right to your eyes.

If you want to run this example you will have to follow the folllowing structure:

| Files             | Type           | Used to                        |
| ----------------- |:--------------:| ------------------------------:|
| staticsmoothie.go | Go programm    | Run the "server"               |
| .md               | index markdown | markdown, contains content     |
| error.md          | error markdown | markdown, contains content     |
| layout.amber      | layout.amber   | providing the basic layout     |
| config.json       | a little conf  | setting title and footer       |
| pages/            | sub-pages dir  | all sub pages md files go here |

A detailed installation example will follow plus what is necessary for cross-compilation!