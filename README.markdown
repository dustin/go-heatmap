# Heatmaps

This is a toolkit for creating [heatmaps][heatmap].  Heatmaps are
awesome.  I use them for a few things, and now you can, too.

There are quite a few things you can do here and they don't all have
the high level documentation they deserve currently.  You can start by
looking at the [example][example] to get a feel for what you can do.
Basically, feed in a bunch of `x,y` type data and get back a heatmap.
Awesome.

Here's an example from a real-live report:

![Northeast US](https://raw.github.com/wiki/dustin/go-heatmap/ne-us.jpg)

## Colors

Colors are always the hardest part, so I attempted to make it easy by
providing three things:

### Predefined Color Schemes

[gheat][gheat] has a set of colors that are available by default under
the `schemes` subpackage.  You can preview them in
[their documentation][ghschemes] and use them directly.

### Drawing Color Schemes

The `schemes` subpackage also has a tool for generating a color scheme
from an image.  Given an image that looks like this:

![Alpha Fire Scheme](https://raw.github.com/dustin/go-heatmap/master/schemes/alphafire.png)

You can use `schemes.FromImage("/path/to/file")` to load it and use it
directly.  Colors near the bottom are for sparsely populated areas and
colors near the top are for areas that are densely populated.

### Computing Color Schemes

You can also compute a color scheme from spec using `schemes.Build`.
This lets you specify starting and ending colors across multiple
segments.  Fun away!

[heatmap]: http://en.wikipedia.org/wiki/Heat_map
[example]: examples/example/example.go
[gheat]: http://www.zetadev.com/software/gheat/0.2/__/doc/html/configuration.html
[ghschemes]: http://www.zetadev.com/software/gheat/0.2/__/doc/html/configuration.html#SECTION003200000000000000000
