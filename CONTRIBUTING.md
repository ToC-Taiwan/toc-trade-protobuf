# Contributing

## Tools

### Conventional Commit

- install git cz tool global

```sh
npm install -g commitizen
npm install -g cz-conventional-changelog
npm install -g conventional-changelog-cli
echo '{ "path": "cz-conventional-changelog" }' > ~/.czrc
```

### Pre-commit

- install git pre-commit tool global(macOS)

```sh
brew install pre-commit
```

- install/modify from config

```sh
pre-commit autoupdate
pre-commit install
pre-commit run --all-files
```

### Modify CHANGELOG

- git-chglog

```sh
brew tap git-chglog/git-chglog
brew install git-chglog
```

```sh
COMMIT_HASH=053ff921b22726015638934803fa228113f58bf1
VERSION=0.1.9
git tag -a v$VERSION $COMMIT_HASH -m $VERSION
git push -u origin --all
git push -u origin --tags
```

### Find ignored files

```sh
find . -type f  | git check-ignore --stdin
```
