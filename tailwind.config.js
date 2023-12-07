/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ['./partials/**/*.html', './pages/**/*.html'],
  purge: ['./partials/**/*.html', './pages/**/*.html'],
  theme: {
    extend: {},
  },
  plugins: [],
}
