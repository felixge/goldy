# Contributing

Contributions are welcome, but please reach out to discuss bigger changes
before putting in too much work.

That being said, I'm very likely to accept patches to help with the following
use cases:

## Big files

Some users might want to work with very large fixtures (e.g. videos?). It might
be nice to provide them with a way to use io.Readers instead of loading
everything into memory. But there should definitley be some discussion around
the API for this beforehand. I'm not willing to make the API more annoying
for everybody just to satisfy this UC.

## Debugging CI server errors:

I could imagine a situation where a test fails on the CI server (e.g. Jenkins)
but not locally. So perhaps a new `GOLDY=verbose` mode could be implemented
that outputs the contents of the in-memory fixture file if comparison with the
on-disk fixture file fails.
