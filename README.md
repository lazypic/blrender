# blrender
- 블랜더를 쉽게 렌더링하기 위한 코드이다.
- EXR만 처리하도록 만든 간단한 코드이다.
- OSX 에서만 테스트됨.

# Install
```
go get github.com/lazypic/blrender
```

# 사용법
- 터미널에서 렌더링할 블랜더파일을 blrender 명령어 뒤에 붙혀준다.

```
$ blrender render.blend
$ blrender /file/location/render.blend
$ blrender ~/test/render.blend
$ blrender ../render.blend
$ blrender .././render.blend
```
