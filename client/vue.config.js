module.exports = {
  devServer: {
    proxy: {
      '^/server': {
        target: 'http://localhost:9000/',
        changeOrigin: true,
        secure: false,
        ws: true,
        pathRewrite: { '^/server': '/' }
      }
    }
  }
}
