/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'dark-primary': '#1a202c', // สีพื้นหลังหลัก
        'dark-secondary': '#2d3748', // สีพื้นหลังรอง
        'dark-accent': '#4a5568', // สีเน้น
        'dark-text': '#edf2f7', // สีข้อความ
        'dark-border': '#718096', // สีขอบ
      },
      fontFamily: {
        sans: ['Prompt', 'sans-serif'],
      }
    },
  },
  plugins: [],
}
