# nav-api

[nav](https://nav.programnotes.cn)'s go serverless api,using pg & gorm

## vercel 里go的api

不是api的util代码必须放在_+目录名下如_pkg

<https://github.com/vercel/vercel/discussions/7000>

### vercel go api和next-auth冲突了

next-auth里通过前端代码实现了auth接口,同时使用go serverless api 在构建时会覆盖?两边的api没法同时存在于一个项目
参考:
<https://github.com/nextauthjs/next-auth/issues/6684>
