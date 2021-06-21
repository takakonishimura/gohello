# Small starter programs

## Program Arguments

```
$(GOPATH)/src/github.com/takakonishimura/gohello/small >>  go build args.go
$(GOPATH)/src/github.com/takakonishimura/gohello/small >>  ./args tay pusheen kurumi
tay & pusheen & kurumi
```

## Standard Input

```
$(GOPATH)/src/github.com/takakonishimura/gohello/small >> go run stdin.go
hi
hi
meow
tay
pusheen
tay
```
hit `<ctrl-D>` to exit

```
2       hi
1	meow
2	tay
1	pusheen
```

## File Input

No files
```
$(GOPATH)/src/github.com/takakonishimura/gohello/small >> go run filein.go             
2021/06/21 16:58:33 no files provided
exit status 1
```

One file
```
$(GOPATH)/src/github.com/takakonishimura/gohello/small >>  go run filein.go tkamb.txt   
2	People
2	Maycomb
3	at
2	of
14	the
2	into
4	in
2	A
2	people’s
2	said
5	was
2	were
2	to
5	he
2	their
2	had
6	Radley
5	and
4	a
2	but
2	by
3	would
```

Multiple files
```
go run filein.go tkamb.txt ge.txt
6	you
6	his
2	by
5	was
13	a
2	People
2	boy
2	were
2	bring
4	your
...
2	never
11	to
2	dare
2	seen
3	it
2	person
3	would
3	A
2	him
2	That
2	into
4	You
2	way
5	he
4	from
2	creep
2	am
3	or
```

## Fetch

```
$(GOPATH)/src/github.com/takakonishimura/gohello/small >> go run fetch.go https://www.gopl.io/ch1/fetch
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
	  "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io"></meta>
  <title>The Go Programming Language</title>
  <script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-68781809-1', 'auto');
  ga('send', 'pageview');
  </script>
  <link rel="stylesheet" type="text/css" href="style.css" />
  <style>
.bio {
	font-size:	95%;
	text-align:	justify;
}
.biblio {
	margin:		0.5em;
	font-size:	110%;
	line-height:	130%;
}
.center {
	text-align:	center;
}
td {
	padding:	0.5em 1em 0.5em 1em;
	vertical-align: top;
}
#toc h1 {
	font-size:	95%;
	font-weight:    normal;
}
hr {
	border:		0;
	border-top:	thin solid #ccc
}
  </style>
</head>
<body>
<table><tr>
<td>
  <a href='http://www.informit.com/store/go-programming-language-9780134190440'><img border='1' src="cover.png" height='600'/></a>
  <br/>
  <div style='text-align: center'>
  <a href='http://www.amazon.com/dp/0134190440'><img border='0' width='150' src="buyfromamazon.png"/></a>
  &nbsp;&nbsp;&nbsp;
  <a href='http://www.informit.com/store/go-programming-language-9780134190440'/><img border='0' height='26' src="informit.png"/></a>
  &nbsp;&nbsp;&nbsp;
  <a href='http://www.barnesandnoble.com/w/1121601944'><img border='0' width='150' src="barnesnoble.png"/></a>
  </div>
</br>
</td>
<td width='500'>
<h1 class='center'>The Go Programming Language</h1>
<p class='biblio center'>
Alan A. A. Donovan &middot; Brian W. Kernighan<br/>
Published Oct 26, 2015 in paperback and Nov 20 in e-book<br/>
Addison-Wesley; 380pp; ISBN:&nbsp;978-0134190440<br/>
<tt>authors</tt><tt>&#0064;</tt><tt>gopl.io</tt>
</p>

<!-- Table of Contents -->
<div id='toc'>
<table><tr>
<td>
<h1>Contents [<a href='ch1.pdf'/>PDF</a>]</h1>
<h1>Preface  [<a href='ch1.pdf'/>PDF</a>]</h1>
<h1>1. Tutorial [<a href='ch1.pdf'/>PDF</a>]</h1>
<h1>2. Program Structure</h1>
<h1>3. Basic Data Types</h1>
<h1>4. Composite Types</h1>
<h1>5. Functions</h1>
<h1>6. Methods</h1>
</td>
<td>
<h1>7. Interfaces</h1>
<h1>8. Goroutines and Channels</h1>
<h1>9. Concurrency with Shared Variables</h1>
<h1>10. Packages and the Go Tool</h1>
<h1>11. Testing</h1>
<h1>12. Reflection</h1>
<h1>13. Low-Level Programming</h1>
<h1>Index  [<a href='ch1.pdf'/>PDF</a>]</h1>
</td>
</tr>
<tr>
<td colspan='2'>
<h1>
  <a href='https://github.com/adonovan/gopl.io/'>Source code</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href='reviews.html'>Reviews</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href='translations.html'>Translations</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href='errata.html'>Errata</a>
</h1>
</td>
</tr>
</table>
</div>

<p class="bio">
Alan Donovan is a Staff Engineer in Google's infrastructure division,
specializing in software development tools.  Since 2012, he has been
working on the Go team, designing libraries and tools for static
analysis.  He is the author of the
<a href='http://golang.org/s/oracle-user-manual'><code>oracle</code></a>,
<a href='http://golang.org/lib/godoc/analysis/help.html'><code>godoc -analysis</code></a>,
<a href='https://github.com/golang/tools/blob/master/refactor/eg/eg.go'><code>eg</code></a>,
and <a href='https://github.com/golang/tools/blob/master/refactor/rename/rename.go'><code>gorename</code></a> tools.
</p> 

<p class="bio">
Brian Kernighan was in the Computing Science Research center at Bell
Labs until 2000, where he worked on languages and tools for Unix.  He
is now a professor in the Computer Science Department at Princeton.
He is the co-author of several books, including
<a href='http://www.amazon.com/dp/0131103628?tracking_id=disfordig-20'>The C Programming Language</a> and
<a href='http://www.amazon.com/dp/020161586X?tracking_id=disfordig-20'>The Practice of Programming</a>.

</td>
</tr></table>
</body>
</html>
```

## Fetch Concurrently

```
go run fetchall.go http://www.gopl.io https://www.golang.org https://godoc.org
0.03s    4154 http://www.gopl.io
0.38s    9951 https://www.golang.org
0.58s   12642 https://godoc.org
0.58s elapsed
```
