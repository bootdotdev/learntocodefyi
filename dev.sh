#!/bin/bash
# Use a trap to catch the Ctrl+C signal and kill all background processes
trap 'kill $(jobs -p)' SIGINT

# Start the Air server in the background
air -c ./.air.toml &

# Start Tailwind CLI with watch mode in the background
npx tailwindcss -i 'styles.css' -o 'public/styles.css' --watch &

# Start Browser-Sync in the background
browser-sync start \
  --files 'public/**/*.html, public/**/*.css' \
  --port 3001 \
  --proxy 'localhost:3000' \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }' &

# Wait for any process to exit
wait
