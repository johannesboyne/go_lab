My first go project
===================

The last couple of days I did some Go coding. Not much, but I always looked towards doing some Go stuff - so I did.

### TLTFA: (Too Long To F*** Around)

Do I prefer it over Node.js now? C'mon!!! Stop trying to fight one technology against another. It is not worth the effort. I encourage both technologies, for different aspects and use-cases.

### What's cool about GO!

Go has some really great features. IMHO the best and useful feature is: cross-compilation! Building platform independent binaries is just great.

E.g. setting up a simple file server:

<pre class="prettyprint lang-go">
http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc")))
</pre>
and have it compiled by doing something like
<pre class="prettyprint lang-bash">
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o appname.linux appname.go
$ GOOS=windows GOARCH=386 go build -o appname.linux appname.go

</pre>

Great, isn't it?

### What I did with GO?

The Website you are browsing is served by a Go process. It is a 84 lines long Go program and it works really well. As mentioned, thanks to cross compilation, it can be run by **any** OS:

<pre class="prettyprint lang-go">
$ ./staticsmoothie
</pre>

So, what is `staticsmoothie` and how does this website work?

---

## Staticsmoothie

> There are only two hard things in Computer Science: 
> cache invalidation and naming things.
> 
> -- Phil Karlton 

`Staticsmootie` is the app behind this website. As mentioned it is a Go programm which parses the incoming request URL, checks whether a coresponding path on the filesystem exists and renders a markdown + amber layout right to your eyes.