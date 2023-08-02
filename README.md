fortune(1) - sample lines from a file

Fortune prints an aphorism chosen at random.

# Usage

```sh
go install bwsd.dev/fortunes@latest
wget -O "$HOME/fortunes" https://raw.githubusercontent.com/bwsd0/fortune/master/lib/fortunes
echo 'export FORTUNES="$HOME/fortunes"' >> ~/.bashrc
```

# Sources

[Jar of fortunes](http://fortunes.cat-v.org/cat-v/)

- [9 front](http://fortunes.cat-v.org/9front/)
- [cat-v](http://fortunes.cat-v.org/cat-v/)
- [FreeBSD](http://fortunes.cat-v.org/freebsd/)
- [FreeBSD (offensive)](http://fortunes.cat-v.org/freebsd/offensive)
- [Kernelnewbies](http://fortunes.cat-v.org/kernelnewbies/)
- [OpenBSD](http://fortunes.cat-v.org/openbsd/)
- [Plan 9](http://fortunes.cat-v.org/plan_9/)
- [Terry A. Davis](https://gist.githubusercontent.com/bwasd/62ceccc10d88d3afc0e0f449001c87c1/raw/9c0408927448cf33b41515840dcc53449d8a2b6b/terry)

To update the fortune file:

```sh
wget -O lib/fortunes https://9fans.github.io/usr/local/plan9/lib/fortunes
```

# History

`fortune(6)` first appeared in the Seventh Edition of Unix in 1979, written by
Ken Arnold.

- [Unix 7th Ed.](http://man.cat-v.org/unix_7th/6/ching)
- [Unix 8th Ed.](http://man.cat-v.org/unix_8th/6/fortune)
- [Plan 9 2nd Ed](http://man.cat-v.org/plan_9_2nd_ed/1/)
- [FreeBSD 1.1](https://svnweb.freebsd.org/base?view=revision&revision=325828)
- [9 Front (July, 2011)](http://man.cat-v.org/9front/1/)[^1]
- [FreeBSD (2017) Removal](https://svnweb.freebsd.org/base?view=revision&revision=325828)[^2]
- [Debian ?](#)

# Notes

[1]: 9Front introduced variants [theo(1)](), [troll(1)]() and [bullshit(1)]()
[2]: [Revision 325828](https://svnweb.freebsd.org/base?view=revision&revision=325828)
> Humour is a funny thing. What is funny to one person is not funny to all
> people. What is insightful to one person is similarly not universal. The
> fortune datfiles have been around a long time and have undoubtedly amused
> people but it's time to acknowledge their subjective, and in some cases at
> least potentially offensive, nature and stop distributing them with the
> imprimatur of the FreeBSD project.

# See also
- [Aphorisms about aphorisms](https://en.wikiquote.org/wiki/Aphorisms)
- [QOTD (Quote of the Day Protocol)](https://tools.ietf.org/html/rfc865) (RFC 865)

# License

MIT
