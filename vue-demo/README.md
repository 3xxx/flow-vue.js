# vue-demo

> A Vue.js project

## Build Setup

``` bash
# install dependencies
# 安装依赖，一般用cnpm……
npm install

# serve with hot reload at localhost:8080
# 本地热更新的开发状态，默认端口8080
npm run dev

# build for production with minification
# 打包给beego等使用
npm run build

# build for production and view the bundle analyzer report
npm run build --report

# run unit tests
npm run unit

# run e2e tests
npm run e2e

# run all tests
npm test
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).



D:\>cd gowork

D:\gowork>cd src

D:\gowork\src>cd github.com

D:\gowork\src\github.com>cd 3xxx

D:\gowork\src\github.com\3xxx>cd vue-demo

D:\gowork\src\github.com\3xxx\vue-demo>cnpm run dev

> vue-demo@1.0.0 dev D:\gowork\src\github.com\3xxx\vue-demo
> webpack-dev-server --inline --progress --config build/webpack.dev.conf.js

 17% building modules 66/96 modules 30 active ...\vue-demo\src\components\docstate.vue{ parser: "babylon" } is deprecated; we now treat it as { parser: "babel" }.
 95% emitting

 DONE  Compiled successfully in 11950ms


 I  Your application is running here: http://localhost:8080
 
```
 proxyTable: {
        "/api": {
            target: "http://localhost/v1/admin", //https://bjys.itdos.com/v1/admin---218.78.187.216/api/v1设置调用的接口域名和端口
            changeOrigin: true,
            pathRewrite: {
                "^/api": "" //用'/api' 代替 'http://218.78.187.216/api/v1'
            }
        },
```

上面配置proxytable没有用，在实际线上的时候没有用。
应该在config——dev.env.js里设置本地请求前缀
```
module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // API_ROOT: '"https://gcb.54lby.com/v1/admin"'  //本地请求前缀
  API_ROOT: '"http://localhost/v1/admin"'  //本地请求前缀
})
```
在config——dev.env.js里设置线上请求前缀
```
module.exports = merge(prodEnv, {
  NODE_ENV: '"production"',
  // API_ROOT: '"https://www.prov.com"'   //线上请求前缀
  API_ROOT: '"https://gcb.54lby.com/v1/admin"'   //线上请求前缀
})
```
修改以上文件需要重新运行cnpm run dev命令生效。

Install fail! Error: EBUSY: resource busy or locked
接下来我就说说自己整理的解决方案，首先肯定要删除 node_modules依赖包。
接下来
1. 如果你的项目里存在 package-lock.json 文件，删除它。并且删除 node_modules。然后再 npm install。
1. 第一步不行的话，运行 npm cache clean --force 或者 npm cache verify ，然后再 npm install / cnpm install。
1. 如果上面的都不行，就升级 npm， npm i -g npm，基本上是可以解决的
最后npm run dev / cnpm run dev

链接：https://www.jianshu.com/p/d2c68801dea7
