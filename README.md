# SequenceTimer

SequenceTimer is a Go application that allows you to set up multiple time intervals (sequences). After activation, it
will play a sound at the end of each interval until the sequence is completed.

## Features

- Define multiple intervals for custom time sequences
- Receive an audio notification when each interval finishes
- Runs sequentially until all intervals are complete

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/JK014/SequenceTimer.git
    ```
2. Navigate to the project directory:
   ```bash
    cd SequenceTimer
    ```
3. Build the project:
   ```bash
    go build
    ```
4. Run the executable:
   ```bash
    ./SequenceTimer
    ```

## Usage

Configure your intervals and start the timer. A sound notification will play once each interval ends.

## License

SequenceTimer is licensed under the GNU General Public License version 3 (GPLv3). You can find a copy of the full
license text in the LICENSE file included with this project, or online
at [GNU.org](https://www.gnu.org/licenses/gpl-3.0.html).
