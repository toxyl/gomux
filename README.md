# GoMux
... is a simple controller for `tmux`.

# Installation
To install GoMux, run the following commands:

```bash
sudo apt install tmux 
git clone https://github.com/toxyl/gomux
cd gomux
chmod +x build.sh
./build.sh
```

# Usage
| Command | Args | Description |
| --- | --- | --- |
| `gomux start` | `[config file]` | Starts a tmux session using the given config file. |
| `gomux daemon` | `[config file]` | Starts a tmux session in the background using the given config file. | 
| `gomux detach` | `[config file]` | Detaches all clients connected to the session started with the given config file. | 
| `gomux list` | | Lists all active sessions. |

# Configs
Here's an example config:
```yaml
name: demo
panes:
- command: echo "Hello world"
- command: echo "My local IP is $(hostname -I)" 
  delay: 2

```

# Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

# License
This project is licensed under the UNLICENSE - see the LICENSE file for details.
