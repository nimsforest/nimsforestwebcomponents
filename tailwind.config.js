/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html",
    "../nimsforestwebviewer/**/templates/**/*.html",
    "../**/templates/**/*.html",
  ],
  theme: {
    extend: {
      colors: {
        'forest-bg': '#F8FAF5',
        'forest-card': '#FFFFFF',
        'forest-green': '#4AA847',
        'forest-green-dark': '#3d8f3c',
        'canopy': '#1E3A1C',
        'solar-gold': '#E8B931',
        'solar-gold-dark': '#c9a029',
        'soil': '#6B5B4E',
        'sprout': '#A8D5A2',
      },
      fontFamily: {
        mono: ['JetBrains Mono', 'ui-monospace', 'SFMono-Regular', 'monospace'],
      },
      typography: {
        DEFAULT: {
          css: {
            '--tw-prose-body': '#6B5B4E',
            '--tw-prose-headings': '#1E3A1C',
            '--tw-prose-lead': '#6B5B4E',
            '--tw-prose-links': '#4AA847',
            '--tw-prose-bold': '#1E3A1C',
            '--tw-prose-counters': '#6B5B4E',
            '--tw-prose-bullets': '#A8D5A2',
            '--tw-prose-hr': '#A8D5A2',
            '--tw-prose-quotes': '#6B5B4E',
            '--tw-prose-quote-borders': '#4AA847',
            '--tw-prose-captions': '#6B5B4E',
            '--tw-prose-code': '#1E3A1C',
            '--tw-prose-pre-code': '#1E3A1C',
            '--tw-prose-pre-bg': '#FFFFFF',
            '--tw-prose-th-borders': 'rgba(168,213,162,0.3)',
            '--tw-prose-td-borders': 'rgba(168,213,162,0.3)',
            'code': {
              fontFamily: "'JetBrains Mono', monospace",
              backgroundColor: 'rgba(168,213,162,0.2)',
              padding: '2px 6px',
              borderRadius: '4px',
              fontWeight: '400',
            },
            'code::before': { content: '""' },
            'code::after': { content: '""' },
            'pre': {
              border: '1px solid rgba(168,213,162,0.3)',
              borderRadius: '8px',
            },
            'th': {
              backgroundColor: 'rgba(168,213,162,0.15)',
            },
          },
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
