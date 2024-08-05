# ArtikelJagd

ArtikelJagd(hunt for articles) is a Go-based game using the Ebiten game library. The goal of the game is to match German nouns to their gender articles.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/RheinhardtSnyman/ArtikelJagd.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd ArtikelJagd
   ```

3. **Install the dependencies:**

   Ensure you have Go installed on your machine. If not, download and install it from [here](https://go.dev/doc/install).

   _Ubuntu_

   Go instalation

   ```bash
   sudo snap install go --classic && export PATH=$PATH:/usr/local/go/bin
   ```

   Ebiten dependencies

   ```bash
    sudo apt-get install libgl1-mesa-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libxrandr-dev libxxf86vm-dev
   ```

   Install the dependencies

   ```bash
   go mod tidy
   ```

## Running the Game

You can run the game in either normal mode or demo mode.

### Demo Mode

Demo mode shows a proof of concept in simple English where the objective is to match colors.

To run the game in demo mode, use:

```bash
go run .\cmd\ArtikelJagd\main.go -demo
```

### Normal Mode

In normal mode, the game is in German.

To run the game in normal mode, use:

```bash
go run .\cmd\ArtikelJagd\main.go
```
