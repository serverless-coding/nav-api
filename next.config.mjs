/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  experimental: {
    appDir: true,
  },
  images: {
    domains: [
      "programnotes.cn",
      "nav.programnotes.cn"
    ],
  }
}

export default nextConfig
