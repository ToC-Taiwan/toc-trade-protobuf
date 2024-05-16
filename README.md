# TOC TRADE PROTOBUF

## VScode setting

- cp settings.template.json .vscode/settings.json

## VScode settings

```json
protoc_path=$(which protoc)
echo '{
    "protoc": {
        "path": "'$protoc_path'",
        "compile_on_save": false,
        "options": [
            "--proto_path=protos/v3",
            "--python_out=src/python",
            "--go_out=src/go"
        ]
    }
}' > .vscode/settings.json
```
