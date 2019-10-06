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

 proxyTable: {
        "/api": {
            target: "http://localhost/v1/admin", //https://bjys.itdos.com/v1/admin---218.78.187.216/api/v1设置调用的接口域名和端口
            changeOrigin: true,
            pathRewrite: {
                "^/api": "" //用'/api' 代替 'http://218.78.187.216/api/v1'
            }
        },


