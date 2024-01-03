/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ['./cmd/app/partials/*.html', './cmd/app/pages/*.html'],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
