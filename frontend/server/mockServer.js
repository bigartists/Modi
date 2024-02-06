const http = require('http')
const { get } = require('lodash')
const routers = [
  {
    path: '/login',
    callback: (req, res) => {
      let result = ''
      req.on('data', data => {
        result += data
      })

      req.on('end', function () {
        const data = JSON.parse(result)
        console.log('🚀 ~ file: mockServer.js:13 ~ data:', data)
        const username = get(data, 'username')
        res.writeHead(200, { 'Content-Type': 'text/json' })
        if (!username) {
          res.end(
            JSON.stringify({
              message: '登录失败',
              status: 0,
              result: {},
            }),
          )
        } else {
          res.end(
            JSON.stringify({
              message: '登录成功',
              status: 1,
              result: {
                username,
              },
            }),
          )
        }
      })
    },
  },
  {
    path: '/register',
    callback: (req, res) => {
      res.writeHead(200, { 'Content-Type': 'text/plain' })
      res.end('register')
    },
  },
]
http
  .createServer((req, res) => {
    routers.forEach(router => {
      handler(router.path, router.callback)(req, res)
    })
  })
  .listen(8080, () => {
    console.log('server is running at http://localhost:8080')
  })

function handler(path, callback) {
  return (req, res) => {
    if (req.url === path) {
      callback(req, res)
    }
  }
}
