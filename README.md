# 4BOD-Go
A Golang reimplementation of 4BOD fantasy console

## Download
- ### [Windows v1.0.0 Download (zip)](https://github.com/Sam36502/4BOD-Go/releases/download/v1.0.0/4BoD-Go_v1.0.0_Windows.zip)
- ### [Linux v1.0.0 Download (zip)](https://github.com/Sam36502/4BOD-Go/releases/download/v1.0.0/4BoD-Go_v1.0.0_Linux.zip)

## Usage
After startup, you can either choose to edit the current
program, or run it. Use arrow keys and enter to navigate
the menus and backspace to return to it again.

In the editor
view, the bar on the right shows what "page" of the editor
you're on which you can scroll up and down using the arrow keys.
Click on pixels in the left side of the window to toggle the bits
of the program. The leftmost 4 bits are the instruction and the
two sets of 4 bits to the right of it (i.e. middle 8 pixels) are
the two operands or arguments. To make it easier to see these divisions
you can toggle an overlay with tab.

At any point you can press 'S' to save your current program to a file
and you can press 'L' to load one. The files are stored in binary format. (`.4bb`)

See [The esolangs page](https://esolangs.org/wiki/4BOD) or
[The original itch.io page](https://puarsliburf.itch.io/4bod-fantaly-console) for more
information on how to use it.

## Examples
The programs in the `examples` directory were mostly assembled with [my 4bod compiler](https://github.com/Sam36502/4BOD-Assembler)
so for each `.4bb` binary file, I've also included the `.4sm` assembly file which might
make it a bit easier to understand what the examples do

## Options
If you want to customise the interface, you can do so with the included `options.json` file.
Options include:
| JSON Key         | Description |
|------------------|-------------|
| `splash_millis`  | How many milliseconds to spend on the splash screen |
| `pixel_size`     | Side length of the 4BOD pixels in real pixels |
| `target_fps`     | What framerate to limit the program to. Helps to see what's actually happening (set to -1 for no limit) |
| `color_fg`       | The colour of foreground pixels |
| `color_bg`       | The colour of background |
| `color_overlay`  | The colour of the editor overlay |
| `old_menu`       | Whether to use the old menu images |
| `editor_overlay` | Whether to have the editor overlay on by default (Should save if turned off with it on) |
| `debug_keycodes` | Whether to display the last pressed keycode on the screen (helpful for changing controls) |
| `controls`       | A list of various keys and their keycodes (see `debug_keycodes`) |

## Changing Keyboard Inputs
The easiest way to change which keys do what is to set the `debug_keycodes` option by setting it to `true`
and starting the machine. Then, you can press the keys you want each thing to do and write down what the
keycode is. After that, you can change the respective `kc_...` settings in the controls part of `options.json`.
You can also use `0` to unbind the key.

