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
        'forest-dark': '#1a1a2e',
        'forest-darker': '#141425',
        'forest-green': '#2ecc71',
        'forest-green-dark': '#27ae60',
        'solar-gold': '#f39c12',
        'solar-gold-dark': '#d68910',
        'life-blue': '#3498db',
        'bark-brown': '#8b6914',
      },
    },
  },
  plugins: [],
}
