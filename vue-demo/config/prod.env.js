'use strict'

const merge = require('webpack-merge')
const prodEnv = require('./prod.env')
 
module.exports = merge(prodEnv, {
  NODE_ENV: '"production"',
  // API_ROOT: '"https://www.prov.com"'   //线上请求前缀
  API_ROOT: '"http://localhost/v1/admin"'   //线上请求前缀
})

// module.exports = {
//   NODE_ENV: '"production"'
// }
