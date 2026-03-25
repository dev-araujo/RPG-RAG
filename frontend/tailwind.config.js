/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        "on-primary-fixed-variant": "#7f2926",
        "surface": "#fffcf7",
        "surface-container-lowest": "#ffffff",
        "surface-container": "#f8f3e8",
        "on-primary-container": "#e3756d",
        "primary-fixed-dim": "#ffb3ad",
        "secondary-fixed": "#e5e2e1",
        "tertiary-container": "#3e2b00",
        "outline": "#89726f",
        "inverse-surface": "#35301e",
        "on-surface-variant": "#564240",
        "on-error": "#ffffff",
        "primary-container": "#5d1010",
        "inverse-primary": "#ffb3ad",
        "secondary-fixed-dim": "#c8c6c5",
        "on-primary-fixed": "#410003",
        "secondary": "#5f5e5e",
        "on-secondary-fixed-variant": "#474746",
        "on-tertiary-container": "#b3914f",
        "outline-variant": "#e5ddd3",
        "primary-fixed": "#ffdad6",
        "error": "#ba1a1a",
        "on-primary": "#ffffff",
        "primary": "#3d0003",
        "error-container": "#ffdad6",
        "surface-container-high": "#f1e8cd",
        "tertiary": "#241700",
        "tertiary-fixed-dim": "#e7c17a",
        "background": "#fffcf7",
        "on-tertiary": "#ffffff",
        "tertiary-fixed": "#ffdea4",
        "surface-variant": "#ebe2c8",
        "on-surface": "#1f1c0b",
        "on-tertiary-fixed-variant": "#5c4205",
        "surface-tint": "#9e403b",
        "on-tertiary-fixed": "#261900",
        "surface-dim": "#e2dabf",
        "surface-container-low": "#fdfaf2",
        "on-background": "#1f1c0b",
        "on-secondary": "#ffffff",
        "on-error-container": "#93000a",
        "surface-container-highest": "#ebe2c8",
        "on-secondary-fixed": "#1c1b1b",
        "secondary-container": "#e2dfde",
        "surface-bright": "#fffcf7",
        "inverse-on-surface": "#faf0d5",
        "on-secondary-container": "#636262"
      },
      fontFamily: {
        "headline": ["Noto Serif"],
        "body": ["Work Sans"],
        "label": ["Space Grotesk"]
      },
      borderRadius: {
        "DEFAULT": "0", "lg": "0", "xl": "0", "full": "9999px"
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/container-queries')
  ],
}
