# go-fotmob — deprecated: do not use

A Go client for [FotMob](https://www.fotmob.com). This is the original public repository, and it
is **not the canonical one and it is not maintained here**. Canonical development, its issues and
its decisions live in the private clubtools GitLab project:
`gitlab.com/heers.it/clubtools/go-fotmob`.

**The code in this repository does not work against FotMob's API as it is today.** It still calls
the old `/api/` prefix that FotMob moved to `/api/data/`; the repair lives in the canonical
repository, and syncing it back here is best-effort — so this snapshot can stay behind
indefinitely.

Two facts a visitor should know:

- **The library is temporary.** clubtools uses it for a pilot and deletes it before GA, in the
  same release that replaces it with a licensed provider. When that happens, this repository will
  be archived and left as-is.
- **A stale upstream that looks authoritative is worse than no upstream.** That is why this
  notice exists: leaving a broken library looking current is the one outcome this repository must
  not produce.

Issues and pull requests here are not the working record.
