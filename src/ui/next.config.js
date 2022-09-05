/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    API_PROTOCOOL: process.env.API_PROTOCOOL,
    API_URL: process.env.API_URL,
    API_PORT: process.env.API_PORT
  },
  async rewrites() {
    return [
      {
        source: '/api/v1/:slug*',
        destination: 'http://localhost:8080/api/v1/:slug*' // Proxy to Backend
      }
    ]
  }
}

module.exports = nextConfig
