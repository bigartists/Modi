const ipMap = {
  mock: 'http://localhost:8080',
}

module.exports = function (env) {
  return {
    '/api/v1': {
      target: ipMap[env],
      changeOrigin: true,
      secure: false,
      pathRewrite: {
        '^/api/v1': '',
      },
    },
  }
}
