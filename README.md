# ocenv
Script that utilizes [tmux](https://github.com/tmux/tmux) and
[fzf](https://github.com/junegunn/fzf) for easier OpenShift
cluster management.

## Usage
Open fzf menu with active sessions
```sh
ocenv
```
Create environment:
```sh
ocenv create <name>
```
Remove environment:
```sh
ocenv rm <name>
```
Create tmux environment for cluster:
```sh
ocenv use <name>
```
List all available anvironments:
```sh
ocenv list
```

## TODO
- perhaps add kubectl support or sum

## Special Thanks
[tmux-sessionizer](https://github.com/ThePrimeagen/tmux-sessionizer) - for
showing me da wae
