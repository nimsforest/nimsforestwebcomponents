/** @type {import('tailwindcss').Config} */

// Palette for the emoji rebrand (issue 411). Two groups of names live here:
//
//   Brand tokens  - the vocabulary the rebrand spec uses (paper, ink, forest,
//                   gold ...). New work should reach for these.
//   Legacy tokens - forest-bg, canopy, solar-gold, soil, sprout and friends.
//                   Around 120 template files across the apps that import nwc
//                   still use them, so they stay, repointed at brand values.
//                   Bumping nwc restyles those apps rather than unstyling them.
//
// Consumers pin nwc by version, so the restyle reaches each app when it bumps.
module.exports = {
  content: [
    "./templates/**/*.html",
    "../nimsforestwebviewer/**/templates/**/*.html",
    "../**/templates/**/*.html",
  ],
  theme: {
    extend: {
      colors: {
        // --- Brand ---
        'paper': '#FAF6EC',        // page background
        'band': '#F5F0E2',         // alternating section background
        'band-2': '#F1EEE1',       // callout blocks, credit strip
        'card': '#FFFDF7',         // cards, pills, form fields
        'panel': '#F9F5E9',        // diagram panel
        'ink': '#16261C',          // headings, strong text
        'ink-2': '#24352A',        // lead paragraphs
        'body-text': '#4A5A4E',    // body copy
        'body-strong': '#35473A',  // list items, chips
        'muted': '#5E6F62',        // nav links
        'mono-label': '#6B7C6D',   // eyebrows, captions. WCAG-checked >=4.5:1 on paper: do not lighten
        'faint': '#8A9A8C',        // table row numbers
        'forest': '#1D4E32',       // primary buttons, accents, links
        'gold': '#B4801F',         // secondary accent, nim rings
        'dark-ground': '#12261B',  // dark sections
        'dark-ink': '#FBF7EC',     // headings on dark
        'dark-body': '#C3D0C4',
        'dark-body-2': '#A9BAAC',
        'dark-accent': '#C7E0B4',  // italics on dark
        'dark-accent-2': '#8FA894',// eyebrows on dark
        // Flow diagram
        'river': '#3F7A8C',
        'drop': '#6E8FA8',
        'leaf': '#6E9450',
        'wind': '#8A9A8C',
        'treehouse': '#8A6A3C',
        'hairline': 'rgba(22,38,28,0.10)',

        // --- Legacy names, repointed ---
        'forest-bg': '#FAF6EC',         // was #F8FAF5
        'forest-card': '#FFFDF7',       // was #FFFFFF
        'forest-green': '#1D4E32',      // was #4AA847
        'forest-green-dark': '#16261C', // was #3d8f3c, now the spec's button hover
        'canopy': '#16261C',            // was #1E3A1C
        'solar-gold': '#B4801F',        // was #E8B931
        'solar-gold-dark': '#8A6A3C',   // was #c9a029
        'soil': '#4A5A4E',              // was #6B5B4E
        'sprout': '#B9C6BA',            // was #A8D5A2. Used at /30-/40 for borders, so it has to stay soft
      },
      fontFamily: {
        display: ['Instrument Serif', 'ui-serif', 'Georgia', 'serif', 'Apple Color Emoji', 'Segoe UI Emoji', 'Noto Color Emoji', 'Nimsforest Emoji'],
        sans: ['Figtree', 'ui-sans-serif', 'system-ui', '-apple-system', 'sans-serif', 'Apple Color Emoji', 'Segoe UI Emoji', 'Noto Color Emoji', 'Nimsforest Emoji'],
        mono: ['IBM Plex Mono', 'ui-monospace', 'SFMono-Regular', 'monospace', 'Apple Color Emoji', 'Segoe UI Emoji', 'Noto Color Emoji', 'Nimsforest Emoji'],
      },
      letterSpacing: {
        eyebrow: '0.15em',
      },
      backgroundImage: {
        'hero-panel': 'linear-gradient(165deg, #F5F1E3 0%, #EFEAD9 100%)',
      },
      typography: {
        DEFAULT: {
          css: {
            '--tw-prose-body': '#4A5A4E',
            '--tw-prose-headings': '#16261C',
            '--tw-prose-lead': '#24352A',
            '--tw-prose-links': '#1D4E32',
            '--tw-prose-bold': '#16261C',
            '--tw-prose-counters': '#6B7C6D',
            '--tw-prose-bullets': '#B9C6BA',
            '--tw-prose-hr': 'rgba(22,38,28,0.10)',
            '--tw-prose-quotes': '#24352A',
            '--tw-prose-quote-borders': '#1D4E32',
            '--tw-prose-captions': '#6B7C6D',
            '--tw-prose-code': '#16261C',
            '--tw-prose-pre-code': '#16261C',
            '--tw-prose-pre-bg': '#FFFDF7',
            '--tw-prose-th-borders': 'rgba(22,38,28,0.14)',
            '--tw-prose-td-borders': 'rgba(22,38,28,0.08)',
            'h1': { fontFamily: "'Instrument Serif', ui-serif, Georgia, serif", fontWeight: '400' },
            'h2': { fontFamily: "'Instrument Serif', ui-serif, Georgia, serif", fontWeight: '400' },
            'h3': { fontFamily: "'Instrument Serif', ui-serif, Georgia, serif", fontWeight: '400' },
            'code': {
              fontFamily: "'IBM Plex Mono', ui-monospace, monospace",
              backgroundColor: 'rgba(22,38,28,0.05)',
              padding: '2px 6px',
              borderRadius: '4px',
              fontWeight: '400',
            },
            'code::before': { content: '""' },
            'code::after': { content: '""' },
            'pre': {
              border: '1px solid rgba(22,38,28,0.10)',
              borderRadius: '8px',
            },
            'th': {
              backgroundColor: 'rgba(22,38,28,0.04)',
            },
          },
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms')({ strategy: 'class' }),
  ],
}
