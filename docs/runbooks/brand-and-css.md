# Brand tokens and the CSS build

nwc owns the look of every NimsForest web UI: the shared layout, the Tailwind
palette, and the self-hosted brand fonts. This runbook covers changing any of
them without unstyling the apps downstream.

## Rebuilding style.css

`static/style.css` is a committed build artifact, not source. After editing
`static/input.css`, `tailwind.config.js` or any template:

```bash
make css        # npx tailwindcss -i static/input.css -o static/style.css --minify
```

**The build scans sibling repos.** `tailwind.config.js` includes
`../**/templates/**/*.html` in its `content` globs, so the utilities that end up
in `style.css` depend on which NimsForest repos are checked out next to nwc at
build time. Build on a machine with the full workspace present, or classes used
only by an app you do not have locally will be missing from the output. Check a
suspicious build with `grep -c 'bg-paper' static/style.css` before committing.

## Palette

Two groups of colour names live in `tailwind.config.js`:

- **Brand tokens** (`paper`, `band`, `card`, `ink`, `body-text`, `mono-label`,
  `forest`, `gold`, `dark-ground`, `river`, `leaf`, `wind`, `hairline` ...) are
  the vocabulary of the rebrand spec. New work uses these.
- **Legacy tokens** (`forest-bg`, `forest-card`, `forest-green`, `canopy`,
  `solar-gold`, `soil`, `sprout` ...) predate the rebrand and are still used by
  roughly 120 template files across the apps that import nwc. They are kept and
  repointed at brand values, so bumping nwc **restyles** those apps instead of
  leaving their markup unstyled. Do not delete a legacy token without first
  removing every use of it downstream.

`mono-label` `#6B7C6D` is the lightest grey that clears 4.5:1 on the paper and
band backgrounds. Do not lighten it.

## Fonts

Instrument Serif (display), Figtree (body/UI) and IBM Plex Mono (labels) are
**self-hosted** in `static/fonts/` as latin + latin-ext woff2 subsets, declared
by `@font-face` at the top of `static/input.css`. Nothing hotlinks Google Fonts:
the UI keeps rendering if fonts.googleapis.com is unreachable, and no visitor
detail leaks to a third party on page load.

To add a weight, fetch the css2 stylesheet with a modern browser User-Agent (an
older UA gets you TTF instead of woff2), download the woff2 files it references
into `static/fonts/`, add the matching `@font-face` block with the same
`unicode-range`, then `make css`.

Emoji are typed as Unicode characters and drawn by the viewer's platform font.
`static/fonts/noto-color-emoji-brand.woff2` (7 KB) sits at the **end** of the
body font stack as a last resort, so a visitor on a platform with no emoji font
sees the brand rather than empty boxes. Everywhere a platform emoji font exists
it wins, because those families are listed ahead of it.

To add a glyph to that subset, re-request it with every glyph the brand uses,
not just the new one, or the others drop out:

```bash
curl -sG -A "<modern browser UA>" \
  --data-urlencode "family=Noto Color Emoji" \
  --data-urlencode "text=🌳🍃✨💧🌊🍂🏡🙂🤖🌱" \
  https://fonts.googleapis.com/css2
```

Then download the woff2 it points at, replace the file, and widen the
`unicode-range` in `input.css` to match.

The SVG artwork in the brand kit is for print and fixed artwork; the one
exception is `static/nimsforest-mark.svg`, the favicon, which is referenced by
the layout. Its artwork derives from Noto Emoji: keep
`static/NOTO-EMOJI-LICENSE.txt` beside it (Apache 2.0, (c) Google LLC).

## Layout configuration

`AppConfig` fields that affect chrome:

| Field | Effect |
|-------|--------|
| `Name` | App label beside the wordmark. Empty or `"forest"` shows "Nimsforest" alone |
| `Emoji` | Header glyph, defaults to 🌳 |
| `NavItems` | Header links. `Primary: true` renders one as a filled pill ("Get started") |
| `CreditStrip` | Mono strip above the header. Empty means no strip |
| `FooterMarks` | Show the 🌳 🍃 ✨ row in the footer |
| `FooterNav` | Repeat NavItems as a footer link row |
| `FullBleed` | Let pages own their width, for full-width bands. Default keeps the centred container that app pages expect |
| `Stylesheets` | Extra stylesheet paths, linked after the shared one. For page-level CSS that does not belong in the shared Tailwind build |

`FullBleed` is the one to watch: app page templates assume the layout supplies
the `max-w-7xl` container and padding. Turning it on without giving each page
its own container leaves content edge to edge.

## Releasing

Consumers pin nwc by version in their `go.mod` (they are not all on the same
one), so a change here reaches each app only when that app bumps. Tag, then bump
the apps you intend to restyle:

```bash
git tag v0.10.0 && git push origin v0.10.0
cd ../<app> && go get github.com/nimsforest/nimsforestwebcomponents@v0.10.0
```
