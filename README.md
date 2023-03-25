# TOC TRADE PROTOBUF

## VScode setting

- cp settings.template.json .vscode/settings.json

```sh
mkdir .vscode
home_path="$HOME"
echo '{
    "protoc": {
        "path": "'$home_path'/sdk_tools/protoc/bin/protoc",
        "compile_on_save": false,
        "options": [
            "--proto_path=protos/v3/app",
            "--proto_path=protos/v3/basic",
            "--proto_path=protos/v3/common",
            "--proto_path=protos/v3/health",
            "--proto_path=protos/v3/history",
            "--proto_path=protos/v3/stream",
            "--proto_path=protos/v3/trade",
            "--python_out=gen/python",
            "--go_out=gen/go",
            "--python_out=gen/python"
        ]
    }
}' > .vscode/settings.json
```
