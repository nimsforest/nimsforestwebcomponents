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
    },
  },
  plugins: [],
}
