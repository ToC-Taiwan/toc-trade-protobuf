# TOC TRADE PROTOBUF

## VScode setting

- cp settings.template.json .vscode/settings.json

```sh
mkdir .vscode
home_path="$HOME"
echo '{
    "protoc": {
        "path": "/Users/timhsu/sdk_tools/protoc/bin/protoc",
        "compile_on_save": false,
        "options": [
            "--proto_path=protos/v3",
            "--python_out=gen/python",
            "--go_out=gen/go"
        ]
    }
}
' > .vscode/settings.json
```

## Install Python protobuf

```sh
make venv
pip3 install -U \
  --no-warn-script-location \
  --no-cache-dir \
  grpcio \
  grpcio-tools
``````
