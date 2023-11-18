const { hostname } = require('os')

/** @type {import('next').NextConfig} */
const hostnames = [
    'media.istockphoto.com',
    'img.freepik.com',
    'www.shutterstock.com'
]

const nextConfig = {
    images: {
        remotePatterns: hostnames.map(hostname =>({
            protocol: 'https',
            hostname
        }))
    }
}

module.exports = nextConfig
