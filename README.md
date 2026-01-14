# ocenv
Script that utilizes [tmux](https://github.com/tmux/tmux) and
[fzf](https://github.com/junegunn/fzf) for easier OpenShift
cluster management.

## Usage
Open fzf menu and switch to selected session
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
List all available environments:
```sh
ocenv list
```
List active environment sessions:
```sh
ocenv list --active
```
Get environment context information
```sh
ocenv info <name>
```

## TODO
- perhaps add kubectl support or sum
- completions mhmhmhm

## Special Thanks
[tmux-sessionizer](https://github.com/ThePrimeagen/tmux-sessionizer) - for
showing me da wae
