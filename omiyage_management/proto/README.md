# proto
omiyage management serviceのproto置き場

## Settings
`make install`

### JetBrains
手動で設定を行わないとimportしてきたもののsuggestが効かなくなるので開発しづらい.

`Languages & Frameworks > Protocol Buffers` の `configuration automatically` を off にして, `proto/src` を import Paths に追加する. 


## Generate
`make gen`

## References
- https://buf.build/docs/introduction/