# go-switch-walls

A Wayland wallpaper switcher with a terminal UI. Integrates with `matugen` for
color scheme generation and `awww` for animated transitions.

## Dependencies

- A Wayland compositor (Hyprland, Niri, SwayWM, etc.)
- [`matugen`](https://github.com/InioX/matugen)
- [`awww`](https://codeberg.org/LGFae/awww)

## Setup

Place your wallpapers (`.png`, `.jpg`) in `~/Pictures/Walls`.

## Installation

**With [`just`](https://github.com/casey/just):**

```sh
just install
```

This builds the binary and moves it to `~/.local/bin` automatically.

**Manually:**

```sh
go build .
mv gsw ~/.local/bin/
```

## Usage

```sh
gsw
```

## Keybindings

| Key            | Action                                                |
| -------------- | ----------------------------------------------------- |
| `j` / `↓`      | Move cursor down                                      |
| `k` / `↑`      | Move cursor up                                        |
| `h` / `←`      | Previous page                                         |
| `l` / `→`      | Next page                                             |
| `Tab`          | Switch focus between wallpaper list and color schemes |
| `s`            | Toggle dark / light mode                              |
| `Enter`        | Apply wallpaper and generate color scheme             |
| `q` / `Ctrl+C` | Quit                                                  |
